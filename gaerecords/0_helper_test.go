package gaerecords

import (
	"testing"
	"os"
	"appengine/datastore"
	"gae-go-testing.googlecode.com/git/appenginetesting"
)

// Asserts that two objects are equal and throws an error if not
func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected to be equal. %v != %v", a, b)
	}
}

// Creates a test appengine context object
func CreateTestAppengineContext() *appenginetesting.Context {
	appengineContext, _ := appenginetesting.NewContext(nil)
	return appengineContext
}

// Creates a test 'people' record manager
func CreateTestPeopleRecordManager() *RecordManager {
	return NewRecordManager(CreateTestAppengineContext(), "people")
}

// Resets the datastore to a default test position
func CreateTestPerson() (*Record, os.Error) {
	
	context := CreateTestAppengineContext()
	people := NewRecordManager(context, "people")
	person := people.New()
	key := person.GetDatastoreKey()
	
	newKey, err := datastore.Put(context, key, person.Fields)
	
	if err != nil {
		return nil, err
	}
	
	// set the person ID
	person.setID(newKey.IntID())
	
	return person, nil
	
}