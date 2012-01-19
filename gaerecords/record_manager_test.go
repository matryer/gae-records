package gaerecords

import (
	"testing"
	"appengine/datastore"
)

func TestNewRecordManager(t *testing.T) {

	m := NewRecordManager(TestContext, "something")
	
	assertEqual(t, "something", m.RecordType())
	assertEqual(t, TestContext, m.appengineContext)
	
}

func TestNew(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	assertEqual(t, people, person.Manager)
	
}

func TestNewFromPropertyListAndKey(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	
	// build a test property list
	var plist datastore.PropertyList = make(datastore.PropertyList, 3)
	plist[0] = datastore.Property{ "name", "Mat", false, false }
	plist[1] = datastore.Property{ "age", int64(29), false, false }
	plist[2] = datastore.Property{ "is_dev", true, false, false }
	
	// create a key
	key := people.NewKeyWithID(246)
	
	person := people.NewFromPropertyListAndKey(plist, key)
	
	if person == nil {
		
		t.Errorf("NewFromPropertyListAndKey should return valid Record")
		
	} else {
		
		assertEqual(t, int64(246), person.ID())
		assertEqual(t, "Mat", person.GetString("name"))
		assertEqual(t, int64(29), person.GetInt("age"))
		assertEqual(t, true, person.GetBool("is_dev"))
		
	}
	
}

func TestRecordType(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	assertEqual(t, "people", people.RecordType())
	
}

