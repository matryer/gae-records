package gaerecords

import (
	"os"
	"fmt"
	"appengine/datastore"
)

type Record struct {
	Fields map[string]interface{}
	Model *Model
}

/*
	Constuctors
*/
func NewRecord(model *Model) *Record {
	
	record := new(Record)
	record.Fields = make(map[string]interface{})
	record.Model = model
	return record
	
}


/*
	Persistence
	----------------------------------------------------------------------
*/

func (r *Record) Load(c <-chan datastore.Property) os.Error {
		
	// load the fields
	for f := range c {
		r.Fields[f.Name] = f.Value
	}
	
	// no errors
	return nil
}

func (r *Record) Save(c chan<- datastore.Property) os.Error {

	for k, v := range r.Fields {
		c <- datastore.Property{
		        Name:  k,
		        Value: v,
		}
	}

	// this channel is finished
	close(c)
	
	// no errors
	return nil
}

/*
	Fields
	----------------------------------------------------------------------
*/

// Gets the value of a field in a record
func (r *Record) Get(key string) interface{} {
	return r.Fields[key]
}
	
// Sets a field in the record.  The value must be an acceptable datastore
// type or another Record
func (r *Record) Set(key string, value interface{}) *Record {
	r.Fields[key] = value
	return r
}

// Gets a string field
func (r *Record) GetString(key string) string {
	return fmt.Sprint(r.Get(key))
}

// Sets the string value of a field
func (r *Record) SetString(key string, value string) *Record {
	return r.Set(key, value)
}

// Gets an int64 field
func (r *Record) GetInt(key string) int64 {
	return r.Get(key).(int64)
}

// Sets the int64 value of a field
func (r *Record) SetInt(key string, value int64) *Record {
	return r.Set(key, value)
}

// Gets a bool field
func (r *Record) GetBool(key string) bool {
	return r.Get(key).(bool)
}

// Sets the bool value of a field
func (r *Record) SetBool(key string, value bool) *Record {
	return r.Set(key, value)
}

// Gets a key field
func (r *Record) GetKeyField(key string) *datastore.Key {
	return r.Get(key).(*datastore.Key)
}

// Sets the key value of a field
func (r *Record) SetKeyField(key string, value *datastore.Key) *Record {
	return r.Set(key, value)
}