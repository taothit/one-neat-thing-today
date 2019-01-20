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
	// NewThingController_Today: start_implement

	// Put your logic here

	return nil
	// NewThingController_Today: end_implement
}
