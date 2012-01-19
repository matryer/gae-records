package gaerecords

import (
	"testing"
)

func TestSimplePropertyList(t *testing.T) {
	
	people := CreateTestPeopleRecordManager()
	
	person := people.New().
							Set("name", "Mat").
							Set("age", 27).
							Set("is_dev", true)
	
	plist := person.GetFieldsAsPropertyList()
	
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