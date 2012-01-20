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

/*
func TestPut(t *testing.T) {
	
	model := CreateTestModel()
	record, _ := CreatePersistedRecord(model)
	
	assertEqual(t, false, record.IsPersisted());
	
	err := record.Put()
	
	if err != nil {
		t.Errorf("Record.Put: %v", err)
	} else {
		assertEqual(t, true, record.IsPersisted());
	}
	
}
*/