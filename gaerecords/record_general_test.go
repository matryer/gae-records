package gaerecords

import (
	"testing"
)

/*
	Fields
*/

func TestSet(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	// Set() should chain
	assertEqual(t, person, person.Set("name", "Mat"))
	
	// did field update?
	assertEqual(t, "Mat", person.Fields["name"])
	
}

func TestGet(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	person.Fields["age"] = 29
	
	assertEqual(t, 29, person.Get("age"))
	
}

func TestSetIDAndGetID(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	assertEqual(t, int64(-1), person.ID())
	assertEqual(t, person, person.setID(123))
	assertEqual(t, int64(123), person.ID())
	
}

func TestIsPersisted(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	assertEqual(t, false, person.IsPersisted())
	
	person.setID(1)
	
	assertEqual(t, true, person.IsPersisted())
	
}
