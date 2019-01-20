package main

import (
	"douthitlab.edu/one-new-thing-today/app"
	"github.com/goadesign/goa"
)

// NewThingController implements the newThing resource.
type NewThingController struct {
	*goa.Controller
}

// NewNewThingController creates a newThing controller.
func NewNewThingController(service *goa.Service) *NewThingController {
	return &NewThingController{Controller: service.NewController("NewThingController")}
}

// Today runs the today action.
func (c *NewThingController) Today(ctx *app.TodayNewThingContext) error {
	name := "OneNewThingToday Service"
	def := "A sweet new service that gives users a new thing to learn each day."
	link := "https://github.com/taothit/one-new-thing-today"
	newThing := &app.NewThing{
		Name:       &name,
		Definition: &def,
		Link:       &link}

	return ctx.OK(newThing)
}
