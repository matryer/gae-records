package gaerecords

import (
	"testing"
)

func TestModelAfterFindEvent(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("afterFindEventModel")
	record, _ := CreatePersistedRecord(t, model)
	
	var called bool = false
	var context *EventContext = nil
	
	model.AfterFind.Do(func(c *EventContext){
		called = true
		context = c
	})
	
	// do something that should trigger the event
	model.Find(record.ID())
	
	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, record.ID(), context.Args[0].(*Record).ID())
	
}