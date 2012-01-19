package gaerecords

import (
	"appengine/datastore"
)

type Model struct {
	recordType string
}

func NewModel(recordType string) *Model {
	
	model := new(Model)
	
	model.recordType = recordType
	
	return model
	
}

func (m *Model) New() *Record {
	return NewRecord(m)
}

func (m *Model) RecordType() string {
	return m.recordType
}

/*
	datastore.Keys
	----------------------------------------------------------------------
*/

// Creates a new datastore Key for this kind of record
func (m *Model) NewKey() *datastore.Key {
	return datastore.NewIncompleteKey(GetAppEngineContext(), m.recordType, nil)
}

// Creates a new datastore Key for this kind of record with the specified ID
func (m *Model) NewKeyWithID(id int64) *datastore.Key {
	return datastore.NewKey(GetAppEngineContext(), m.recordType, "", int64(id), nil)
}