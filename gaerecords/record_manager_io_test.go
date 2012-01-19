package gaerecords

import (
	"testing"
)

func TestFind(t *testing.T) {
	
	testPerson, _ := CreateTestPerson(t)
	people := CreateTestPeopleRecordManager(t)
	
	assertEqual(t, int64(1), testPerson.ID())
	
	foundPerson := people.Find(1)
	
	if foundPerson == nil {
		
		t.Errorf(".Find(1) should find record")
		
	} else {
		
		assertEqual(t, testPerson.Fields["name"], foundPerson.Fields["name"])
		assertEqual(t, testPerson.Fields["age"], foundPerson.Fields["age"])
		
		assertEqual(t, testPerson.ID(), foundPerson.ID())
		
	}
	
}