package gaerecords

import (
	"os"
	"fmt"
	"testing"
	"appengine/datastore"
	"gae-go-testing.googlecode.com/git/appenginetesting"
)

var modelIndex int = 1

// Creates a test appengine context object
func UseTestAppEngineContext() {

	AppEngineContext, _ = appenginetesting.NewContext(nil)

}

func CreateTestRecord() *Record {
	return NewRecord(CreateTestModel())
}

func CreateTestModel() *Model {
	modelIndex++
	return NewModel(fmt.Sprint("model-", modelIndex))
}
func CreateTestModelWithPropertyType(kind string) *Model {
	return NewModel(kind)
}

func CreatePersistedRecord(t *testing.T, model *Model) (*Record, os.Error) {

	context := GetAppEngineContext()
	person := model.New()
	key := person.DatastoreKey()

	person.
		SetString("name", "Mat").
		SetInt64("age", 29)

	newKey, err := datastore.Put(context, key, datastore.PropertyLoadSaver(person))

	t.Logf(">> Created new record with ID: '%v'", newKey)

	if err != nil {
		return nil, err
	}

	// set the person ID
	person.setID(newKey.IntID())

	return person, nil

}

func TestSetup(t *testing.T) {

	UseTestAppEngineContext()
	t.Logf("<<< Test context created %v >>>", AppEngineContext)

}
