# gaerecords

gaerecords is a high performance and lightweight wrapper around [appengine/datastore](http://code.google.com/appengine/docs/go/overview.html), providing Active Record and DBO style management of data.

## Project status

Ready to use

## Usage

    // create a new model for 'People'
    People := gaerecords.NewModel("People")

    // create a new person
    mat := People.New()
    mat.
     SetString("name", "Mat")
     SetInt64("age", 28)
     .Put()

    // load person with ID 1
    person, _ := People.Find(1)

    // change some fields
    person.SetInt64("age", 29).Put()

    // load all People
    peeps, _ := People.FindAll()

    // delete mat
    mat.Delete()

    // delete user with ID 2
    People.Delete(2)

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
    
## Installation

    git clone git://github.com/matryer/gae-records.git
    cd gae-records/gaerecords
    gomake install
    
## Testing

    gotest
    
To properly test the datastore, [gae-go-testing.googlecode.com/git/appenginetesting](http://code.google.com/p/gae-go-testing/) is required.
    
## License

This software is licensed under the terms of the [MIT License](http://en.wikipedia.org/wiki/MIT_License).

## Contributing

We are always keen on getting new people involved on our projects, if you have any ideas, issues or feature requests please get involved.

## Support

Please log defects and feature requests using the issue tracker on github.

## Roadmap

The following items are being considered for future effort (please get in touch if you have a view on these items, or would like other features including)

 * Parent and child records (mirroring Parent and child keys in datastore)
 * Nicer handling of relationships (i.e. Record or []Record as field value)
 * More shortcuts for common queries (like paging, first, last, etc.)

## About

gaerecords was written by [Mat Ryer](http://matryer.com/), follow me on [Twitter](http://www.twitter.com/matryer)
