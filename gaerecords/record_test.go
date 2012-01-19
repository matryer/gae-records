package gaerecords

import (
	"testing"
	"appengine/datastore"
)

/*
	Relationships
*/
func TestSetParentAndParent(t *testing.T) {

	people := CreateTestPeopleRecordManager()
	parent := people.New()

	child := people.New()
	
	assertEqual(t, child, child.SetParent(parent))
	assertEqual(t, parent, child.Parent())

}

/*
	Fields
*/

func TestSet(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	person := people.New()
	
	// Set() should chain
	assertEqual(t, person, person.Set("name", "Mat"))
	
	// did field update?
	assertEqual(t, "Mat", person.Fields["name"])
	
}

func TestGet(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	person := people.New()
	
	person.Fields["age"] = 29
	
	assertEqual(t, 29, person.Get("age"))
	
}

func TestSetIDAndGetID(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	person := people.New()
	
	assertEqual(t, RecordID(-1), person.ID())
	assertEqual(t, person, person.setID(123))
	assertEqual(t, RecordID(123), person.ID())
	
}

func TestIsPersisted(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	person := people.New()
	
	assertEqual(t, false, person.IsPersisted())
	
	person.setID(1)
	
	assertEqual(t, true, person.IsPersisted())
	
}

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
