package gaerecords

import (
	"os"
	"fmt"
	"appengine"
	"appengine/datastore"
)

// Represents a single model. A model is a class of data.
//  // create a new model for 'people'
//  people := NewModel("people")
type Model struct {

	// Event that gets triggered after a record has been created.
	// Useful for initializing Records.
	//
	//   Args[0] - The *Record that has been created
	AfterNew Event

	// Event that gets triggered after a record of this kind has been
	// found.  Useful for any processing of records after they have been loaded.
	// For any operations that cause multiple finds (i.e. All()) this event
	// will be triggered for each record.
	//
	//   Args[0] - The *Record that has been found.
	//
	AfterFind Event

	// Event that gets triggered before a record is deleted. The Args will
	// always contain the ID of the record being deleted, and depending on the
	// operation, the second argument could be the *Record itself.
	//
	// Setting Cancel to true will cancel the delete operation.
	//
	//   Args[0] - ID (int64) of the record that is about to be deleted.
	//   Args[1] - *Record to be deleted (if available)
	//
	BeforeDelete Event

	// Event that gets triggered after a record has been deleted by ID. The Args will
	// always contain the ID of the record being deleted, and depending on the
	// operation, the second argument could be the *Record itself.
	///
	// The EventContext is the same one that was passed to BeforeDelete.
	//
	//   Args[0] - ID (int64) of the record that was just deleted.
	//   Args[1] - *Record that was deleted (if available)
	//
	AfterDelete Event

	// Event that gets triggered before a record gets Put into the datastore.
	// Use Args[0].(*Record).IsPersisted() to find out whether the record is being
	// saved or updated.
	//
	// Setting Cancel to true will prevent the record from being Put
	// 
	//   Args[0] - The *Record that is about to be Put
	//
	BeforePut Event

	// Event that gets triggered after a record has been Put.
	// The EventContext is the same one that was passed to BeforePut.
	//
	//   Args[0] - The *Record that was just Put
	// 
	AfterPut Event

	// Event that gets triggered after a record field has been changed
	// using one of the Set*() methods.
	//
	//   Args[0] - The record that changed
	//   Args[1] - The key of the field that changed
	//   Args[2] - The new value of the field
	//   Args[3] - The old value of the field (or nil if it's a new field)
	OnChanged Event

	// internal string holding the 'type' of this model,
	// or the kind of data this model works with
	recordType string
	
	// internal storage of appengine context to use for this model.
	specificAppengineContext appengine.Context
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
	AppEngine Context
	----------------------------------------------------------------------
*/

func (m *Model) AppEngineContext() appengine.Context {
	if m.specificAppengineContext == nil {
		m.specificAppengineContext = AppEngineContext
	}
	return m.specificAppengineContext
}

func (m *Model) SetAppEngineContext(context appengine.Context) *Model {
	m.specificAppengineContext = context
	
	// chain
	return m
}

func (m *Model) UseGlobalAppEngineContext() *Model {
	
	m.SetAppEngineContext(nil)
	
	// chain
	return m
}

/*
	Persistence
	----------------------------------------------------------------------
*/

// Finds the record of this type with the specified id.
//  people := NewModel("people")
//  firstPerson := people.Find(1)
//
// Raises events:
//   Model.AfterFind with Args(record)
func (m *Model) Find(id int64) (*Record, os.Error) {
	return findOneByID(m, id)
}

// Finds all records of this type.
//   people := NewModel("people")
//   everyone := people.All()
//
// Raises events for each record:
//   Model.AfterFind with Args(record)
func (m *Model) All() ([]*Record, os.Error) {
	return findAll(m)
}

// Deletes a single record of this type.  Returns nil if successful, otherwise
// the datastore error that was returned.
//   people := NewModel("people")
//   people.Delete(1)
//
// Raises events:
//   Model.BeforeDelete with Args(id, nil)
//   Model.AfterDelete with Args(id, nil)
// Note: The Record will not be passed to the events.
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
