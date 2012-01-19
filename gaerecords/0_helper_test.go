package gaerecords

import (
	"testing"
	"os"
	"appengine"
	"appengine/datastore"
	"gae-go-testing.googlecode.com/git/appenginetesting"
)

// Asserts that two objects are equal and throws an error if not
func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected to be equal. %v != %v", a, b)
	}
}

var TestContext *appenginetesting.Context

// Creates a test appengine context object
func CreateTestAppengineContext(t *testing.T) appengine.Context {
	
	if TestContext == nil {
		t.Logf("<<< Test context created >>>")
		TestContext, _ = appenginetesting.NewContext(nil)
	}
	
	return TestContext
	
}

// Creates a test 'people' record manager
func CreateTestPeopleRecordManager(t *testing.T) *RecordManager {
	return NewRecordManager(CreateTestAppengineContext(t), "people")
}

// Resets the datastore to a default test position
func CreateTestPerson(t *testing.T) (*Record, os.Error) {
	
	context := CreateTestAppengineContext(t)
	people := NewRecordManager(context, "people")
	person := people.New()
	key := person.DatastoreKey()
	
	person.Set("name", "Mat").Set("age", int64(29))
	
	var plist datastore.PropertyList = person.GetFieldsAsPropertyList()
	
	newKey, err := datastore.Put(context, key, &plist)
	
	if err != nil {
		return nil, err
	}
	
	// set the person ID
	person.setID(newKey.IntID())
	
	return person, nil
	
}
