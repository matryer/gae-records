package gaerecords

import (
	"testing"
)

func TestFindOneByID(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("modeltwo")
	record, err := CreatePersistedRecord(t, model)
	
	if err != nil {
		t.Errorf("CreatePersistedRecord: %v", err)
	}
	if record == nil {
		t.Errorf("CreatePersistedRecord didn't create the record")
	}
	
	loadedRecord, err2 := FindOneByID(model, record.ID())
	
	if err2 != nil {
		t.Errorf("FindOneByID: %v", err2)
	}
	if loadedRecord == nil {
		t.Errorf("FindOneByID didn't create the record")
	}
	
	if record != nil && loadedRecord != nil {
		
		assertEqual(t, record.ID(), loadedRecord.ID())
		
		assertEqual(t, "Mat", record.GetString("name"))
		assertEqual(t, int64(29), record.GetInt64("age"))
		assertEqual(t, model.RecordType(), record.Model().RecordType())
		
	}
	
}

func TestFindAll(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("findallmodel")
	record1, _ := CreatePersistedRecord(t, model)
	record2, _ := CreatePersistedRecord(t, model)
	record3, _ := CreatePersistedRecord(t, model)
	record4, _ := CreatePersistedRecord(t, model)
	record5, _ := CreatePersistedRecord(t, model)
	
	records, err := FindAll(model)
	
	if err != nil {
		t.Errorf("FindAll: %v", err)
	} else {
		
		// validate the records
		
		assertEqual(t, 5, len(records))
		
		assertEqual(t, record1.ID(), records[0].ID())
		assertEqual(t, record2.ID(), records[1].ID())
		assertEqual(t, record3.ID(), records[2].ID())
		assertEqual(t, record4.ID(), records[3].ID())
		assertEqual(t, record5.ID(), records[4].ID())
		
	}
	
}

func TestPutOne_Create(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("modelthree")
	record := model.New()
	
	assertEqual(t, false, record.IsPersisted())
	
	// set some fields
	record.
		SetString("name", "Mat").
		SetBool("dev", true).
		SetInt64("age", 29)
	
	err := PutOne(record)
	
	if err != nil {
		t.Errorf("PutOne: %v", err)
	}
	
	assertEqual(t, true, record.IsPersisted())
	
	// ensure the record ID was updated
	if record.ID() == NoIDValue {
		t.Errorf("Record ID shouldn't be %v", NoIDValue)
	}
	
	// reload the record
	loadedRecord, _ := FindOneByID(model, record.ID())
	
	assertEqual(t, "Mat", loadedRecord.GetString("name"))
	assertEqual(t, true, loadedRecord.GetBool("dev"))
	assertEqual(t, int64(29), loadedRecord.GetInt64("age"))
	
}

func TestPutOne_Update(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("modelthree")
	record := model.New()
	
	assertEqual(t, false, record.IsPersisted())
	
	// set some fields
	record.
		SetString("name", "Mat").
		SetBool("dev", true).
		SetInt64("age", 29)
	
	err := PutOne(record)
	
	if err != nil {
		t.Errorf("PutOne: %v", err)
	}
	
	// make some changes
	record.
		SetString("name", "Laurie").
		SetBool("dev", false).
		SetInt64("age", 27)
	
	err = PutOne(record)
	
	if err != nil {
		t.Errorf("PutOne: %v", err)
	}
	
	// reload the record
	loadedRecord, _ := FindOneByID(model, record.ID())
	
	assertEqual(t, "Laurie", loadedRecord.GetString("name"))
	assertEqual(t, false, loadedRecord.GetBool("dev"))
	assertEqual(t, int64(27), loadedRecord.GetInt64("age"))

	
}
