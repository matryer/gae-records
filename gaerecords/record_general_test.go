package gaerecords

import (
	"testing"
	"appengine/datastore"
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

func TestGetAndSetString(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	assertEqual(t, person, person.SetString("name", "Mat"))
	assertEqual(t, "Mat", person.Fields["name"])
	assertEqual(t, "Mat", person.GetString("name"))
	
}

func TestGetAndSetInt(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	assertEqual(t, person, person.SetInt("age", 27))
	assertEqual(t, int64(27), person.Fields["age"])
	assertEqual(t, int64(27), person.GetInt("age"))
	
}

func TestGetAndSetBool(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	assertEqual(t, person, person.SetBool("field", true))
	assertEqual(t, true, person.Fields["field"])
	assertEqual(t, true, person.GetBool("field"))
	
}

func TestGetAndSetKey(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	var key *datastore.Key = people.NewKey()
	
	assertEqual(t, person, person.SetKey("field", key))
	assertEqual(t, key, person.Fields["field"])
	assertEqual(t, key, person.GetKey("field"))
	
}

func TestSetIDAndGetID(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	person := people.New()
	
	assertEqual(t, NoIDValue, person.ID())
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
