package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("discover", func() {
	Title("One-new-thing-today")
	Description("Discover one new thing a day.")
	Server("discoverysvr", func() {
		Services("neatThing")
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("neatThing", func() {
	Description("Allows access to discover neat things.")

	Method("neatThingToday", func() {
		HTTP(func() {
			GET("/neat/thing/today")
		})

		Result(NeatThingMedia)
	})

	Files("/openapi.json", "../../gen/http/openapi.json")
})

// NeatThingMedia represents a neat thing with its name, description/definition,
// and a great illustrative link.
var NeatThingMedia = Type("application/vnd.douthitlab.neatthing", func() {
	TypeName("NeatThing")
	Field(1, "name", String, "The neat thing")
	Field(2, "definition", String, "What the neat thing is")
	Field(3, "link", String, "Illustrative link for the neat thing", func() {
		Format(FormatURI)
	})
	Field(4, "date", String, "When this was a neat thing", func() {
		Format(FormatDateTime)
	})
	Field(5, "bibliography", ArrayOf(String))

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
