package gaerecords

import (
	"testing"
	"appengine/datastore"
)

/*
	------------------------------------------------------------
	Persistence
	------------------------------------------------------------
*/

func TestLoad(t *testing.T) {
	
	record := CreateTestRecord()
	
	c := make(chan datastore.Property)
	
	go record.Load(c)
	
	c <- datastore.Property{
	        Name:  "name",
	        Value: "Mat",
	}
	c <- datastore.Property{
	        Name:  "age",
	        Value: 27,
	}
	c <- datastore.Property{
	        Name:  "dev",
	        Value: true,
	}
	
	close(c)
	
	// ensure it took the fields
	assertEqual(t, "Mat", record.Fields()["name"])
	assertEqual(t, 27, record.Fields()["age"])
	assertEqual(t, true, record.Fields()["dev"])
	
}

func TestSave(t *testing.T) {
	
	record := CreateTestRecord()
	record.Fields()["name"] = "Mat"
	record.Fields()["age"] = 27
	record.Fields()["dev"] = true
	
	savedProperties := make(map[string]interface{})
	
	c := make(chan datastore.Property)
	
	go func() {
		for property := range c {
			savedProperties[property.Name] = property.Value
		}
	}()
	
	record.Save(c)	
	
	// ensure it saved the fields
	assertEqual(t, "Mat", savedProperties["name"])
	assertEqual(t, 27, savedProperties["age"])
	assertEqual(t, true, savedProperties["dev"])
	
}

func TestPut_Create(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("modelthree")
	record := model.New()
	
	record.
		SetString("name", "Mat").
		SetBool("dev", true).
		SetInt64("age", 29)
	
	assertEqual(t, false, record.IsPersisted());
	
	err := record.Put()
	
	if err != nil {
		t.Errorf("Record.Put: %v", err)
	} else {
		assertEqual(t, true, record.IsPersisted());
	}
	
	// reload the record
	loadedRecord, err := FindOneByID(model, record.ID())
	
	if err != nil {
		t.Errorf("FindOneByID: %v", err)
	}
	
	assertEqual(t, "Mat", loadedRecord.GetString("name"))
	assertEqual(t, true, loadedRecord.GetBool("dev"))
	assertEqual(t, int64(29), loadedRecord.GetInt64("age"))
	
}

func TestPut_Update(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("modelthree")
	record := model.New()
	
	assertEqual(t, false, record.IsPersisted())
	
	// set some fields
	record.
		SetString("name", "Mat").
		SetBool("dev", true).
		SetInt64("age", 29)
	
	err := record.Put()
	
	if err != nil {
		t.Errorf("PutOne: %v", err)
	}
	
	// make some changes
	record.
		SetString("name", "Laurie").
		SetBool("dev", false).
		SetInt64("age", 27)
	
	err = record.Put()
	
	if err != nil {
		t.Errorf("PutOne: %v", err)
	}
	
	// reload the record
	loadedRecord, _ := FindOneByID(model, record.ID())
	
	assertEqual(t, "Laurie", loadedRecord.GetString("name"))
	assertEqual(t, false, loadedRecord.GetBool("dev"))
	assertEqual(t, int64(27), loadedRecord.GetInt64("age"))

	
}

