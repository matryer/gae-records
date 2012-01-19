package gaerecords

import (
	"testing"
)

func TestRecordSave_NewRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	
	person1 := people.New()
	person1.Fields["name"] = "Mat"
	person1.Fields["age"] = int64(27)
	
	result, _ := people.Save(person1)
	assertEqual(t, true, result)
	
	if person1.ID() == NoIDValue {
		t.Errorf("SaveChanges() should cause the ID to be updated")
	}
	
}

func TestRecordSave_ExistingRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	
	person1 := people.New()
	person1.Set("name", "Mat")
	person1.Set("age", int64(29))
	
	// save the person
	result, _ := person1.Save()
	assertEqual(t, true, result)
	
	// old ID
	id := person1.ID()
	
	// find this record
	person1, _ = people.Find(id)
	
	assertEqual(t, "Mat", person1.Get("name"))
	assertEqual(t, int64(29), person1.Get("age"))
	
	
	// change the name and age
	person1.Set("name", "Laurie")
	person1.Set("age", int64(27))
	
	result, _ = person1.Save()
	assertEqual(t, true, result)
	
	// ID should not change
	assertEqual(t, id, person1.ID())
	
	// find this record
	person1, _ = people.Find(id)
	
	assertEqual(t, "Laurie", person1.Get("name"))
	assertEqual(t, int64(27), person1.Get("age"))
	
}