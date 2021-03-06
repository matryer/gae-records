package gaerecordsapp

import (
	"goweb"
	"fmt"
	"strconv"
	"gaerecords"
)

/*
	A RESTful controller for 
*/
type PeopleController struct { }

func (cr *PeopleController) New(cx *goweb.Context) {
	
	renderTemplate(cx.ResponseWriter, "People/New", nil)
	
}

func (cr *PeopleController) Create(cx *goweb.Context) {
	
	// create a new appengine context
	gaerecords.CreateAppEngineContext(cx.Request)
	
	// create a new person
	person := People.New()
	
	// get the fields from the HTTP request
	var name string = cx.Request.FormValue("name")
	age, _ := strconv.Atoi64(cx.Request.FormValue("age"))
	
	// set the fields
	person.
		SetString("name", name).
		SetInt64("age", age)
		
	// save it
	err := person.Put()
	
	if err == nil {
		
		// success - redirect to see this person
		cx.RespondWithLocation(fmt.Sprint("/people/", person.ID()))
		
	} else {
		
		// failed - write the error
		cx.ResponseWriter.Write([]byte(fmt.Sprintf("Error: %v", err)))
		
	}
	
}
func (cr *PeopleController) DeleteConfirm(id string, cx *goweb.Context) {

	// create a new appengine context
	gaerecords.CreateAppEngineContext(cx.Request)

	// get the person ID from the URL
	personID, _ := strconv.Atoi64(id)
	
	// load the person
	person, _ := People.Find(personID)

	// create the template data
	data := map[string]interface{}{
		"PersonID":id,
		"PersonName":person.GetString("name"),
	}

	renderTemplate(cx.ResponseWriter, "People/Delete", data)
	
}
func (cr *PeopleController) Delete(id string, cx *goweb.Context) {

	// create a new appengine context
	gaerecords.CreateAppEngineContext(cx.Request)

	// get the person ID from the URL
	personID, _ := strconv.Atoi64(id)
	
	// load the person
	person, _ := People.Find(personID)
	
	// delete the person
	person.Delete()
	
	// send them on their way
	cx.RespondWithLocation("/people");

}
func (cr *PeopleController) DeleteMany(cx *goweb.Context) {
	cx.RespondWithNotImplemented()
}
func (cr *PeopleController) Read(id string, cx *goweb.Context) {
	
	// create a new appengine context
	gaerecords.CreateAppEngineContext(cx.Request)
	
	// get the person ID from the URL
	personID, _ := strconv.Atoi64(id)
	
	// load the person
	person, _ := People.Find(personID)
	
	// create the template data
	data := map[string]interface{}{
		"Person": person.Fields(),
		"PersonID":id,
	}
	
	// render the view
	renderTemplate(cx.ResponseWriter, "People/View", data)
	
}

// /people
func (cr *PeopleController) ReadMany(cx *goweb.Context) {

	// create a new appengine context
	gaerecords.CreateAppEngineContext(cx.Request)
	
	// load all people
	people, _ := People.All()
	
	// collect the fields as an array for the view
	peopleData := make([]map[string]interface{}, 0)
	for _, person := range people {
		
		// save the ID
		person.Set("ID", person.ID())
		peopleData = append(peopleData, person.Fields())
		
	}

	// create the template data
	data := map[string]interface{}{
		"People": peopleData,
	}

	// render the view
	renderTemplate(cx.ResponseWriter, "People/Index", data)

}

func (cr *PeopleController) Edit(id string, cx *goweb.Context) {

	// create a new appengine context
	gaerecords.CreateAppEngineContext(cx.Request)
	
	// get the person ID from the URL
	personID, _ := strconv.Atoi64(id)
	
	// load the person
	person, _ := People.Find(personID)
	
	// create the template data
	data := map[string]interface{}{
		"Person": person.Fields(),
		"PersonID":id,
	}
	
	// render the view
	renderTemplate(cx.ResponseWriter, "People/Edit", data)

}

func (cr *PeopleController) Update(id string, cx *goweb.Context) {

	// create a new appengine context
	gaerecords.CreateAppEngineContext(cx.Request)
	
	// get the person ID from the URL
	personID, _ := strconv.Atoi64(id)
	
	// load the person
	person, _ := People.Find(personID)
	
	// get the fields from the HTTP request
	var name string = cx.Request.FormValue("name")
	age, err := strconv.Atoi64(cx.Request.FormValue("age"))
	
	if err == nil {
		
		// set the fields
		person.
			SetString("name", name).
			SetInt64("age", age)

		// save it
		err := person.Put()

		if err == nil {

			// success - redirect to see this person
			cx.RespondWithLocation(fmt.Sprint("/people/", person.ID()))

		} else {

			// failed - write the error
			cx.ResponseWriter.Write([]byte(fmt.Sprintf("Error: %v", err)))

		}
		
	} else {
	
		// failed to get age - write the error
		cx.ResponseWriter.Write([]byte(fmt.Sprintf("Error getting age: %v", err)))
		
	}
	
}
func (cr *PeopleController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithNotImplemented()
}
