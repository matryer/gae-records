package gaerecords

import (
	"testing"
	"appengine/datastore"
)

func TestSetDatastoreKey(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	var key *datastore.Key = people.NewKeyWithID(123, nil)
	
	person.SetDatastoreKey(key)
	
	// ensure the ID was updated
	assertEqual(t, int64(123), person.ID())
	
}

func TestDatastoreKeyChangingIDInvalidatesCache(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	var key *datastore.Key = person.DatastoreKey()
	
	person.setID(123)
	
	var keyAfterID *datastore.Key = person.DatastoreKey()
	
	if key.Eq(keyAfterID) {
		t.Errorf("Key cache should be invalidated after setID is called")
	}
	
}

func TestDatastoreKeyChangingParent(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	parent := people.New()
	child := people.New()
	
	var key *datastore.Key = child.DatastoreKey()

	if key.Parent() != nil {
		t.Errorf("The Parent() of the key should be nil")
	}

	child.SetParent(parent)
	
	var key2 *datastore.Key = child.DatastoreKey()

	if key2.Parent() == nil {
		t.Errorf("The Parent() of the key should not be nil after .SetParent() is called on the Record")
	}
	
}

func TestDatastoreKeyForPersistedRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New().setID(123)
	
	var key *datastore.Key = person.DatastoreKey()
	
	assertEqual(t, int64(123), key.IntID())
	assertEqual(t, people.RecordType(), key.Kind())
	
}

func TestDatastoreKeyForUnpersistedRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	var key *datastore.Key = person.DatastoreKey()
	
	assertEqual(t, int64(0), key.IntID())
	assertEqual(t, people.RecordType(), key.Kind())
	
}

func TestDatastoreKeyForPersistedRecordWithParentRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	
	parent := people.New().setID(123)
	person := people.New().setID(456).SetParent(parent)
	
	var key *datastore.Key = person.DatastoreKey()
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

func TestDatastoreKeyForUnpersistedRecordWithParentRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	
	parent := people.New().setID(123)
	person := people.New().SetParent(parent)
	
	var key *datastore.Key = person.DatastoreKey()
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
