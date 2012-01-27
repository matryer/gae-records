package gaerecords

import (
	"testing"
	"appengine/datastore"
)

func TestNewQuery(t *testing.T) {

	model := CreateTestModelWithPropertyType("newQueryModel")
	query := model.NewQuery()

	assertNotNil(t, query, "query")

}

func TestFindByQuery_WithQuery(t *testing.T) {

	model := CreateTestModelWithPropertyType("findByQueryModel")

	record1, _ := CreatePersistedRecord(t, model)
	record2, _ := CreatePersistedRecord(t, model)
	record3, _ := CreatePersistedRecord(t, model)
	record4, _ := CreatePersistedRecord(t, model)
	record5, _ := CreatePersistedRecord(t, model)

	query := model.NewQuery()
	records, err := model.FindByQuery(query)

	if err != nil {
		t.Errorf("FindByQuery: %v", err)
	} else {

		assertEqual(t, 5, len(records))
		assertEqual(t, record1.ID(), records[0].ID())
		assertEqual(t, record2.ID(), records[1].ID())
		assertEqual(t, record3.ID(), records[2].ID())
		assertEqual(t, record4.ID(), records[3].ID())
		assertEqual(t, record5.ID(), records[4].ID())

		assertEqual(t, model, record1.Model())
		assertEqual(t, model, record2.Model())
		assertEqual(t, model, record3.Model())
		assertEqual(t, model, record4.Model())
		assertEqual(t, model, record5.Model())

	}

}

func TestFindByQuery_WithFunc(t *testing.T) {

	model := CreateTestModelWithPropertyType("findByQueryWithFuncModel")

	// create 5 records
	CreatePersistedRecord(t, model)
	CreatePersistedRecord(t, model)
	CreatePersistedRecord(t, model)
	CreatePersistedRecord(t, model)
	CreatePersistedRecord(t, model)

	var called bool = false

	records, err := model.FindByQuery(func(q *datastore.Query) {
		q.Limit(2)
		called = true
	})

	assertEqual(t, true, called)

	if err != nil {
		t.Errorf("%v", err)
	} else {
		assertEqual(t, 2, len(records))
	}

}

func TestFindByQuery_WithFilter_ThroughFunc(t *testing.T) {

	model := CreateTestModelWithPropertyType("findByQueryWithFuncModel")

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

	var called bool = false

	records, err := model.FindByQuery(func(q *datastore.Query) {
		q.Filter("Style=", "A")
		called = true
	})

	assertEqual(t, true, called)

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

func TestFindByQuery_WithFilter_Directly(t *testing.T) {

	model := CreateTestModelWithPropertyType("findByQueryDirectlyModel")

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

	query := model.NewQuery().Filter("Style=", "A")
	records, err := model.FindByQuery(query)

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

func TestFindByQuery_WithFilter_WithLowercaseValues(t *testing.T) {

	model := CreateTestModelWithPropertyType("findByQueryDirectlyWithLowercaseValuesModel")

	record1, _ := CreatePersistedRecord(t, model)
	record2, _ := CreatePersistedRecord(t, model)
	record3, _ := CreatePersistedRecord(t, model)
	record4, _ := CreatePersistedRecord(t, model)
	record5, _ := CreatePersistedRecord(t, model)

	// set some fields
	record1.SetString("style", "A")
	record2.SetString("style", "B")
	record3.SetString("style", "A")
	record4.SetString("style", "B")
	record5.SetString("style", "A")

	// save the new states
	record1.Put()
	record2.Put()
	record3.Put()
	record4.Put()
	record5.Put()

	query := model.NewQuery().Filter("style=", "A")
	records, err := model.FindByQuery(query)

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
