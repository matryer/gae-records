package gaerecords

import (
	"os"
	"fmt"
	"appengine/datastore"
)

// The key value that indicates there is no ID
var NoIDValue int64 = 0

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
func NewRecord(model *Model) *Record {
	
	record := new(Record)
	record.fields = make(map[string]interface{})
	record.model = model
	return record
	
}

/*
	Properties
	----------------------------------------------------------------------
*/
func (r *Record) Model() *Model {
	return r.model
}

/*
	IDs
	----------------------------------------------------------------------
*/

// Gets the ID for this record
func (r *Record) ID() int64 {
	return r.recordID
}

// Sets the ID for this record
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

func (r *Record) Load(c <-chan datastore.Property) os.Error {
		
	// load the fields
	for f := range c {
		r.fields[f.Name] = f.Value
	}
	
	// no errors
	return nil
}

func (r *Record) Save(c chan<- datastore.Property) os.Error {

	for k, v := range r.fields {
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


/*
	Datastore Key
	----------------------------------------------------------------------
*/

// Gets the datastore key for this record
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
// datastore or not
func (r *Record) IsPersisted() bool {
	return r.recordID != NoIDValue
}

/*
	Fields
	----------------------------------------------------------------------
*/

func (r *Record) Fields() map[string]interface{} {
	return r.fields
}

// Gets the value of a field in a record
func (r *Record) Get(key string) interface{} {
	return r.fields[key]
}
	
// Sets a field in the record.  The value must be an acceptable datastore
// type or another Record
func (r *Record) Set(key string, value interface{}) *Record {
	r.fields[key] = value
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

// Gets an int64 field
func (r *Record) GetInt(key string) int64 {
	return r.Get(key).(int64)
}

// Sets the int64 value of a field
func (r *Record) SetInt(key string, value int64) *Record {
	return r.Set(key, value)
}

// Gets a bool field
func (r *Record) GetBool(key string) bool {
	return r.Get(key).(bool)
}

// Sets the bool value of a field
func (r *Record) SetBool(key string, value bool) *Record {
	return r.Set(key, value)
}

// Gets a key field
func (r *Record) GetKeyField(key string) *datastore.Key {
	return r.Get(key).(*datastore.Key)
}

// Sets the key value of a field
func (r *Record) SetKeyField(key string, value *datastore.Key) *Record {
	return r.Set(key, value)
}