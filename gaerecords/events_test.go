package gaerecords

import (
	"testing"
)

type TestObject struct {
	OnSomething Event
}

func TestSimpleEvent(t *testing.T) {

	obj := new(TestObject)
	called := false

	obj.OnSomething.Do(func(e *EventContext) {
		called = true
	})

	// trigger the event
	obj.OnSomething.Trigger()

	// the callback method should have been called
	assertEqual(t, true, called)

}

func TestEventWithManyCallbacks(t *testing.T) {

	obj := new(TestObject)
	called := 0

	obj.OnSomething.Do(func(e *EventContext) {
		called++
	})
	obj.OnSomething.Do(func(e *EventContext) {
		called++
	})
	obj.OnSomething.Do(func(e *EventContext) {
		called++
	})

	obj.OnSomething.Trigger()

	assertEqual(t, 3, called)

}

func TestEventTriggerResponse(t *testing.T) {

	obj := new(TestObject)
	var context *EventContext

	obj.OnSomething.Do(func(e *EventContext) {
		context = e
	})

	response := obj.OnSomething.Trigger()

	assertEqual(t, context, response)

}

func TestEventWithArguments(t *testing.T) {

	obj := new(TestObject)
	called := false
	var context *EventContext

	obj.OnSomething.Do(func(e *EventContext) {
		called = true
		context = e
	})

	var arg1 string = "one"
	var arg2 int = 2
	var arg3 bool = true

	// trigger the event
	obj.OnSomething.Trigger(arg1, arg2, arg3)

	// the callback method should have been called
	assertEqual(t, true, called)

	assertEqual(t, arg1, context.Args[0])
	assertEqual(t, arg2, context.Args[1])
	assertEqual(t, arg3, context.Args[2])

}

func TestEventContextIndex(t *testing.T) {

	obj := new(TestObject)
	firstIndex := -1
	secondIndex := -1
	thirdIndex := -1

	obj.OnSomething.Do(func(e *EventContext) {
		firstIndex = e.Index
	})
	obj.OnSomething.Do(func(e *EventContext) {
		secondIndex = e.Index
	})
	obj.OnSomething.Do(func(e *EventContext) {
		thirdIndex = e.Index
	})

	obj.OnSomething.Trigger()

	assertEqual(t, 0, firstIndex)
	assertEqual(t, 1, secondIndex)
	assertEqual(t, 2, thirdIndex)

}

func TestEventCancellation(t *testing.T) {

	obj := new(TestObject)
	called := 0

	obj.OnSomething.Do(func(e *EventContext) {
		called++
	})
	obj.OnSomething.Do(func(e *EventContext) {
		called++

		// cancel here
		e.Cancel = true

	})
	obj.OnSomething.Do(func(e *EventContext) {
		called++
	})

	output := obj.OnSomething.Trigger()

	assertEqual(t, 2, called)
	assertEqual(t, true, output.Cancel)

}

func TestEventData(t *testing.T) {

	obj := new(TestObject)

	obj.OnSomething.Do(func(e *EventContext) {
		e.Set("name", "Mat")
	})

	data := obj.OnSomething.Trigger().Data()
	assertEqual(t, "Mat", data["name"])

}

func TestEventContextSet(t *testing.T) {

	context := new(EventContext)

	assertEqual(t, context, context.Set("name", "Mat"))

	assertEqual(t, "Mat", context.Data()["name"])

}

func TestEventTriggerWithContext(t *testing.T) {

	obj := new(TestObject)
	customContext := new(EventContext)

	var context *EventContext

	obj.OnSomething.Do(func(e *EventContext) {
		context = e
	})

	obj.OnSomething.TriggerWithContext(customContext)

	assertEqual(t, context, customContext)

}
