package gaerecords

import (
	"testing"
	"appengine/datastore"
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

func TestFindByFilter(t *testing.T) {

	model := CreateTestModelWithPropertyType("findByFilterDirectlyModel")

	record1, _ := CreatePersistedRecord(t, model)
	record2, _ := CreatePersistedRecord(t, model)
	record3, _ := CreatePersistedRecord(t, model)
	record4, _ := CreatePersistedRecord(t, model)
	record5, _ := CreatePersistedRecord(t, model)

	// set some fields
	record1.SetString("Style", "A")
	record2.SetString("Style", "B")
	record3.SetString("Style", "A")
	record4.SetString("Style", "B")
	record5.SetString("Style", "A")

	// save the new states
	record1.Put()
	record2.Put()
	record3.Put()
	record4.Put()
	record5.Put()

	records, err := model.FindByField("Style=", "A")

	if err != nil {
		t.Errorf("%v", err)
	} else {

		assertEqual(t, 3, len(records))

		if len(records) != 3 {

			t.Errorf("3 records expected, not %v.", len(records))

		} else {

			assertEqual(t, record1.ID(), records[0].ID())
			assertEqual(t, record3.ID(), records[1].ID())
			assertEqual(t, record5.ID(), records[2].ID())

		}

	}

}

func TestFindByFilter_WithModifierFunc(t *testing.T) {

	model := CreateTestModelWithPropertyType("findByQueryDirectlyWithModifierFuncModel")

	record1, _ := CreatePersistedRecord(t, model)
	record2, _ := CreatePersistedRecord(t, model)
	record3, _ := CreatePersistedRecord(t, model)
	record4, _ := CreatePersistedRecord(t, model)
	record5, _ := CreatePersistedRecord(t, model)

	// set some fields
	record1.SetString("Style", "A")
	record2.SetString("Style", "B")
	record3.SetString("Style", "A")
	record4.SetString("Style", "B")
	record5.SetString("Style", "A")

	// save the new states
	record1.Put()
	record2.Put()
	record3.Put()
	record4.Put()
	record5.Put()

	records, err := model.FindByField("Style=", "B", func(q *datastore.Query) {
		q.Limit(1)
	})

	if err != nil {
		t.Errorf("%v", err)
	} else {

		assertEqual(t, 1, len(records))

		if len(records) != 1 {

			t.Errorf("1 record expected, not %v.", len(records))

		} else {

			assertEqual(t, record2.ID(), records[0].ID())

		}

	}

}
