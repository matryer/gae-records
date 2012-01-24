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

func TestConfigureRecord(t *testing.T) {

	record := new(Record)

	record.needsPersisting = true

	model := CreateTestModelWithPropertyType("configureRecord")
	key := model.NewKeyWithID(123)

	assertEqual(t, record, record.configureRecord(model, key))

	assertEqual(t, false, record.NeedsPersisting())
	assertEqual(t, model, record.model)
	assertEqual(t, key, record.datastoreKey)

}

func TestNeedsPersisting(t *testing.T) {

	model := CreateTestModelWithPropertyType("needsPersistingTestModel")

	// create a non-persisted record
	record := model.New()
	withMessage("NeedsPersisting should be true with New record")
	assertEqual(t, true, record.NeedsPersisting())

	// save it
	record.Put()
	withMessage("NeedsPersisting should be false with record that was just Put")
	assertEqual(t, false, record.NeedsPersisting())

	// load it again
	record, _ = model.Find(record.ID())
	withMessage("NeedsPersisting should be false with record that was just loaded (with Find)")
	assertEqual(t, false, record.NeedsPersisting())

	// change something
	record.Set("name", "Mat")
	withMessage("NeedsPersisting should be true after changing something with Set()")
	assertEqual(t, true, record.NeedsPersisting())

	// load them all
	records, _ := model.FindAll()
	record = records[0]

	withMessage("NeedsPersisting should be false after loading with All()")
	assertEqual(t, false, record.NeedsPersisting())

	// change something (but don't actually change the value)
	record.Set("name", "Mat").Put()
	record.Set("name", "Mat")

	withMessage("NeedsPersisting should be false after changing something to the same value")
	assertEqual(t, false, record.NeedsPersisting())

}

func TestSetNeedsPersisting(t *testing.T) {

	model := CreateTestModelWithPropertyType("needsPersistingModel")
	record := model.New()
	withMessage("NeedsPersisting should be true with New record")
	assertEqual(t, true, record.NeedsPersisting())

	assertEqual(t, record, record.SetNeedsPersisting(false))

	withMessage("NeedsPersisting should be false after SetNeedsPersisting(false)")
	assertEqual(t, false, record.NeedsPersisting())

	assertEqual(t, record, record.SetNeedsPersisting(true))

	withMessage("NeedsPersisting should be false after SetNeedsPersisting(true)")
	assertEqual(t, true, record.NeedsPersisting())

}

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

	assertEqual(t, false, record.IsPersisted())

	err := record.Put()

	if err != nil {
		t.Errorf("Record.Put: %v", err)
	} else {
		assertEqual(t, true, record.IsPersisted())
	}

	// reload the record
	loadedRecord, err := model.Find(record.ID())

	if err != nil {
		t.Errorf("findOneByID: %v", err)
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
		t.Errorf("putOne: %v", err)
	}

	// make some changes
	record.
		SetString("name", "Laurie").
		SetBool("dev", false).
		SetInt64("age", 27)

	err = record.Put()

	if err != nil {
		t.Errorf("putOne: %v", err)
	}

	// reload the record
	loadedRecord, _ := model.Find(record.ID())

	assertEqual(t, "Laurie", loadedRecord.GetString("name"))
	assertEqual(t, false, loadedRecord.GetBool("dev"))
	assertEqual(t, int64(27), loadedRecord.GetInt64("age"))

}

func TestRecordDelete(t *testing.T) {

	model := CreateTestModelWithPropertyType("findAllmodel")
	record, _ := CreatePersistedRecord(t, model)

	recordId := record.ID()

	err := record.Delete()

	if err != nil {
		t.Errorf("deleteOne: %v", err)
	}

	// the record should no longer be 'Persisted'
	assertEqual(t, false, record.IsPersisted())
	assertEqual(t, NoIDValue, record.ID())

	// try and load it
	loadedRecord, err := model.Find(recordId)

	if err == nil || loadedRecord != nil {
		t.Errorf("Error expected when trying to findOneByID a deleted record. The loaded record is: %v", loadedRecord)
	}

}
