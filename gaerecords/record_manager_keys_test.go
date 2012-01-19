package gaerecords

import (
	"testing"
)

func TestGetKey(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	key := people.NewKey(nil)
	
	assertEqual(t, people.RecordType(), key.Kind())
	
	key2 := people.NewKey(key)
	assertEqual(t, key, key2.Parent())
	
}

func TestNewKeyWithID(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	key := people.NewKeyWithID(123, nil)
	
	assertEqual(t, people.RecordType(), key.Kind())
	assertEqual(t, int64(123), key.IntID())
	
	key2 := people.NewKeyWithID(456, key)
	assertEqual(t, key, key2.Parent())
	assertEqual(t, int64(456), key2.IntID())
	
}