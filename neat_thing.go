package main

import (
	"douthitlab.edu/one-neat-thing-today/app"
	"github.com/goadesign/goa"
)

// NeatThingController implements the neatThing resource.
type NeatThingController struct {
	*goa.Controller
}

// NewNeatThingController creates a neatThing controller.
func NewNeatThingController(service *goa.Service) *NeatThingController {
	return &NeatThingController{Controller: service.NewController("NeatThingController")}
}

// Today runs the today action.
func (c *NeatThingController) Today(ctx *app.TodayNeatThingContext) error {
	// NeatThingController_Today: start_implement

	// Put your logic here

	res := &app.NeatThing{}
	return ctx.OK(res)
	// NeatThingController_Today: end_implement
}
