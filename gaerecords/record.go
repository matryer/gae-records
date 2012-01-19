// gaerecords is an Active record style wrapper around Google App Engine datastore
package gaerecords

import (
	"appengine/datastore"
)

// A map of the fields of a record
type RecordFields map[string]interface{}

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
	r := new(Record)
	r.Fields = make(RecordFields)
	r.recordID = -1
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
	
	// invalidate the key
	r.datastoreKey = nil
	
	// chain
	return r
}

// Whether this record has been persisted in the
// datastore or not
func (r *Record) IsPersisted() bool {
	return r.recordID > -1
}

/*
	Parentage
*/

// Sets the parent record of this record.
func (r *Record) SetParent(parent *Record) *Record {
	r.parent = parent
	return r
}

// Gets the parent record of this record or nil if it has no parent.
func (r *Record) Parent() *Record {
	return r.parent
}

// Gets whether this record has a parent or not.  Same as record.Parent() != nil.
func (r *Record) HasParent() bool {
	return r.Parent() != nil
}

/*
	Persistence
*/

// Gets the datastore key for this record
func (r *Record) GetDatastoreKey() *datastore.Key {
	
	if r.datastoreKey == nil {
	
		var key *datastore.Key
		var parentKey *datastore.Key
	
		if r.HasParent() {
			parentKey = r.Parent().GetDatastoreKey()
		}
	
		if r.IsPersisted() {
			key = r.Manager.GetKeyWithID(r.ID(), parentKey)
		} else {
			key = r.Manager.GetKey(parentKey)
		}
	
		r.datastoreKey = key
	
	}
	
	return r.datastoreKey
	
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
