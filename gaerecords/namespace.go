// Package containing a high performance and lightweight wrapper around appengine/datastore,
// providing Active Record and DBO style management of data.
//
//  // create a new model for 'people'
//  people := NewModel("people")
//  
//  // create a new person
//  mat := people.New()
//  mat.
//    SetString("name", "Mat").
//    SetInt64("age", 28).
//    Put()
//
//  // load person with ID 1
//  person := people.Find(1)
//
//  // change some fields
//  person.SetInt64("age", 29).Put()
//
//  // load all people
//  peeps := people.All()
//
//  // delete mat
//  mat.Delete()
//
// Supported types are the same as those supported by the datastore.Property's Value object
// plus any array or slice of those types, which will be saved as a property with multiple values.
// (see http://code.google.com/appengine/docs/go/datastore/reference.html#Property)
//    - int64
//    - []int64
//    - bool
//    - []bool
//    - string
//    - []string
//    - float64
//    - []float64
//    - *datastore.Key
//    - []*datastore.Key
//    - Time
//    - []Time
//    - appengine.BlobKey
//    - []appengine.BlobKey
//
// KNOWN BUG: []byte properties are not currently working.  See https://github.com/matryer/gae-records/issues/1
package gaerecords
