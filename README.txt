gae-records: Active Record like wrapper for Google App Engine Datasource in Go

(THIS PROJET IS UNDER CONSTRUCTION AND NOT YET READY FOR USE: But feel free to take a look around)

---

This project aims to simplify the interactions with the appengine/datastore in Go by wrapping the functionality in easy to use classes.

---

Package containing a high performance and lightweight wrapper around appengine/datastore,
providing Active Record and DBO style management of data.

 // create a new model for 'people'
 people := NewModel("people")
 
 // create a new person
 mat := people.New()
 mat.
   SetString("name", "Mat")
   SetInt64("age", 28)
   .Put()

 // load person with ID 1
 person := people.Find(1)

 // change some fields
 person.SetInt64("age", 29).Put()

 // load all people
 peeps := people.All()

 // delete mat
 mat.Delete()