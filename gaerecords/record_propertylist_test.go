package gaerecords

import (
	"testing"
	"appengine/datastore"
)

func TestGetFieldsAsPropertyList(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	
	person := people.New().
							Set("name", "Mat").
							Set("age", 27).
							Set("is_dev", true)
	
	var plist datastore.PropertyList = person.GetFieldsAsPropertyList()
	
	assertEqual(t, 3, len(plist))
	
	// FIXME IF THESE TESTS FAIL: 
	// don't trust the order of the fields
	// loopup each key when validating
	
	assertEqual(t, "name", plist[0].Name)
	assertEqual(t, "Mat", plist[0].Value)
	
	assertEqual(t, "age", plist[2].Name)
	assertEqual(t, 27, plist[2].Value)
	
	assertEqual(t, "is_dev", plist[1].Name)
	assertEqual(t, true, plist[1].Value)
	
}

func TestSetFieldsFromPropertyList(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	
	person := people.New()
	
	// build a test property list
	var plist datastore.PropertyList = make(datastore.PropertyList, 3)
	plist[0] = datastore.Property{ "name", "Mat", false, false }
	plist[1] = datastore.Property{ "age", int64(29), false, false }
	plist[2] = datastore.Property{ "is_dev", true, false, false }
	
	// inject it into the record
	person.SetFieldsFromPropertyList(plist)
	
	assertEqual(t, "Mat", person.Fields["name"])
	assertEqual(t, int64(29), person.Fields["age"])
	assertEqual(t, true, person.Fields["is_dev"])
	
}