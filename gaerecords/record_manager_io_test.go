package gaerecords

import (
	"testing"
)

func TestFind(t *testing.T) {
	
	_, err := CreateTestPerson(t)
	//people := CreateTestPeopleRecordManager()
	
	assertEqual(t, "", err.String())
	//assertEqual(t, -100, testPerson.ID())
	
}