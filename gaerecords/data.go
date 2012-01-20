package gaerecords

import (
	"os"
	"appengine/datastore"
)

func LoadOneByID(model *Model, id int64) (*Record, os.Error) {
	
	key := model.NewKeyWithID(id)
	
	var record *Record = NewRecord(model)
	
	err := datastore.Get(GetAppEngineContext(), key, datastore.PropertyLoadSaver(record))
	
	if err != nil {
	
		// return the error
		return nil, err
	
	} else {
	
		// build and return the record
		return record.SetDatastoreKey(key), nil
	
	}

	return nil, nil
	
}