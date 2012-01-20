package gaerecords

import (
	"testing"
	"appengine/datastore"
)

func TestSetDatastoreKey(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	var key *datastore.Key = people.NewKeyWithID(123)

	person.SetDatastoreKey(key)

	// ensure the ID was updated
	assertEqual(t, int64(123), person.ID())

}

func TestDatastoreKeyChangingIDInvalidatesCache(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	var key *datastore.Key = person.DatastoreKey()

	person.setID(123)

	var keyAfterID *datastore.Key = person.DatastoreKey()

	if key.Eq(keyAfterID) {
		t.Errorf("Key cache should be invalidated after setID is called")
	}

}

func TestDatastoreKeyForPersistedRecord(t *testing.T) {

	people := CreateTestModelWithPropertyType("modeltwo")
	person := people.New().setID(123)

	var key *datastore.Key = person.DatastoreKey()

	assertEqual(t, int64(123), key.IntID())
	assertEqual(t, people.RecordType(), key.Kind())

}

func TestDatastoreKeyForUnpersistedRecord(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	var key *datastore.Key = person.DatastoreKey()

	assertEqual(t, int64(0), key.IntID())
	assertEqual(t, people.RecordType(), key.Kind())

}

func TestIsPersisted(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	assertEqual(t, false, person.IsPersisted())

	person.setID(1)

	assertEqual(t, true, person.IsPersisted())

}
