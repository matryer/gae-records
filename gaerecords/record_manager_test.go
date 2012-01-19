package gaerecords

import (
	"testing"
)

func TestNewRecordManager(t *testing.T) {

	context := CreateTestAppengineContext()
	m := NewRecordManager(context, "something")
	
	assertEqual(t, "something", m.RecordType())
	assertEqual(t, context, m.appengineContext)
	
}

func TestNew(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	person := people.New()
	
	assertEqual(t, people, person.Manager)
	
}

func TestRecordType(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	assertEqual(t, "people", people.RecordType())
	
}

