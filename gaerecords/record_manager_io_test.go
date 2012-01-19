package gaerecords

import (
	"testing"
)

func TestFind(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	testPerson, _ := CreateTestPerson(t)
		
	foundPerson, _ := people.Find(testPerson.ID())
	
	if foundPerson == nil {
		
		t.Errorf(".Find(1) should find record")
		
	} else {
		
		assertEqual(t, testPerson.Fields["name"], foundPerson.Fields["name"])
		assertEqual(t, testPerson.Fields["age"], foundPerson.Fields["age"])
		
		assertEqual(t, testPerson.ID(), foundPerson.ID())
		
	}
	
}

func TestSave_NewRecord(t *testing.T) {
	
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

func TestSave_ExistingRecord(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	var result bool
	
	person1 := people.New()
	person1.Set("name", "Mat")
	person1.Set("age", int64(29))
	
	// save the person
	result, _ = people.Save(person1)
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
	
	result, _ = people.Save(person1)
	assertEqual(t, true, result)
	
	// ID should not change
	assertEqual(t, id, person1.ID())
	
	// find this record
	person1, _ = people.Find(id)
	
	assertEqual(t, "Laurie", person1.Get("name"))
	assertEqual(t, int64(27), person1.Get("age"))
	
}

func TestAll(t *testing.T) {

	// create 10 people
	people := CreateTestPeopleRecordManager(t)
	person1 := people.New().SetString("name", "Mat")
	person2 := people.New().SetString("name", "Mat")
	person3 := people.New().SetString("name", "Mat")
	person4 := people.New().SetString("name", "Mat")
	person5 := people.New().SetString("name", "Mat")
	person6 := people.New().SetString("name", "Mat")
	person7 := people.New().SetString("name", "Mat")
	person8 := people.New().SetString("name", "Mat")
	person9 := people.New().SetString("name", "Mat")
	person10 := people.New().SetString("name", "Mat")
	
	// save them all
	person1.Save()
	person2.Save()
	person3.Save()
	person4.Save()
	person5.Save()
	person6.Save()
	person7.Save()
	person8.Save()
	person9.Save()
	person10.Save()
	
	// get all
	peeps, err := people.All()
	
	if err != nil {
		t.Errorf(".All() shouldn't raise error: %v", err)
	}
	
	t.Errorf("%v", peeps)
	return
	
	assertEqual(t, 10, len(peeps))
	assertEqual(t, person1.ID(), peeps[0].ID())
	assertEqual(t, person2.ID(), peeps[1].ID())
	assertEqual(t, person3.ID(), peeps[2].ID())
	assertEqual(t, person4.ID(), peeps[3].ID())
	assertEqual(t, person5.ID(), peeps[4].ID())
	assertEqual(t, person6.ID(), peeps[5].ID())
	assertEqual(t, person7.ID(), peeps[6].ID())
	assertEqual(t, person8.ID(), peeps[7].ID())
	assertEqual(t, person9.ID(), peeps[8].ID())
	assertEqual(t, person10.ID(), peeps[9].ID())
	
}


