package gaerecords

import (
	"appengine"
	"appengine/datastore"
)

// Represents a kind of records that can be interacted with
type RecordManager struct {
	recordType string
	appengineContext appengine.Context
}

/*
	Creates a new record manager responsible for managing
	the specified type of data
*/
func NewRecordManager(context appengine.Context, recordType string) *RecordManager {
	m := new(RecordManager)
	m.recordType = recordType
	m.appengineContext = context
	return m
}

// Creates a new record managed by this manager
func (m *RecordManager) New() *Record {
	r := NewRecord()
	r.Manager = m
	return r
}

// Gets a string representing the record type managed by this manager
func (m *RecordManager) RecordType() string {
	return m.recordType
}

/*
	Keys
*/

// Gets a datastore Key for this kind of record
func (m *RecordManager) GetKey(parent *datastore.Key) *datastore.Key {
	return datastore.NewIncompleteKey(m.appengineContext, m.RecordType(), parent)
}

// Gets a datastore Key for this kind of record with the specified ID
func (m *RecordManager) GetKeyWithID(id int64, parent *datastore.Key) *datastore.Key {
	return datastore.NewKey(m.appengineContext, m.RecordType(), "", int64(id), parent)
}

/*
	Data access
*/

// Finds a single record by its ID
func (m *RecordManager) Find(id int64) *Record {
	
	key := m.GetKeyWithID(id, nil)
	
	var plist datastore.PropertyList
	
	err := datastore.Get(m.appengineContext, key, &plist)
	
	if err != nil {
		
		// TODO: handle error
		
	} else {
		
		// build the record
		record := m.New()
		record.SetFieldsFromPropertyList(plist)
		
		return record
		
	}
	
	return nil
}