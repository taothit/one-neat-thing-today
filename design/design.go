package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Discovery", func() {
	Title("A new thing each day")
	Description("Discover something new each day by being given an item from a curated list.")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("newThing", func() {
	Description("Something new for you")
	DefaultMedia(NewThingMedia)
	BasePath("/new/thing")
	Action("today", func() {
		Routing(GET("today"))
		Description("GET the new thing for the day")
		Response(OK, NewThingMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
})

// NewThingMedia represents a new thing with its name, description/definition,
// and a great illustrative link.
var NewThingMedia = MediaType("application/vnd.douthitlab.newthing", func() {
	TypeName("NewThing")
	ContentType("application/json")
	Attributes(func() {
		Attribute("name", String, "The new thing")
		Attribute("definition", String, "What the new thing is")
		Attribute("link", String, "Illustrative link for the new thing")
		Attribute("date", DateTime, "When this was a new thing")
		Attribute("bibliography", ArrayOf(String))
	})
	View("default", func() {
		Attribute("name")
		Attribute("definition")
		Attribute("link")
	})
	View("full", func() {
		Attribute("name")
		Attribute("definition")
		Attribute("link")
		Attribute("date")
		Attribute("bibliography")
	})
	View("name", func() {
		Attribute("name")
	})
	View("name+definition", func() {
		Attribute("name")
		Attribute("definition")
	})
	View("name+link", func() {
		Attribute("name")
		Attribute("link")
	})
})
