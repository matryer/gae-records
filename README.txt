gae-records: Active Record like wrapper for Google App Engine Datasource in Go

(THIS PROJET IS UNDER CONSTRUCTION AND NOT YET READY FOR USE: But feel free to take a look around)

---

This project aims to simplify the interactions with the appengine/datastore in Go by wrapping the functionality in easy to use classes.

---

// define a type of record
people := NewModel("people")

// create a new person
person := people.New().
  SetString("name", "Mat").
  SetInt("age", 28)

// save this person
person.Put()

// Load and iterate over all people
for p := range people.All() {
  fmt.Println("%v is %v years old", p.GetString("name"), p.GetInt("age"))
}

// load a specific person
mat := people.Find(1)

// change something and save it
mat.SetInt("age", 29).Put()