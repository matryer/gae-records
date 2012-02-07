package gaerecords

import (
	"testing"
	"appengine/datastore"
)

func TestNewKey(t *testing.T) {

	model := CreateTestModel()
	key := model.NewKey()

	assertEqual(t, model.RecordType(), key.Kind())
	assertEqual(t, "", key.StringID())
	assertEqual(t, true, key.Incomplete())

}

func TestNewKeyWithID(t *testing.T) {

	model := CreateTestModelWithPropertyType("modelname1")
	key := model.NewKeyWithID(123)

	assertEqual(t, model.RecordType(), key.Kind())
	assertEqual(t, int64(123), key.IntID())
	assertEqual(t, "", key.StringID())
	assertEqual(t, false, key.Incomplete())

}

func TestNewKeyWithParent(t *testing.T) {

	model := CreateTestModel()
	parentKey := datastore.NewIncompleteKey(model.AppEngineContext(), "parent-key-kind", nil)
	key := model.NewKeyWithParent(parentKey)

	assertEqual(t, model.RecordType(), key.Kind())
	assertEqual(t, "", key.StringID())
	assertEqual(t, true, key.Incomplete())

	assertEqual(t, "parent-key-kind", key.Parent().Kind())

}

func TestNewKeyWithIDWithParent(t *testing.T) {

	model := CreateTestModelWithPropertyType("modelname2")
	parentKey := datastore.NewIncompleteKey(model.AppEngineContext(), "parent-key-kind", nil)
	key := model.NewKeyWithIDAndParent(123, parentKey)

	assertEqual(t, model.RecordType(), key.Kind())
	assertEqual(t, int64(123), key.IntID())
	assertEqual(t, "", key.StringID())
	assertEqual(t, false, key.Incomplete())
	
	assertEqual(t, "parent-key-kind", key.Parent().Kind())

}
