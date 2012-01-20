package gaerecords

import (
	"os"
	"appengine/datastore"
)

func LoadOneByID(model *Model, id int64) (*Record, os.Error) {
	
	key := model.NewKeyWithID(id)
	
	var record *Record = NewRecord(model)
	
	err := datastore.Get(GetAppEngineContext(), key, datastore.PropertyLoadSaver(record))
	
	if err == nil {
	
		// build and return the record
		return record.SetDatastoreKey(key), nil
	
	}

	return nil, err
	
}

func PutOne(record *Record) os.Error {
	
	newKey, err := datastore.Put(GetAppEngineContext(), record.DatastoreKey(), datastore.PropertyLoadSaver(record))
	
	if err == nil {
		
		// update the record key
		record.SetDatastoreKey(newKey)
		
		return nil
		
	}
	
	return err
	
}
