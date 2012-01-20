package gaerecords

/*
	Event
	----------------------------------------------------------------------
*/

// Type that enables eventing on an object.
//   // define a struct
//   type MyObject struct {
//
//     // add an event
//     OnSomething Event
//
//   }
// 
//   // create an instance of our type
//   obj := new(MyObject)
//
//   // add an event listener
//   obj.OnSomething.Do(func(e *EventContext){
//     // TODO: handle the event
//   })
//
//   // Trigger the event
//   obj.OnSomething.Trigger()
type Event struct {
	Callbacks []func(*EventContext)
}

// Adds a callback func to this event.  When Trigger() is called, the func passed
// in will get called, provided no other funcs have cancelled the event beforehand.
func (e *Event) Do(f func(*EventContext)) {
	e.Callbacks = append(e.Callbacks, f)
}

// Triggers the event.  All funcs that have been registered with the Do() method will
// be called.  If any of the funcs sets the EventContext.Cancel property to true, no
// more callbacks will be called.
//
// Trigger() returns the EventContext that was passed through each callback which is useful
// for checking if the event chain was cancelled, or if any data was collected along the way.
func (e *Event) Trigger(args ...interface{}) *EventContext {
	
	var context *EventContext = new(EventContext)
	context.Args = args
	context.Cancel = false
	
	for index, c := range e.Callbacks {
		
		// update the index
		context.Index = index
		
		// call the callback
		c(context)
		
		// do we need to cancel?
		if context.Cancel == true {
			break
		}
		
	}
	
	return context
	
}

/*
	EventContext
	----------------------------------------------------------------------
*/

// Type that provides context to event callbacks.
type EventContext struct {
	
	// Whether the event should be cancelled or not.  If set to true inside a 
	// callback func, no subsequent callbacks will be called.
	Cancel bool
	
	// Array holding the arguments passed to Trigger() if any.
	Args []interface{}
	
	// The index of this callback in the chain.  Will be 0 for first callback etc.
	Index int
	
	data map[string]interface{}
}

// Sets some data.
func (c *EventContext) Set(key string, value interface{}) *EventContext {
	
	// set the value
	c.Data()[key] = value;
	
	// chain
	return c
	
}

// Gets a map[string]interface{} of the data for this context.  Will return an
// empty (but non-nil) map if no data has been provided.
func (c *EventContext) Data() map[string]interface{} {
	if c.data == nil {
		c.data = make(map[string]interface{})
	}
	return c.data
}