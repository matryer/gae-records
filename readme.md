# gaerecords

gaerecords is a lightweight wrapper around [appengine/datastore](http://code.google.com/appengine/docs/go/overview.html), providing Active Record and DBO style management of data.

## Project status

Ready to use

## Usage examples

### Creating records

    // create a new model for 'People'
    People := gaerecords.NewModel("People")

    // create a new person
    mat := People.New()
    mat.
     SetString("name", "Mat")
     SetInt64("age", 28)
     .Put()

### Changing records

    // change some fields and Put it
    person.SetInt64("age", 29).Put()

### Deleting records

    // delete mat
    mat.Delete()

    // delete user with ID 2
    People.Delete(2)

### Finding records

    // load person with ID 1
    person, _ := People.Find(1)

    // load all People
    peeps, _ := People.FindAll()
    
### Counting records
    
    // get the total number of people
    total, _ := People.Count()
    
    // get the total number of male people
    totalMen, _ := People.Count(func(q *datastore.Query){
      q.Filter("IsMale=", true)
    })
    
### Working with pages of records
    
    // get three pages of people with 10 records on each page
    peoplePageOne, _ := People.FindByPage(1, 10)
    peoplePageTwo, _ := People.FindByPage(2, 10)
    peoplePageThree, _ := People.FindByPage(3, 10)

    // get the number of pages if we have 10 records per page
    totalPages = People.LoadPagingInfo(10, 1).TotalPages

### Finding records by querying

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

    // find all people that are 'active'
    activePeople, _ := People.FindByFilter("Active=", true)

### Binding to model events

    // using events, make sure 'People' records always get
    // an 'updatedAt' value set before being put (created and updated)
    People.BeforePut.On(func(c *gaerecords.EventContext){
      person := c.Args[0].(*Record)
      person.SetTime("updatedAt", datastore.SecondsToTime(time.Seconds()))
    })
    
## Concepts

gaerecords provides two main types that represent the different data for your project.
A <code>Model</code> describes a type of data, and a <code>Record</code> is a single entity
or instance of that type.

Creating models is as easy as calling the <code>gaerecords.NewModel</code> method.

For example, a typical blogging application might define these models:

    Authors := gaerecords.NewModel("Author")
    Posts := gaerecords.NewModel("Post")
    Comments := gaerecords.NewModel("Comment")
    
And to create a new blog post is as simple as:

    newPost := Posts.New()
    newPost.SetString("title", "My Blog Post").
            SetString("body", "My blog text goes here...")
    // save it
    newPost.Put()

<code>Model</code>'s also support events for when interesting things happen to records, and you
can bind to these by providing an additional initializer func(*Model) to the <code>NewModel</code> method.

    People := gaerecords.NewModel("people", func(model *gaerecords.Model){
      
      // bind to the BeforePut event
      model.BeforePut.On(func(e *gaerecords.EventContext){
        
        // do something before records are saved
        
      })
      
    })

### Model
    
<code>Model</code> objects allow you to perform operations on sets of data, such as create a 
new record, find records etc.

### Record

<code>Record</code> objects allow to you perform operations on a specific entity, such as set fields,
save changes, delete it.

### Event

The <code>Event</code> type (and its <code>EventContext</code> younger brother) allows you to bind your
own callbacks to the lifecycle of records.  For example, before or after a <code>Record</code> gets <code>Put</code> (saved),
or after a <code>Record</code> has been deleted.

The <code>Model</code> has the events, but actions to records can cause the callbacks to get run.
    
## Installation

    git clone git://github.com/matryer/gae-records.git
    cd gae-records/gaerecords
    gomake install
    
## Testing

### Testing gaerecords

    gotest
    
To properly test the datastore, [gae-go-testing.googlecode.com/git/appenginetesting](http://code.google.com/p/gae-go-testing/) is required.
    
### Testing your own code that uses gaerecords

The [gae-go-testing.googlecode.com/git/appenginetesting](http://code.google.com/p/gae-go-testing/) library allows you to test code that relies on Google App Engine.

Follow these recommended patterns and practices:

1. Have a test file with an 'early' name (i.e. <code>0\_setup\_test.go</code> - to ensure the tests get run before others) that setup the expected state of your datastore.  <code>appenginetesting</code> does persist entities across different runs, so be sure to delete any that you plan to create or at least be aware of this when writing tests.
1. In the setup file, create a test App Engine context that you can reuse for your entire test suite.  See our [test setup file](https://github.com/matryer/gae-records/blob/master/gaerecords/0_app_helper_test.go) for an example of how a test App Engine context is created.
1. Don't assume that the IDs will be consistent.  It is best practice to discover the ID of a newly created record if you plan to refer to that record later, instead of assuming it will have an ID of 1.
1. Remember to <code>Close()</code> your appenginetesting context as per the documentation, otherwise you'll end up with lots of Python processes running in Terminal.  It is a good idea to do this in a 'late' test file (i.e. <code>z\_cleanup\_test.go</code>).  See our [clean-up file](https://github.com/matryer/gae-records/blob/master/gaerecords/z_cleanup_test.go) for an example (maybe even take that file wholesale?)
    
## License

This software is licensed under the terms of the [MIT License](http://en.wikipedia.org/wiki/MIT_License).

## Contributing

We are always keen on getting new people involved on our projects, if you have any ideas, issues or feature requests please get involved.

## Support

Please log defects and feature requests using the issue tracker on github.

## Roadmap

The following items are being considered for future effort (please get in touch if you have a view on these items, or would like other features including)

 * Parent and child records (mirroring Parent and child keys in datastore)
 * Strongly typed records
 * Better handling of relationships (i.e. Record or []Record as field value)
 * More shortcuts for common queries (like paging, first, last, etc.)

## About

gaerecords was written by Mat Ryer, follow me on [Twitter](http://www.twitter.com/matryer)
