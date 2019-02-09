//go:generate goagen bootstrap -d douthitlab.edu/one-neat-thing-today/design

package main

import (
	"flag"
	"fmt"

	"douthitlab.edu/one-neat-thing-today/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

var port = flag.String("port", "8080", "port service listens on")

func main() {
	flag.Parse()
	// Create service
	service := goa.New("Discovery")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "neatThing" controller
	c := NewNeatThingController(service)
	app.MountNeatThingController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(fmt.Sprintf(":%s", *port)); err != nil {
		service.LogError("startup", "err", err)
	}

}
