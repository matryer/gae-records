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
	cachedKey *datastore.Key
	
	Manager *RecordManager
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

// Sets a field in the record
func (r *Record) Set(k string, v interface{}) *Record {
	r.Fields[k] = v
	return r
}

// Gets the value of a field in a record
func (r *Record) Get(k string) interface{} {
	return r.Fields[k]
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
	r.cachedKey = nil
	
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

func (r *Record) SetParent(parent *Record) *Record {
	r.parent = parent
	return r
}

func (r *Record) Parent() *Record {
	return r.parent
}

// TODO: test me
func (r *Record) HasParent() bool {
	return r.Parent() != nil
}

/*
	Persistence
*/

// Gets the datastore key for this record
func (r *Record) GetDatastoreKey() *datastore.Key {
	
	if r.cachedKey == nil {
	
		var key *datastore.Key
		var parentKey *datastore.Key
	
		if r.HasParent() {
			parentKey = r.Parent().GetDatastoreKey()
		}
	
		if r.IsPersisted() {
			key = datastore.NewKey(r.Manager.appengineContext, r.Manager.RecordType(), "", int64(r.ID()), parentKey)
		} else {
			key = datastore.NewIncompleteKey(r.Manager.appengineContext, r.Manager.RecordType(), parentKey)
		}
	
		r.cachedKey = key
	
	}
	
	return r.cachedKey
	
}

/*
	PropertyList
*/

func (r *Record) GetFieldsAsPropertyList() datastore.PropertyList {
	
	var list datastore.PropertyList = make(datastore.PropertyList, len(r.Fields))
	var counter int = 0
	
	for k, v := range r.Fields {
		list[counter] = datastore.Property { k, v, true, false }
		counter++
	}
	
	return list
	
}