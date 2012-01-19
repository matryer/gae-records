package gaerecords

import (
	"appengine/datastore"
)

type RecordFields map[string]interface{}

type RecordID int64

/*
	Record
	------------------------------------------------------------
*/
type Record struct {
	
	recordID RecordID
	parent *Record
	
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
func (r *Record) ID() RecordID {
	return r.recordID
}

// Sets the ID for this record
func (r *Record) setID(id RecordID) *Record {
	r.recordID = id
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

func (r *Record) GetDatastoreKey() *datastore.Key {
	
	var key *datastore.Key
	var parentKey *datastore.Key
	
	if r.HasParent() {
		parentKey = r.GetDatastoreKey()
	}
	
	if r.IsPersisted() {
		key = datastore.NewKey(r.Manager.appengineContext, r.Manager.RecordType(), "", int64(r.ID()), parentKey)
	} else {
		key = datastore.NewIncompleteKey(r.Manager.appengineContext, r.Manager.RecordType(), parentKey)
	}
	
	return key
	
}

