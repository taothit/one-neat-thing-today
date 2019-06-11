package main

import (
	"github.com/goadesign/goa"
	"github.com/taothit/one-neat-thing-today/app"
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

	res := &app.NeatThing{}
	


	return ctx.OK(res)
}
