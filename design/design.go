package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("discover", func() {
	Title("One-new-thing-today")
	Description("Discover one new thing a day.")
	Produces("application/json")
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
		Result(Int)

		HTTP(func() {
			GET("/neat/thing/today")
		})
	})

	Files("/openapi.json", "../../gen/http/openapi.json")
})
