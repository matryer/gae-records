package gaerecords

import (
	"testing"
)

func TestGetKey(t *testing.T) {
	
	model := CreateTestModel()
	key := model.NewKey()
	
	assertEqual(t, model.RecordType(), key.Kind())
	assertEqual(t, "", key.StringID())
	assertEqual(t, true, key.Incomplete())
	
}

func TestNewKeyWithID(t *testing.T) {
	
	model := CreateTestModel()
	key := model.NewKeyWithID(123)
	
	assertEqual(t, model.RecordType(), key.Kind())
	assertEqual(t, int64(123), key.IntID())
	assertEqual(t, "", key.StringID())
	assertEqual(t, false, key.Incomplete())
	
}