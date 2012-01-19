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

// Creates a new datastore Key for this kind of record
func (m *RecordManager) NewKey() *datastore.Key {
	return datastore.NewIncompleteKey(m.appengineContext, m.RecordType(), nil)
}

// Creates a new datastore Key for this kind of record with the specified ID
func (m *RecordManager) NewKeyWithID(id int64) *datastore.Key {
	return datastore.NewKey(m.appengineContext, m.RecordType(), "", int64(id), nil)
}

/*
	Data access
*/

// Finds a single record by its ID
func (m *RecordManager) Find(id int64) *Record {
	
	key := m.NewKeyWithID(id)
	
	var plist datastore.PropertyList
	
	err := datastore.Get(m.appengineContext, key, &plist)
	
	if err != nil {
		
		// TODO: handle error
		
	} else {
		
		// build the record
		record := m.New()
		
		record.SetFieldsFromPropertyList(plist)
		record.SetDatastoreKey(key)
		
		return record
		
	}
	
	return nil
}