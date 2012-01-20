// Package containing a high performance and lightweight wrapper around appengine/datastore,
// providing Active Record and DBO style management of data.
//
// To persist records, use Record.Put().  To retrieve them, use Model.Find(id).
//
//  // create a new model for 'people'
//  people := NewModel("people")
//  
//  // create a new person
//  mat := people.New()
//  mat.
//    SetString("name", "Mat")
//    SetInt64("age", 28)
//    .Put()
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
package gaerecords
