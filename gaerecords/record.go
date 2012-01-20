package gaerecords

import (
	"os"
	"fmt"
	"appengine/datastore"
)

// The ID value of a record that indicates there is no ID.  A record
// will have no ID if it has not yet been saved, or if it has been deleted.
var NoIDValue int64 = 0

// Represents a single record of data (like a single row in a database, or a single resource
// on a web server).  Synonymous with an Entity in appengine/datastore.
type Record struct {
	fields map[string]interface{}

	model *Model

	datastoreKey *datastore.Key

	recordID int64
}

/*
	Constuctors
	----------------------------------------------------------------------
*/

// Creates a new record of the given Model type.  Not recommended.  Instead call the
// New() method on the model object itself.
func NewRecord(model *Model) *Record {

	record := new(Record)
	record.model = model
	return record

}

/*
	Properties
	----------------------------------------------------------------------
*/

// Gets the current Model object describing this type of record.
func (r *Record) Model() *Model {
	return r.model
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

// CAUTION: This method does NOT load persisted records.  See Find().
// PropertyLoadSaver.Load: Takes a channel of datastore.Property objects and
// applies them to the internal Fields() object.
// Used internally by the datastore.
func (r *Record) Load(c <-chan datastore.Property) os.Error {

	// load the fields
	for f := range c {
		r.Fields()[f.Name] = f.Value
	}

	// no errors
	return nil
}

// CAUTION: This method does NOT persist records.  See Put().
// PropertyLoadSaver.Save: Writes datastore.Property objects and
// representing the Fields() of this record to the specified channel.
// Used internally by the datastore to persist the values.
func (r *Record) Save(c chan<- datastore.Property) os.Error {

	for k, v := range r.Fields() {
		c <- datastore.Property{
			Name:  k,
			Value: v,
		}
	}

	// this channel is finished
	close(c)

	// no errors
	return nil
}

// Saves or updates this record.  Returns nil if successful, otherwise returns the os.Error
// that was retrned by appengime/datastore.
//  record.Put()
func (r *Record) Put() os.Error {
	return putOne(r)
}

// Deletes this record.  Returns nil if successful, otherwise returns the os.Error
// that was retrned by appengime/datastore.
//   record.Delete()
func (r *Record) Delete() os.Error {
	return deleteOne(r)
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

	// does the key have an ID?
	if key.IntID() > 0 {

		// set the ID
		r.setID(key.IntID())

	}

	// set the key
	r.datastoreKey = key

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

// Gets the value of a field in a record.  Strongly typed alternatives are provided and recommended
// to use where possible.
func (r *Record) Get(key string) interface{} {
	
	if r == nil {
		panic(fmt.Sprintf("gaerecords: Cannot Get(\"%v\") property from a nil Record", key))
	}
	
	return r.Fields()[key]
}

// Sets a field in the record.  The value must be an acceptable datastore
// type or another Record.  Strongly typed alternatives are provided and recommended
// to use where possible.
func (r *Record) Set(key string, value interface{}) *Record {
	r.Fields()[key] = value
	return r
}

// Gets a string field
func (r *Record) GetString(key string) string {
	return fmt.Sprint(r.Get(key))
}

// Sets the string value of a field
func (r *Record) SetString(key string, value string) *Record {
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

// Gets the bool value of a field with the specified key.
func (r *Record) GetBool(key string) bool {
	return r.Get(key).(bool)
}

// Sets the bool value of a field with the specified key.
func (r *Record) SetBool(key string, value bool) *Record {
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


/*
	Errors
	----------------------------------------------------------------------
*/

// Causes the record to panic
func (r *Record) panic(message string) {
	
}
