package gaerecords

import (
	"testing"
)

func TestFind(t *testing.T) {

	model := CreateTestModel()
	record, _ := CreatePersistedRecord(t, model)

	loadedRecord, err := model.Find(record.ID())

	if err != nil {
		t.Errorf("Model.Find: %v", err)
	}

	assertEqual(t, record.ID(), loadedRecord.ID())

}

func TestModelDelete(t *testing.T) {

	model := CreateTestModelWithPropertyType("findAllmodel")
	record, _ := CreatePersistedRecord(t, model)

	recordId := record.ID()

	err := model.Delete(recordId)

	if err != nil {
		t.Errorf("deleteOneByID: %v", err)
	}

	// try and load it
	loadedRecord, err := model.Find(recordId)

	if err == nil || loadedRecord != nil {
		t.Errorf("Error expected when trying to findOneByID a deleted record. The loaded record is: %v", loadedRecord)
	}

}

func TestFindAll(t *testing.T) {

	model := CreateTestModelWithPropertyType("modelAll")
	record1, _ := CreatePersistedRecord(t, model)
	record2, _ := CreatePersistedRecord(t, model)
	record3, _ := CreatePersistedRecord(t, model)
	record4, _ := CreatePersistedRecord(t, model)
	record5, _ := CreatePersistedRecord(t, model)

	records, err := model.FindAll()

	if err != nil {
		t.Errorf("Model.FindAll: %v", err)
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
