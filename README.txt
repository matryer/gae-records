gaerecords: Go package containing a high performance and lightweight wrapper around appengine/datastore, providing Active Record and DBO style management of data.

Project status: Ready to use

---

Examples:

  // create a new model for 'People'
  People := gaerecords.NewModel("People")

  // create a new person
  mat := People.New()
  mat.
   SetString("name", "Mat")
   SetInt64("age", 28)
   .Put()

  // load person with ID 1
  person := People.Find(1)

  // change some fields
  person.SetInt64("age", 29).Put()

  // load all People
  peeps, _ := People.FindAll()

  // delete mat
  mat.Delete()
  
  // find the first three People by passing a func(*datastore.Query)
  // to the FindByQuery method
  firstThree, _ := People.FindByQuery(func(q *datastore.Query){
    q.Limit(3)
  })
  
  // build your own query and use that
  var ageQuery *datastore.Query = People.NewQuery().
    Limit(3).Order("-age")
  
  // use FindByQuery with a query object
  oldestThreePeople, _ := People.FindByQuery(ageQuery)
  
  // using events, make sure 'People' records always get
  // an 'updatedAt' value set before being put (created and updated)
  People.BeforePut.On(func(c *gaerecords.EventContext){
    person := c.Args[0].(*Record)
    person.SetTime("updatedAt", datastore.SecondsToTime(time.Seconds()))
  })
  
---

Read the documentation by getting the source and running this command:
  
  godoc -http=:6060 -path="path/to/gae-records/gaerecords/"
  
then visit:

  http://localhost:6060/pkg/gaerecords/