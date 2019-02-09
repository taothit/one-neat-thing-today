package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Discovery", func() {
	Title("A neat thing each day")
	Description("Discover something neat each day by being given an item from a curated list.")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("neatThing", func() {
	Description("Something neat for you")
	DefaultMedia(NeatThingMedia)
	BasePath("/neat/thing")
	Action("today", func() {
		Routing(GET("today"))
		Description("GET the neat thing for the day")
		Response(OK, NeatThingMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
})

// NeatThingMedia represents a neat thing with its name, description/definition,
// and a great illustrative link.
var NeatThingMedia = MediaType("application/vnd.douthitlab.neatthing", func() {
	TypeName("NeatThing")
	ContentType("application/json")
	Attributes(func() {
		Attribute("name", String, "The neat thing")
		Attribute("definition", String, "What the neat thing is")
		Attribute("link", String, "Illustrative link for the neat thing")
		Attribute("date", DateTime, "When this was a neat thing")
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
