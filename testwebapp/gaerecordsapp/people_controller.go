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
	
	// get the fields
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
func (cr *PeopleController) Delete(id string, cx *goweb.Context) {
	cx.RespondWithNotImplemented()
}
func (cr *PeopleController) DeleteMany(cx *goweb.Context) {
	cx.RespondWithNotImplemented()
}
func (cr *PeopleController) Read(id string, cx *goweb.Context) {
	
	
	
}

// /people
func (cr *PeopleController) ReadMany(cx *goweb.Context) {
	cx.RespondWithNotImplemented()
}
func (cr *PeopleController) Update(id string, cx *goweb.Context) {
	cx.RespondWithNotImplemented()
}
func (cr *PeopleController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithNotImplemented()
}
