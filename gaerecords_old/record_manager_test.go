package gaerecords

import (
	"testing"
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

func TestRecordType(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	assertEqual(t, "people", people.RecordType())
	
}

