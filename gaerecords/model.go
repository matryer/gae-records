package gaerecords

import (
	"os"
	"fmt"
	"appengine/datastore"
)

// Represents a single model. A model is a class of data.
//  // create a new model for 'people'
//  people := NewModel("people")
type Model struct {
	recordType string
}

// Creates a new model for data classified by the specified recordType.
// 
// For example, the following code creates a new Model called 'people':
//
//   people := NewModel("people")
func NewModel(recordType string) *Model {

	model := new(Model)

	model.recordType = recordType

	return model

}

// Creates a new record of this type.
//   people := NewModel("people")
//   person1 := people.New()
//   person2 := people.New()
func (m *Model) New() *Record {
	return NewRecord(m)
}

// Gets the record type of the model as a string.  This is the string you specify
// when calling NewModel(string) and is used as the Kind in the datasource keys.
func (m *Model) RecordType() string {
	return m.recordType
}

// Gets a human readable string representation of this model.
func (m *Model) String() string {
	return fmt.Sprintf("{Model:%v}", m.RecordType())
}

/*
	Persistence
	----------------------------------------------------------------------
*/

// Finds the record of this type with the specified id.
//  people := NewModel("people")
//  firstPerson := people.Find(1)
func (m *Model) Find(id int64) (*Record, os.Error) {
	return findOneByID(m, id)
}

// Finds all records of this type.
//   people := NewModel("people")
//   everyone := people.All()
func (m *Model) All() ([]*Record, os.Error) {
	return findAll(m)
}

// Deletes a single record of this type.  Returns nil if successful, otherwise
// the datastore error that was returned.
//   people := NewModel("people")
//   people.Delete(1)
func (m *Model) Delete(id int64) os.Error {
	return deleteOneByID(m, id)
}

/*
	datastore.Keys
	----------------------------------------------------------------------
*/

// Creates a new datastore Key for this kind of record.
func (m *Model) NewKey() *datastore.Key {
	return datastore.NewIncompleteKey(GetAppEngineContext(), m.recordType, nil)
}

// Creates a new datastore Key for this kind of record with the specified ID.
func (m *Model) NewKeyWithID(id int64) *datastore.Key {
	return datastore.NewKey(GetAppEngineContext(), m.recordType, "", int64(id), nil)
}
