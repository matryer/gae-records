package gaerecords

import (
	"os"
	"fmt"
	"reflect"
	"appengine"
	"appengine/datastore"
)

// Model struct represents a single model. A model is a class of data and a Model object is used
// to interact with the datastore including reading and writing records of this type.
//
// The NewModel method creates a new model. 
// The argument specified is a string description of the type of data, which is used
// in the *datastore.Key for records of this kind.
//  // create a new model for 'people'
//  People := NewModel("people")
//
//  // create a new model for 'books'
//  Books := NewModel("books")
type Model struct {

	// AfterNew gets triggered after a record has been created.
	// Useful for initializing Records.
	//
	//   Args[0] - The *Record that has been created
	//
	AfterNew Event

	// AfterFind gets triggered after a record of this kind has been
	// found.  Useful for any processing of records after they have been loaded.
	// For any operations that cause multiple finds (i.e. All() or FindByQuery()) this event
	// will be triggered once for each record.
	//
	//   Args[0] - The *Record that has been found.
	//
	AfterFind Event

	// BeforeDelete gets triggered before a record is deleted. The Args will
	// always contain the ID of the record being deleted, and depending on the
	// operation, the second argument could be the *Record itself.
	//
	// Setting Cancel to true will cancel the delete operation.
	//
	//   Args[0] - ID (int64) of the record that is about to be deleted.
	//   Args[1] - *Record to be deleted (if available)
	//
	BeforeDelete Event

	// AfterDelete gets triggered after a record has been deleted by ID. The Args will
	// always contain the ID of the record being deleted, and depending on the
	// operation, the second argument could be the *Record itself.
	///
	// The EventContext is the same one that was passed to BeforeDelete.
	//
	//   Args[0] - ID (int64) of the record that was just deleted.
	//   Args[1] - *Record that was deleted (if available)
	//
	AfterDelete Event

	// BeforePut gets triggered before a record gets Put into the datastore.
	// Use Args[0].(*Record).IsPersisted() to find out whether the record is being
	// saved or updated.
	//
	// Setting Cancel to true will prevent the record from being Put
	// 
	//   Args[0] - The *Record that is about to be Put
	//
	BeforePut Event

	// AfterPut gets triggered after a record has been Put.
	// The EventContext is the same one that was passed to BeforePut.
	//
	//   Args[0] - The *Record that was just Put
	// 
	AfterPut Event

	// OnChanged gets triggered after a record field has been changed
	// using one of the Set*() methods.
	//
	//   Args[0] - The record that changed
	//   Args[1] - The key of the field that changed
	//   Args[2] - The new value of the field
	//   Args[3] - The old value of the field (or nil if it's a new field)
	//
	OnChanged Event

	// recordType is the internal string holding the 'type' of this model,
	// or the kind of data this model works with
	recordType string

	// specificAppengineContext is the internal storage of appengine context to use for this model.
	specificAppengineContext appengine.Context
}

// NewModel creates a new model for data classified by the specified recordType.
// 
// For example, the following code creates a new Model called 'people':
//
//   people := NewModel("people")
func NewModel(recordType string) *Model {

	model := new(Model)

	model.recordType = recordType

	return model

}

// New creates a new record of this type.
//   people := NewModel("people")
//   person1 := people.New()
//   person2 := people.New()
func (m *Model) New() *Record {
	return NewRecord(m)
}

// RecordType gets the record type of the model as a string.  This is the string you specify
// when calling NewModel(string) and is used as the Kind in the datasource keys.
func (m *Model) RecordType() string {
	return m.recordType
}

// String gets a human readable string representation of this model.
func (m *Model) String() string {
	return fmt.Sprintf("{Model:%v}", m.RecordType())
}

/*
	AppEngine Context
	----------------------------------------------------------------------
*/

// AppEngineContext gets the appengine.Context to use for datastore interactions for this model.
// If a specific one has been provided (via Model.SetAppEngineContext()) that 
// context is used, otherwise the global AppEngineContext object is returned.
func (m *Model) AppEngineContext() appengine.Context {

	// do we have a specific model context?
	if m.specificAppengineContext == nil {

		// use the global one
		return AppEngineContext

	}

	// use the specific one
	return m.specificAppengineContext

}

// SetAppEngineContext tells this model to use the specified appengine.Context instead of the global
// AppEngineContext object for its interactions with the datastore.
func (m *Model) SetAppEngineContext(context appengine.Context) *Model {

	// set the context
	m.specificAppengineContext = context

	// chain
	return m
}

// UseGlobalAppEngineContext tells this model to use the global AppEngineContext object for its interactions with the datastore, 
// instead of one provided by Model.SetAppEngineContext().
func (m *Model) UseGlobalAppEngineContext() *Model {

	// set the model specific context to nil so it uses the
	// global one when Model.AppEngineContext() is called.
	m.SetAppEngineContext(nil)

	// chain
	return m
}

/*
	Persistence
	----------------------------------------------------------------------
*/

// Find finds the record of this type with the specified id.
//  people := NewModel("people")
//  firstPerson := people.Find(1)
//
// Raises events:
//   Model.AfterFind with Args(record)
func (m *Model) Find(id int64) (*Record, os.Error) {

	key := m.NewKeyWithID(id)

	var record *Record = new(Record)

	err := datastore.Get(m.AppEngineContext(), key, datastore.PropertyLoadSaver(record))

	if err == nil {

		// setup the record object
		record.configureRecord(m, key)

		// raise the AfterFind event on the model
		m.AfterFind.Trigger(record)

		// return the record
		return record, nil

	}

	return nil, err

}

// FindAll finds all records of this type.
//   people := NewModel("people")
//   everyone := people.All()
//
// Raises events for each record:
//   Model.AfterFind with Args(record)
func (m *Model) FindAll() ([]*Record, os.Error) {
	return m.FindByQuery(m.NewQuery())
}

// FindByField finds records of this Model's type where the filterString (e.g. FieldName=) matches
// the specified value.  To add multiple filters, use FindByQuery() instead.
//
// You can pass an optional query modifier func (of type func(*datastore.Query)) that will be called
// before the query is run to allow you to add additional properties to the query.
//
// For valid filter strings, see http://code.google.com/appengine/docs/go/datastore/reference.html#Query.Filter
//
//  // get everyone over 60 years old
//  oldPeople, _ := People.FindByField("Age>", 60)
//
//  // get one person over 60
//  oldMan, _ := People.FindByField("Age>", 60, func(q *datastore.Query){
//	  q.Limit(1)
//  })
func (m *Model) FindByField(filterString string, value interface{}, queryModifier ...func(*datastore.Query)) ([]*Record, os.Error) {

	query := m.NewQuery().Filter(filterString, value)

	// let the modifier do its work if there is one
	if len(queryModifier) == 1 {
		queryModifier[0](query)
	}

	return m.FindByQuery(query)

}

/*
	Queries
	----------------------------------------------------------------------
*/

// NewQuery creates a new datastore.Query for accessing records represented
// by the model.  For advanced use only.  Consider instead one of the 
// Find* methods.
func (m *Model) NewQuery() *datastore.Query {
	return datastore.NewQuery(m.RecordType())
}

// FindByQuery finds Records handled by this Model.
//
// Returns an array of records as the first argument,
// or an error as the second return argument.
//
// The queryOrFunc argument may be one of:
//
//   *datastore.Query
// The specified query will be used to find records.
//   func(*datastore.Query)
// A new query will be created and the specified function will be
// used to further configure the query.
//
// Example:
//  model := NewModel("people")
//  women, err := model.FindByQuery(func(q *datastore.Query){
//	  q.Filter("sex=", "male")
//  })
//
// bug(matryer): Currently, filters on queries are not working. See https://github.com/matryer/gae-records/issues/2
func (m *Model) FindByQuery(queryOrFunc interface{}) ([]*Record, os.Error) {

	var query *datastore.Query

	if reflect.TypeOf(queryOrFunc).Kind() == reflect.Func {

		// create a new query
		query = m.NewQuery()

		// ask the func to configure the query
		queryOrFunc.(func(*datastore.Query))(query)

	} else {

		// just use the query
		query = queryOrFunc.(*datastore.Query)

	}

	var records []*Record
	keys, err := query.GetAll(m.AppEngineContext(), &records)

	if err == nil {

		// update the key for each loaded record
		for index, record := range records {
			record.configureRecord(m, keys[index])
			m.AfterFind.Trigger(record)
		}

		return records, nil

	}

	return nil, err

}

// Delete deletes a single record of this type.  Returns nil if successful, otherwise
// the datastore error that was returned.
//   people := NewModel("people")
//   people.Delete(1)
//
// Raises events:
//   Model.BeforeDelete with Args(id, nil)
//   Model.AfterDelete with Args(id, nil)
// Note: The Record will not be passed to the events.
func (m *Model) Delete(id int64) os.Error {

	// trigger the BeforeDeleteByID event
	context := m.BeforeDelete.Trigger(id, nil)

	if !context.Cancel {

		err := datastore.Delete(m.AppEngineContext(), m.NewKeyWithID(id))

		if err == nil {

			// trigger the AfterDeleteByID event
			m.AfterDelete.TriggerWithContext(context)

		}

		// return the error
		return err

	}

	return ErrOperationCancelledByEventCallback

}

/*
	datastore.Keys
	----------------------------------------------------------------------
*/

// NewKey creates a new datastore Key for this kind of record.
func (m *Model) NewKey() *datastore.Key {
	return datastore.NewIncompleteKey(m.AppEngineContext(), m.recordType, nil)
}

// NewKeyWithID creates a new datastore Key for this kind of record with the specified ID.
func (m *Model) NewKeyWithID(id int64) *datastore.Key {
	return datastore.NewKey(m.AppEngineContext(), m.recordType, "", int64(id), nil)
}
