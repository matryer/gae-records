package gaerecords

import (
	"os"
	"appengine/datastore"
)

func findAll(model *Model) ([]*Record, os.Error) {

	query := datastore.NewQuery(model.RecordType())

	var records []*Record
	keys, err := query.GetAll(GetAppEngineContext(), &records)

	if err == nil {

		// update the key for each loaded record
		for index, record := range records {
			record.SetModel(model).SetDatastoreKey(keys[index])
			model.AfterFind.Trigger(record)
		}

		return records, nil

	}

	return nil, err

}

func findOneByID(model *Model, id int64) (*Record, os.Error) {

	key := model.NewKeyWithID(id)

	var record *Record = NewRecord(model)

	err := datastore.Get(GetAppEngineContext(), key, datastore.PropertyLoadSaver(record))

	if err == nil {

		// set the key
		record.SetDatastoreKey(key)
		
		// raise the AfterFind event on the model
		model.AfterFind.Trigger(record)
		
		// return the record
		return record, nil

	}

	return nil, err

}

func deleteOne(record *Record) os.Error {

	err := datastore.Delete(GetAppEngineContext(), record.DatastoreKey())

	if err == nil {

		// clean up the record
		record.setID(NoIDValue)

	}

	return err

}

func deleteOneByID(model *Model, id int64) os.Error {
	
	// trigger the BeforeDeleteByID event
	context := model.BeforeDeleteByID.Trigger(id)
	
	if !context.Cancel {
		
		err := datastore.Delete(GetAppEngineContext(), model.NewKeyWithID(id))
		
		if err == nil {
			
			// trigger the AfterDeleteByID event
			model.AfterDeleteByID.TriggerWithContext(context)
			
		}
		
		// return the error
		return err
		
	}
	
	return ErrOperationCancelledByEventCallback
	
}

func putOne(record *Record) os.Error {

	// trigger the BeforePut event on the model
	context := record.Model().BeforePut.Trigger(record)

	if !context.Cancel {
		
		newKey, err := datastore.Put(GetAppEngineContext(), record.DatastoreKey(), datastore.PropertyLoadSaver(record))
		
		if err == nil {

			// update the record key
			record.SetDatastoreKey(newKey)
			
			// trigger the AfterPut event
			record.Model().AfterPut.TriggerWithContext(context)

			return nil

		}
		
		return err
		
	}

	return ErrOperationCancelledByEventCallback

}
