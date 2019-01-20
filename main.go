//go:generate goagen bootstrap -d douthitlab.edu/one-new-thing-today/design

package main

import (
	"douthitlab.edu/one-new-thing-today/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("Discovery")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "new" controller
	c := NewNewThingController(service)
	app.MountNewThingController(service, c)

	s := NewSwaggerController(service)
	app.MountSwaggerController(service, s)
	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
