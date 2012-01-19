package gaerecords

import (
	"os"
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

// Creates a new record and initialises it using data from the
// propertyList, and the specified key.
func (m *RecordManager) NewFromPropertyListAndKey(propertyList datastore.PropertyList, key *datastore.Key) *Record {
	
	return m.New().
					SetFieldsFromPropertyList(propertyList).
					SetDatastoreKey(key)
	
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

// Finds a single record by its ID.
func (m *RecordManager) Find(id int64) (*Record, os.Error) {
	
	key := m.NewKeyWithID(id)
	
	var plist datastore.PropertyList
	
	err := datastore.Get(m.appengineContext, key, &plist)
	
	if err != nil {
		
		// return the error
		return nil, err
		
	} else {
		
		// build and return the record
		return m.NewFromPropertyListAndKey(plist, key), nil
		
	}
	
	return nil, nil
	
}

// Saves a record.
func (m *RecordManager) Save(record *Record) (bool, os.Error) {
	
	var plist datastore.PropertyList = record.GetFieldsAsPropertyList()
	newKey, err := datastore.Put(m.appengineContext, record.DatastoreKey(), &plist)
	
	if err != nil {
		return false, err
	} else {
		
		// update the record
		if !record.IsPersisted() {
			record.SetDatastoreKey(newKey)
		}
		
		return true, nil
		
	}
	
	return false, nil
	
}

// Loads all records.
func (m *RecordManager) All() ([]*Record, os.Error) {
	
	query := datastore.NewQuery(m.RecordType())
	
	var plists []*datastore.PropertyList
	_, err := query.GetAll(m.appengineContext, &plists)
	
	if err != nil {
		return nil, err
	} else {
		
		/*
		// make an array to hold these records
		records := make([]*Record, 0, len(plists))
		
		// build each record
		for index, _ := range plists {
			
			// create a new record
			record := m.New()
			
			// set the data
			record.SetFieldsFromPropertyList(plists)
			record.SetDatastoreKey(keys[index])
			
			// save the record
			records = append(records, record)
			
		}
		
		return records, nil
		*/
	}
	
	return nil, nil
	
}