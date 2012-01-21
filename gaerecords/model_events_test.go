package gaerecords

import (
	"os"
	"testing"
)

func TestModelAfterFindEvent(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterFindEventModel")
	record, _ := CreatePersistedRecord(t, model)

	var called bool = false
	var context *EventContext = nil

	model.AfterFind.Do(func(c *EventContext) {
		called = true
		context = c
	})

	// do something that should trigger the event
	model.Find(record.ID())

	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, record.ID(), context.Args[0].(*Record).ID())

}

func TestModelAfterFindEvent_withFindAll(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterFindEventALLModel")
	CreatePersistedRecord(t, model)
	record2, _ := CreatePersistedRecord(t, model)

	var called int = 0
	var lastContext *EventContext = nil

	model.AfterFind.Do(func(c *EventContext) {
		called++
		lastContext = c
	})

	// do something that should trigger the event
	model.All()

	assertEqual(t, 2, called)
	assertNotNil(t, lastContext.Args[0], "context.Args[0]")
	assertEqual(t, record2.ID(), lastContext.Args[0].(*Record).ID())

}

func TestModelBeforeDeleteByIDEvent(t *testing.T) {

	model := CreateTestModelWithPropertyType("beforeDeleteEventModel")
	record, _ := CreatePersistedRecord(t, model)
	var loadedRecord *Record

	var called bool = false
	var context *EventContext = nil

	model.BeforeDelete.Do(func(c *EventContext) {
		called = true
		context = c
		loadedRecord, _ = model.Find(record.ID())
	})

	// do something that should trigger the event
	model.Delete(record.ID())

	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, record.ID(), context.Args[0].(int64))
	assertEqual(t, nil, context.Args[1])
	assertEqual(t, record.ID(), loadedRecord.ID())

}

func TestModelBeforeDeleteByIDEvent_Cancellation(t *testing.T) {

	model := CreateTestModelWithPropertyType("beforeDeleteEventCancelModel")
	record, _ := CreatePersistedRecord(t, model)

	var called bool = false
	var context *EventContext = nil

	model.BeforeDelete.Do(func(c *EventContext) {
		called = true
		context = c

		// cancel the delete
		c.Cancel = true

	})

	// do something that should trigger the event
	err := model.Delete(record.ID())

	if err != ErrOperationCancelledByEventCallback {
		t.Errorf("ErrOperationCancelledByEventCallback Error should be returned if the delete operation was cancelled by an event callback")
	}

	foundRecord, _ := model.Find(record.ID())

	assertEqual(t, record.ID(), foundRecord.ID())

}

func TestModelAfterDeleteByID(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterDeleteByIDModel")
	record, _ := CreatePersistedRecord(t, model)
	var err os.Error = nil
	var loadedRecord *Record = nil

	var called bool = false
	var context *EventContext = nil

	model.AfterDelete.Do(func(c *EventContext) {
		called = true
		context = c
		loadedRecord, err = model.Find(record.ID())
	})

	if loadedRecord != nil {
		t.Errorf("loadedRecord should return nil after delete")
	}

	// do something that should trigger the event
	model.Delete(record.ID())

	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, nil, context.Args[1])
	assertEqual(t, record.ID(), context.Args[0].(int64))

}

func TestBeforeAndAfterDeleteByIDEventsShareContext(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterDeleteByIDModel")
	record, _ := CreatePersistedRecord(t, model)

	var context1 *EventContext = nil
	var context2 *EventContext = nil

	model.BeforeDelete.Do(func(c *EventContext) {
		context1 = c
	})
	model.AfterDelete.Do(func(c *EventContext) {
		context2 = c
	})

	// trigger the event
	model.Delete(record.ID())

	// make sure they match
	assertEqual(t, context1, context2)

}

func TestModelBeforeDeleteEvent(t *testing.T) {

	model := CreateTestModelWithPropertyType("beforeDeleteEventModel")
	record, _ := CreatePersistedRecord(t, model)
	var loadedRecord *Record

	var called bool = false
	var context *EventContext = nil

	model.BeforeDelete.Do(func(c *EventContext) {
		called = true
		context = c
		loadedRecord, _ = model.Find(record.ID())
	})

	recordID := record.ID()

	// do something that should trigger the event
	record.Delete()

	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, recordID, context.Args[0].(int64))
	assertEqual(t, record, context.Args[1])
	assertEqual(t, recordID, loadedRecord.ID())

}

func TestModelBeforeDeleteEvent_Cancellation(t *testing.T) {

	model := CreateTestModelWithPropertyType("beforeDeleteEventCancelModel")
	record, _ := CreatePersistedRecord(t, model)

	var called bool = false
	var context *EventContext = nil

	model.BeforeDelete.Do(func(c *EventContext) {
		called = true
		context = c

		// cancel the delete
		c.Cancel = true

	})

	recordID := record.ID()

	// do something that should trigger the event
	err := record.Delete()

	if err != ErrOperationCancelledByEventCallback {
		t.Errorf("ErrOperationCancelledByEventCallback Error should be returned if the delete operation was cancelled by an event callback")
	}

	foundRecord, _ := model.Find(record.ID())

	assertEqual(t, recordID, foundRecord.ID())

}

func TestModelAfterDelete(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterDeleteByIDModel")
	record, _ := CreatePersistedRecord(t, model)
	var err os.Error = nil
	var loadedRecord *Record = nil

	var called bool = false
	var context *EventContext = nil

	model.AfterDelete.Do(func(c *EventContext) {
		called = true
		context = c
		loadedRecord, err = model.Find(record.ID())
	})

	if loadedRecord != nil {
		t.Errorf("loadedRecord should return nil after delete")
	}

	recordID := record.ID()

	// do something that should trigger the event
	record.Delete()

	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, record, context.Args[1])
	assertEqual(t, recordID, context.Args[0].(int64))

}

func TestBeforeAndAfterDeleteEventsShareContext(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterDeleteByIDModel")
	record, _ := CreatePersistedRecord(t, model)

	var context1 *EventContext = nil
	var context2 *EventContext = nil

	model.BeforeDelete.Do(func(c *EventContext) {
		context1 = c
	})
	model.AfterDelete.Do(func(c *EventContext) {
		context2 = c
	})

	// trigger the event
	record.Delete()

	// make sure they match
	assertEqual(t, context1, context2)

}

func TestModelBeforePutEvent(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterFindEventModel")
	record := model.New().Set("something", "something")

	var called bool = false
	var context *EventContext = nil

	model.BeforePut.Do(func(c *EventContext) {
		called = true
		context = c
	})

	// do something that should trigger the event
	record.Put()

	assertEqual(t, true, record.IsPersisted())

	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, record.ID(), context.Args[0].(*Record).ID())

}

func TestModelBeforePutEvent_Cancellation(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterFindEventModel")
	record := model.New().Set("something", "something")

	var called bool = false
	var context *EventContext = nil

	model.BeforePut.Do(func(c *EventContext) {
		called = true
		context = c

		c.Cancel = true

	})

	// do something that should trigger the event
	err := record.Put()

	if err != ErrOperationCancelledByEventCallback {
		t.Errorf("ErrOperationCancelledByEventCallback Error should be returned if the delete operation was cancelled by an event callback")
	}

	assertEqual(t, false, record.IsPersisted())

	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, record.ID(), context.Args[0].(*Record).ID())

}

func TestModelAfterPutEvent(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterFindEventModel")
	record := model.New().Set("something", "something")

	var called bool = false
	var context *EventContext = nil

	model.AfterPut.Do(func(c *EventContext) {
		called = true
		context = c
	})

	// do something that should trigger the event
	record.Put()

	assertEqual(t, true, record.IsPersisted())

	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, record.ID(), context.Args[0].(*Record).ID())

}

func TestBeforeAndAfterPutEventsShareContext(t *testing.T) {

	model := CreateTestModelWithPropertyType("puttingSharedContext")
	record, _ := CreatePersistedRecord(t, model)

	var context1 *EventContext = nil
	var context2 *EventContext = nil

	model.BeforePut.Do(func(c *EventContext) {
		context1 = c
	})
	model.AfterPut.Do(func(c *EventContext) {
		context2 = c
	})

	// trigger the event
	record.Put()

	// make sure they match
	assertEqual(t, context1, context2)

}

func TestAfterNewEvent(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterNewEventModel")

	var called bool = false
	var context *EventContext = nil

	model.AfterNew.Do(func(c *EventContext) {
		context = c
		called = true
	})

	newRecord := model.New()

	assertEqual(t, true, called)
	if called {
		assertEqual(t, newRecord, context.Args[0])
	}

}

func TestOnChangedEvent(t *testing.T) {

	model := CreateTestModelWithPropertyType("onChangedEventModel")
	record := model.New()

	var called bool = false
	var context *EventContext

	// sign up to the OnChanged event
	model.OnChanged.Do(func(c *EventContext) {
		called = true
		context = c
	})

	// change something
	record.Set("name", "Mat")

	assertEqual(t, true, called)
	if called {
		assertEqual(t, record, context.Args[0])
		assertEqual(t, "name", context.Args[1])
		assertEqual(t, "Mat", context.Args[2])
		assertEqual(t, nil, context.Args[3])
	}

	// do it again - to check the 'old' argument
	called = false

	record.Set("name", "Laurie")

	assertEqual(t, true, called)
	if called {
		assertEqual(t, record, context.Args[0])
		assertEqual(t, "name", context.Args[1])
		assertEqual(t, "Laurie", context.Args[2])
		assertEqual(t, "Mat", context.Args[3])
	}

}
