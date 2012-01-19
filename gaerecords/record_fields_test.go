package gaerecords

import (
	"testing"
	"appengine/datastore"
)

func TestSet(t *testing.T) {
	
	people := CreateTestModel()
	person := people.New()
	
	// Set() should chain
	assertEqual(t, person, person.Set("name", "Mat"))
	
	// did field update?
	assertEqual(t, "Mat", person.Fields()["name"])
	
}

func TestGet(t *testing.T) {
	
	people := CreateTestModel()
	person := people.New()
	
	person.fields["age"] = 29
	
	assertEqual(t, 29, person.Get("age"))
	
}

func TestGetAndSetString(t *testing.T) {
	
	people := CreateTestModel()
	person := people.New()
	
	assertEqual(t, person, person.SetString("name", "Mat"))
	assertEqual(t, "Mat", person.Fields()["name"])
	assertEqual(t, "Mat", person.GetString("name"))
	
}

func TestGetAndSetInt(t *testing.T) {
	
	people := CreateTestModel()
	person := people.New()
	
	assertEqual(t, person, person.SetInt("age", 27))
	assertEqual(t, int64(27), person.Fields()["age"])
	assertEqual(t, int64(27), person.GetInt("age"))
	
}

func TestGetAndSetBool(t *testing.T) {
	
	people := CreateTestModel()
	person := people.New()
	
	assertEqual(t, person, person.SetBool("field", true))
	assertEqual(t, true, person.Fields()["field"])
	assertEqual(t, true, person.GetBool("field"))
	
}

func TestGetAndSetKeyField(t *testing.T) {
	
	people := CreateTestModel()
	person := people.New()
	
	var key *datastore.Key = datastore.NewIncompleteKey(appEngineContext, "Entity", nil)
	
	assertEqual(t, person, person.SetKeyField("field", key))
	assertEqual(t, key, person.Fields()["field"])
	assertEqual(t, key, person.GetKeyField("field"))
	
}