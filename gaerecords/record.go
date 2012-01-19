// gaerecords is an Active record style wrapper around Google App Engine datastore
package gaerecords

import (
	"os"
	"appengine/datastore"
)

// A map of the fields of a record
type RecordFields map[string]interface{}

// The key value that indicates there is no ID
var NoIDValue int64 = 0

// Represents a single record
type Record struct {
	
	recordID int64
	parent *Record
	datastoreKey *datastore.Key
	
	// The RecordManager responsible for handling interactions with this record
	Manager *RecordManager
	
	// The fields that make up the data of this record
	// managed via the Get() and Set() methods
	Fields RecordFields
	
}

// Creates a new record
func NewRecord() *Record {
	
	// create the record
	r := new(Record)
	
	// prepare the fields
	r.Fields = make(RecordFields)
	
	// Start off with no ID
	r.recordID = NoIDValue
	
	// return it
	return r
}

/*
	Fields
*/

// Sets a field in the record.  The value must be an acceptable datastore
// type or another Record
func (r *Record) Set(key string, value interface{}) *Record {
	r.Fields[key] = value
	return r
}

// Gets the value of a field in a record
func (r *Record) Get(key string) interface{} {
	return r.Fields[key]
}


/*
	ID Management
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

// Whether this record has been persisted in the
// datastore or not
func (r *Record) IsPersisted() bool {
	return r.recordID != NoIDValue
}

/*
	IO
*/

// Saves the record to the datastore.
func (r *Record) Save() (bool, os.Error) {
	return r.Manager.Save(r)
}



/*
	DatastoreKey
*/

// Gets the datastore key for this record
func (r *Record) DatastoreKey() *datastore.Key {
	
	if r.datastoreKey == nil {
	
		var key *datastore.Key
	
		if r.IsPersisted() {
			key = r.Manager.NewKeyWithID(r.ID())
		} else {
			key = r.Manager.NewKey()
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

/*
	PropertyList
*/

// Creates a datastore.PropertyList containing the fields from the record
func (r *Record) GetFieldsAsPropertyList() datastore.PropertyList {
	
	var list datastore.PropertyList = make(datastore.PropertyList, len(r.Fields))
	var counter int = 0
	
	for k, v := range r.Fields {
		list[counter] = datastore.Property { k, v, false, false }
		counter++
	}
	
	return list
	
}

// Sets the fields in the record to match those of the specified datastore.PropertyList
func (r *Record) SetFieldsFromPropertyList(plist datastore.PropertyList) {
	
	for _, property := range plist {
		r.Fields[property.Name] = property.Value
	}
	
}
