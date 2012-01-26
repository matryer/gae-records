package gaerecords

import (
	"os"
	"fmt"
	"reflect"
	"appengine"
	"appengine/datastore"
)

// The ID value of a record that indicates there is no ID.  A record
// will have no ID if it has not yet been saved, or if it has been deleted.
var NoIDValue int64 = 0

// Represents a single record of data (like a single row in a database, or a single resource
// on a web server).  Synonymous with an Entity in appengine/datastore.
type Record struct {

	// internal storage of record field data.
	fields map[string]interface{}

	// a reference to the model describing the
	// type of this record.
	model *Model

	// an internal cache of the datastore.Key
	datastoreKey *datastore.Key

	// whether the record needs persisting or not
	needsPersisting bool

	// internal storage of this record's ID
	recordID int64
}

/*
	Constuctors
	----------------------------------------------------------------------
*/

// Creates a new record of the given Model type.  Not recommended.  Instead call the
// New() method on the model object itself.
func NewRecord(model *Model) *Record {

	// create and setup the record
	record := new(Record).
		configureRecord(model, nil).
		SetNeedsPersisting(true)

	// trigger the event
	model.AfterNew.Trigger(record)

	// return the record
	return record
}

/*
	Properties
	----------------------------------------------------------------------
*/

// Gets the current Model object representing the type of this record.
func (r *Record) Model() *Model {
	return r.model
}

// Sets the current Model object representing the type of this record.  It is recommended that
// you create records with model.New() or use NewRecord(*Model) instead of using this method
// directly.
func (r *Record) SetModel(model *Model) *Record {
	r.model = model
	return r
}

// Gets a human readable string representation of this record
func (r *Record) String() string {

	if r.IsPersisted() {
		return fmt.Sprintf("{Record:model=%v,id=%v}", r.model.String(), r.ID())
	}

	return fmt.Sprintf("{Record:model=%v}", r.model.String())
}

/*
	IDs
	----------------------------------------------------------------------
*/

// Gets the unique ID for this record.  A record will be assigned a unique ID
// only when it is persisted in the datastore.  Otherwise, the ID will be equal to NoIDValue.
// Use IsPersisted() to check if a record has been persisted in the datastore or not.
func (r *Record) ID() int64 {
	return r.recordID
}

// Sets the ID for this record.  Used internally.
func (r *Record) setID(id int64) *Record {

	// set the record ID
	r.recordID = id

	r.invalidateDatastoreKey()

	// chain
	return r
}

/*
	Persistence
	----------------------------------------------------------------------
*/

// Gets whether this record needs persisting in the datastore or not.  If this
// record is synched with the datastore (as far as this record knows) it will
// return false, otherwise, if something has changed or this is a new record, true
// will be returned.
func (r *Record) NeedsPersisting() bool {
	return r.needsPersisting
}

// Sets whether this record needs persisting or not.  Advanced use only.
func (r *Record) SetNeedsPersisting(value bool) *Record {
	r.needsPersisting = value
	return r
}

// CAUTION: This method does NOT load persisted records.  See Find().
// PropertyLoadSaver.Load takes a channel of datastore.Property objects and
// applies them to the internal Fields() object.
// Used internally by the datastore.
func (r *Record) Load(c <-chan datastore.Property) os.Error {

	// load the fields
	for f := range c {

		if f.Multiple {

			// do we already have this value?
			if r.Fields()[f.Name] == nil {

				// create a slice to hold these objects
				r.Fields()[f.Name] = make([]interface{}, 0, 0)

			}

			// add this object to the slice
			r.Fields()[f.Name] = reflect.Append(reflect.ValueOf(r.Fields()[f.Name]), reflect.ValueOf(f.Value)).Interface()

		} else {

			// load single value
			r.Fields()[f.Name] = f.Value

		}

	}

	// no errors
	return nil
}

// CAUTION: This method does NOT persist records.  See Put().
// PropertyLoadSaver.Save writes datastore.Property objects and
// representing the Fields() of this record to the specified channel.
// Used internally by the datastore to persist the values.
func (r *Record) Save(c chan<- datastore.Property) os.Error {

	for k, v := range r.Fields() {

		if reflect.TypeOf(v).Kind() == reflect.Array || reflect.TypeOf(v).Kind() == reflect.Slice {

			// multiple values - iterate over each value
			// and add them as seperate properties

			value := reflect.ValueOf(v)
			l := value.Len()

			for i := 0; i < l; i++ {

				thisVal := value.Index(i)

				// create the property
				c <- datastore.Property{
					Name:     k,
					Value:    thisVal.Interface(),
					Multiple: true,
				}

			}

		} else {

			// single value
			// create the property
			c <- datastore.Property{
				Name:     k,
				Value:    v,
				Multiple: false,
			}

		}

	}

	// this channel is finished
	close(c)

	// no errors
	return nil
}

// (Internal) Configures a Record after it has been found or created using means other than
// model.New() or NewRecord(model).
func (r *Record) configureRecord(model *Model, key *datastore.Key) *Record {

	return r.
		SetModel(model).
		SetDatastoreKey(key).
		SetNeedsPersisting(false)

}

// Saves or updates this record.  Returns nil if successful, otherwise returns the os.Error
// that was retrned by the datastore.
//  record.Put()
//
// Raises events:
//   Model.BeforePut with Args(record)
//   Model.AfterPut with Args(record)
func (r *Record) Put() os.Error {

	// trigger the BeforePut event on the model
	context := r.Model().BeforePut.Trigger(r)

	if !context.Cancel {

		newKey, err := datastore.Put(GetAppEngineContext(), r.DatastoreKey(), datastore.PropertyLoadSaver(r))

		if err == nil {

			// update the record
			r.SetDatastoreKey(newKey).SetNeedsPersisting(false)

			// trigger the AfterPut event
			r.Model().AfterPut.TriggerWithContext(context)

			return nil

		}

		return err

	}

	return ErrOperationCancelledByEventCallback

}

// Deletes this record.  Returns nil if successful, otherwise returns the os.Error
// that was retrned by appengime/datastore.
//   record.Delete()
//
// Raises events:
//   Model.BeforeDelete with Args(id, record)
//   Model.AfterDelete with Args(id, record)
// Note: The Record will be passed to the events.
func (r *Record) Delete() os.Error {

	// trigger the BeforeDeleteByID event
	context := r.Model().BeforeDelete.Trigger(r.ID(), r)

	if !context.Cancel {

		err := datastore.Delete(GetAppEngineContext(), r.DatastoreKey())

		if err == nil {

			// clean up the record
			r.setID(NoIDValue)

			// trigger the AfterDeleteByID event
			r.Model().AfterDelete.TriggerWithContext(context)

		}

		return err

	}

	return ErrOperationCancelledByEventCallback

}

/*
	Datastore Key
	----------------------------------------------------------------------
*/

// Gets the appengine/datastore Key for this record.  If this record is persisted in the
// datastore it wil be a complete key, otherwise, this method will return an incomplete key.
func (r *Record) DatastoreKey() *datastore.Key {

	if r.datastoreKey == nil {

		var key *datastore.Key

		if r.IsPersisted() {
			key = r.model.NewKeyWithID(r.ID())
		} else {
			key = r.model.NewKey()
		}

		r.datastoreKey = key

	}

	return r.datastoreKey

}

// Sets the datastore Key and updates the records ID if needed
func (r *Record) SetDatastoreKey(key *datastore.Key) *Record {

	if key == nil {

		r.setID(NoIDValue)
		r.datastoreKey = nil

	} else {

		// does the key have an ID?
		if key.IntID() > 0 {

			// set the ID
			r.setID(key.IntID())

		}

		// set the key
		r.datastoreKey = key

	}

	// chain
	return r

}

// Invalidates the internally cached datastore key for this
// record so that when it is next requested via DatastoreKey() it will
// be regenerated to match the corrected state
func (r *Record) invalidateDatastoreKey() {
	r.datastoreKey = nil
}

// Whether this record has been persisted in the
// datastore or not, i.e. record.ID != NoIDValue
func (r *Record) IsPersisted() bool {
	return r.recordID != NoIDValue
}

/*
	Fields
	----------------------------------------------------------------------
*/

// Gets the internal storage map (map[string]interface{}) that contains the
// persistable fields for this record.  Instead of manipulating this object directly,
// you should use the Get*() and Set*() methods.
func (r *Record) Fields() map[string]interface{} {

	// ensure we have a map to store the fields
	if r.fields == nil {
		r.fields = make(map[string]interface{})
	}

	return r.fields

}

/*
	Getting Fields
	----------------------------------------------------------------------
*/

// Gets the value of a field in a record.  Strongly typed alternatives are provided and recommended
// to use where possible.
func (r *Record) Get(key string) interface{} {

	if r == nil {
		panic(fmt.Sprintf("gaerecords: Cannot Get(\"%v\") property from a nil Record", key))
	}

	return r.Fields()[key]
}

// Gets an []interface{} of the multiple values contained in a single property.
// For example:
//  // create a model
//  model := NewModel("people")
//
//  // create an item with a 'tags' property containing a slice of strings
//  // and Put this item
//  item, _ := model.New().Set("tags", []string{"one", "two", "three"}).Put()
// 
//  // load the item again
//  item = model(item.ID())
//
//  // get the tags
//  for i, tag := range item.GetMultiple("tags") {
//	  // cast tag and do soemthing with it
//  }
func (r *Record) GetMultiple(key string) []interface{} {
	return r.Get(key).([]interface{})
}

// Gets the i'th item from an array or slice property.  If you plan to iterate over
// all of the items, see GetMultiple() instead.
func (r *Record) GetMultipleItem(key string, i int) interface{} {
	return r.GetMultiple(key)[i]
}

// Gets the number of items in an array or slice property.  If you plan to iterate over
// all of the items, see GetMultiple() instead.
func (r *Record) GetMultipleLen(key string) int {
	return len(r.GetMultiple(key))
}

/*
	Setting
	----------------------------------------------------------------------
*/

// Sets a field in the record.  The value must be an acceptable datastore
// type or another Record.  Strongly typed alternatives are provided and recommended
// to use where possible.
//
// bug(matryer): Setting values of type []byte fails. See https://github.com/matryer/gae-records/issues/1
func (r *Record) Set(key string, value interface{}) *Record {

	fields := r.Fields()
	oldValue := fields[key]
	fields[key] = value

	// trigger the OnChanged event if we need to
	if r.model.OnChanged.HasCallbacks() {
		r.model.OnChanged.Trigger(r, key, value, oldValue)
	}

	if value != oldValue {
		r.SetNeedsPersisting(true)
	}

	return r
}

/*
	Strongly typed getters and setters
	----------------------------------------------------------------------
*/

// Gets a string field
func (r *Record) GetString(key string) string {
	return fmt.Sprint(r.Get(key))
}

// Sets the string value of a field
func (r *Record) SetString(key string, value string) *Record {
	return r.Set(key, value)
}

func (r *Record) SetMultipleStrings(key string, value []string) *Record {
	return r.Set(key, value)
}

// Gets the int64 value of a field with the specified key.
func (r *Record) GetInt64(key string) int64 {
	return r.Get(key).(int64)
}

// Sets the int64 value of a field with the specified key.
func (r *Record) SetInt64(key string, value int64) *Record {
	return r.Set(key, value)
}

func (r *Record) SetMultipleInt64s(key string, value []int64) *Record {
	return r.Set(key, value)
}

// Gets the float64 value of a field with the specified key.
func (r *Record) GetFloat64(key string) float64 {
	return r.Get(key).(float64)
}

// Sets the float64 value of a field with the specified key.
func (r *Record) SetFloat64(key string, value float64) *Record {
	return r.Set(key, value)
}

func (r *Record) SetMultipleFloat64s(key string, value []float64) *Record {
	return r.Set(key, value)
}

// Gets the bool value of a field with the specified key.
func (r *Record) GetBool(key string) bool {
	return r.Get(key).(bool)
}

// Sets the bool value of a field with the specified key.
func (r *Record) SetBool(key string, value bool) *Record {
	return r.Set(key, value)
}

func (r *Record) SetMultipleBools(key string, value []bool) *Record {
	return r.Set(key, value)
}

// Gets the *datastore.Key value of a field with the specified key.
func (r *Record) GetKeyField(key string) *datastore.Key {
	return r.Get(key).(*datastore.Key)
}

// Sets the *datastore.Key value of a field with the specified key.
func (r *Record) SetKeyField(key string, value *datastore.Key) *Record {
	return r.Set(key, value)
}

func (r *Record) SetMultipleKeys(key string, value []*datastore.Key) *Record {
	return r.Set(key, value)
}

// Gets the appengine.BlobKey value of a field with the specified key.
func (r *Record) GetBlobKey(key string) appengine.BlobKey {
	return r.Get(key).(appengine.BlobKey)
}

// Sets the appengine.BlobKey value of a field with the specified key.
func (r *Record) SetBlobKey(key string, value appengine.BlobKey) *Record {
	return r.Set(key, value)
}

func (r *Record) SetMultipleBlobKeys(key string, value []appengine.BlobKey) *Record {
	return r.Set(key, value)
}

// Gets the []byte value of a field with the specified key.
func (r *Record) GetBytes(key string) []byte {
	return r.Get(key).([]byte)
}

// Sets the []byte value of a field with the specified key.
func (r *Record) SetBytes(key string, value []byte) *Record {
	return r.Set(key, value)
}

func (r *Record) SetMultipleBytes(key string, value [][]byte) *Record {
	return r.Set(key, value)
}

// Gets the datastore.Time value of a field with the specified key.
func (r *Record) GetTime(key string) datastore.Time {
	return r.Get(key).(datastore.Time)
}

// Sets the datastore.Timevalue of a field with the specified key.
func (r *Record) SetTime(key string, value datastore.Time) *Record {
	return r.Set(key, value)
}

func (r *Record) SetMultipleTimes(key string, value []datastore.Time) *Record {
	return r.Set(key, value)
}

/*
	Errors
	----------------------------------------------------------------------
*/

// Causes the record to panic
func (r *Record) panic(message string) {

}
