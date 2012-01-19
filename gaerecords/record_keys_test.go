package gaerecords

import (
	"testing"
	"appengine/datastore"
)

func TestGetDatastoreKeyForPersistedRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	person := people.New().setID(123)
	
	var key *datastore.Key = person.GetDatastoreKey()
	
	assertEqual(t, int64(123), key.IntID())
	assertEqual(t, people.RecordType(), key.Kind())
	
}

func TestGetDatastoreKeyForUnpersistedRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	person := people.New()
	
	var key *datastore.Key = person.GetDatastoreKey()
	
	assertEqual(t, int64(0), key.IntID())
	assertEqual(t, people.RecordType(), key.Kind())
	
}

func TestGetDatastoreKeyForPersistedRecordWithParentRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	
	parent := people.New().setID(123)
	person := people.New().setID(456).SetParent(parent)
	
	var key *datastore.Key = person.GetDatastoreKey()
	parentKey := key.Parent()
	
	if parentKey == nil {

		t.Errorf("key.Parent() shouldn't be nil")
		
	} else {
	
		assertEqual(t, int64(123), parentKey.IntID())
		assertEqual(t, people.RecordType(), parentKey.Kind())
	
	}
	
	assertEqual(t, int64(456), key.IntID())
	assertEqual(t, people.RecordType(), key.Kind())
	
}

func TestGetDatastoreKeyForUnpersistedRecordWithParentRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	
	parent := people.New().setID(123)
	person := people.New().SetParent(parent)
	
	var key *datastore.Key = person.GetDatastoreKey()
	parentKey := key.Parent()
	
	if parentKey == nil {

		t.Errorf("key.Parent() shouldn't be nil")
		
	} else {
	
		assertEqual(t, int64(123), parentKey.IntID())
		assertEqual(t, people.RecordType(), parentKey.Kind())
	
	}
	
	assertEqual(t, int64(0), key.IntID())
	assertEqual(t, people.RecordType(), key.Kind())
	
}
