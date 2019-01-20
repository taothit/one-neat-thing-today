package main

import (
	"douthitlab.edu/one-new-thing-today/app"
	"github.com/goadesign/goa"
)

// NewController implements the new resource.
type NewController struct {
	*goa.Controller
}

// NewNewController creates a new controller.
func NewNewController(service *goa.Service) *NewController {
	return &NewController{Controller: service.NewController("NewController")}
}

// Today runs the today action.
func (c *NewController) Today(ctx *app.TodayNewContext) error {
	// NewController_Today: start_implement

	// Put your logic here

	return nil
	// NewController_Today: end_implement
}
