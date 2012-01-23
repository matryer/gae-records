package gaerecords

import (
	"testing"
	"appengine"
	"appengine/datastore"
)

func TestSet(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	// Set() should chain
	assertEqual(t, person, person.Set("name", "Mat"))

	// did field update?
	assertEqual(t, "Mat", person.Fields()["name"])

}

func TestGet(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	person.Fields()["age"] = 29

	assertEqual(t, 29, person.Get("age"))

}

func TestGetAndSetString(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	assertEqual(t, person, person.SetString("name", "Mat"))
	assertEqual(t, "Mat", person.Fields()["name"])
	assertEqual(t, "Mat", person.GetString("name"))

}

func TestGetAndSetInt(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	assertEqual(t, person, person.SetInt64("age", 27))
	assertEqual(t, int64(27), person.Fields()["age"])
	assertEqual(t, int64(27), person.GetInt64("age"))

}

func TestGetAndSetFloat64(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	assertEqual(t, person, person.SetFloat64("age", 27.5))
	assertEqual(t, float64(27.5), person.Fields()["age"])
	assertEqual(t, float64(27.5), person.GetFloat64("age"))

}

func TestGetAndSetBool(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	assertEqual(t, person, person.SetBool("field", true))
	assertEqual(t, true, person.Fields()["field"])
	assertEqual(t, true, person.GetBool("field"))

}

func TestGetAndSetKeyField(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	var key *datastore.Key = datastore.NewIncompleteKey(GetAppEngineContext(), "Entity", nil)

	assertEqual(t, person, person.SetKeyField("field", key))
	assertEqual(t, key, person.Fields()["field"])
	assertEqual(t, key, person.GetKeyField("field"))

}

func TestGetAndSetTimeField(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	var time datastore.Time = 1

	assertEqual(t, person, person.SetTime("field", time))
	assertEqual(t, time, person.Fields()["field"])
	assertEqual(t, time, person.GetTime("field"))

}

func TestGetAndSetBlobKeyField(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	var key appengine.BlobKey = "blob"

	assertEqual(t, person, person.SetBlobKey("field", key))
	assertEqual(t, key, person.Fields()["field"])
	assertEqual(t, key, person.GetBlobKey("field"))

}

func TestGetAndSetBytes(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	var bytes []byte = []byte("This is a test value")

	assertEqual(t, person, person.SetBytes("field", bytes))
	assertEqual(t, string(bytes), string(person.Fields()["field"].([]byte)))
	assertEqual(t, string(bytes), string(person.GetBytes("field")))

}

func TestDifferentValueTypes(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	err := person.
		//Set("1", int(1)).
		//Set("2", uint8(1)).     
		//Set("3", uint16(1)).    
		//Set("4", uint32(1)).    
		//Set("5", uint64(1)).    
		//Set("6", int8(1)).      
		//Set("7", int16(1)).     
		//Set("8", int32(1)).     
		Set("9", int64(1)).
		//Set("10", float32(1.1)).   
		Set("11", float64(1.1)).
		//Set("12", complex64(1.1)). 
		//Set("13", complex128(1.1)).
		Set("14", true).
		Put()

	if err != nil {
		t.Errorf("TestDifferentValueTypes failed with error: %v", err)
	}

}

func TestMultipleValues(t *testing.T) {
	
	  //	- int64
    //	- bool
    //	- string
    //	- float64
    //	- *Key
    //	- Time
    //	- appengine.BlobKey
    //	- []byte (up to 1 megabyte in length)
    // This set is smaller than the set of valid struct field types that the
  
	
	
}


