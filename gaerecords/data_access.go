package gaerecords

import (
	"os"
	"reflect"
	"appengine/datastore"
)

// FindRecordByKey finds a single record for the given model by the specified datastore.Key.
func FindRecordByKey(model *Model, key *datastore.Key) (*Record, os.Error) {
	
	var record *Record = new(Record)

	err := datastore.Get(model.AppEngineContext(), key, datastore.PropertyLoadSaver(record))

	if err == nil {

		// setup the record object
		record.configureRecord(model, key)

		// raise the AfterFind event on the model
		if model.AfterFind.HasCallbacks() {
			model.AfterFind.Trigger(record)
		}

		// return the record
		return record, nil

	}

	return nil, err
	
}

// FindRecord finds a single record for the specified model with the specified ID.
func FindRecord(model *Model, id int64) (*Record, os.Error) {

	key := model.NewKeyWithID(id)
	return FindRecordByKey(model, key)

}

// FindRecordWithParent finds a single child record of parent, for the specified model, with the specified ID.
//
// If you have the parent, call parent.Find(model, id) directly.
func FindRecordWithParent(model *Model, id int64, parent *Record) (*Record, os.Error) {

	key := model.NewKeyWithIDAndParent(id, parent.DatastoreKey())
	return FindRecordByKey(model, key)
	
}

// FindByQuery finds Records handled by the specified model.
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
//  women, err := FindByQuery(model, func(q *datastore.Query){
//	  q.Filter("sex=", "male")
//  })
//
// Raises events for each record:
//   Model.AfterFind with Args(record)
//
// It is cleaner to call Model.FindByQuery instead.
func FindByQuery(model *Model, queryOrFunc interface{}) ([]*Record, os.Error) {

	var query *datastore.Query

	if reflect.TypeOf(queryOrFunc).Kind() == reflect.Func {

		// create a new query
		query = model.NewQuery()

		// ask the func to configure the query
		queryOrFunc.(func(*datastore.Query))(query)

	} else {

		// just use the query
		query = queryOrFunc.(*datastore.Query)

	}

	var records []*Record
	keys, err := query.GetAll(model.AppEngineContext(), &records)

	if err == nil {

		// configure each loaded record
		for index, record := range records {

			// configure the record
			record.configureRecord(model, keys[index])

			if model.AfterFind.HasCallbacks() {
				model.AfterFind.Trigger(record)
			}

		}

		return records, nil

	}

	return nil, err

}