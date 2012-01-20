package gaerecords

import (
	"os"
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
	Persistence
	----------------------------------------------------------------------
*/

func (m *Model) Find(id int64) (*Record, os.Error) {
	return FindOneByID(m, id)
}

func (m *Model) All() ([]*Record, os.Error) {
	return FindAll(m)
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