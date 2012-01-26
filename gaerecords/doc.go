// gaerecords provides a high performance and lightweight wrapper around appengine/datastore,
// providing Active Record and DBO style management of data.
//
//  // create a new model for 'People'
//  People := gaerecords.NewModel("People")
//  
//  // create a new person
//  mat := People.New()
//  mat.
//   SetString("name", "Mat")
//   SetInt64("age", 28)
//   .Put()
//  
//  // load person with ID 1
//  person := People.Find(1)
//  
//  // change some fields
//  person.SetInt64("age", 29).Put()
//  
//  // load all People
//  peeps, _ := People.FindAll()
//  
//  // delete mat
//  mat.Delete()
//
//  // delete user with ID 2
//  People.Delete(2)
//  
//  // find the first three People by passing a func(*datastore.Query)
//  // to the FindByQuery method
//  firstThree, _ := People.FindByQuery(func(q *datastore.Query){
//    q.Limit(3)
//  })
//  
//  // build your own query and use that
//  var ageQuery *datastore.Query = People.NewQuery().
//    Limit(3).Order("-age")
//  
//  // use FindByQuery with a query object
//  oldestThreePeople, _ := People.FindByQuery(ageQuery)
//  
//  // using events, make sure 'People' records always get
//  // an 'updatedAt' value set before being put (created and updated)
//  People.BeforePut.On(func(c *gaerecords.EventContext){
//    person := c.Args[0].(*Record)
//    person.SetTime("updatedAt", datastore.SecondsToTime(time.Seconds()))
//  })
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
// bug(matryer): []byte properties are not currently working.  See https://github.com/matryer/gae-records/issues/1
package gaerecords
