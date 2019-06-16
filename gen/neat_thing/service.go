// Code generated by goa v3.0.2, DO NOT EDIT.
//
// neatThing service
//
// Command:
// $ goa gen github.com/taothit/one-neat-thing-today/design

package neatthing

import (
	"context"

	neatthingviews "github.com/taothit/one-neat-thing-today/gen/neat_thing/views"
)

// Allows access to discover neat things.
type Service interface {
	// NeatThingToday implements neatThingToday.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "full"
	//	- "name"
	//	- "name+definition"
	//	- "name+link"
	NeatThingToday(context.Context) (res *NeatThing, view string, err error)
	// NewNeatThing implements newNeatThing.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "full"
	//	- "name"
	//	- "name+definition"
	//	- "name+link"
	NewNeatThing(context.Context, *NeatThing) (res *NeatThing, view string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "neatThing"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"neatThingToday", "newNeatThing"}

// NeatThing is the result type of the neatThing service neatThingToday method.
type NeatThing struct {
	// The neat thing
	Name *string
	// What the neat thing is
	Definition *string
	// Illustrative link for the neat thing
	Link *string
	// When this was a neat thing
	Date         *string
	Bibliography []string
}

// NewNeatThing initializes result type NeatThing from viewed result type
// NeatThing.
func NewNeatThing(vres *neatthingviews.NeatThing) *NeatThing {
	var res *NeatThing
	switch vres.View {
	case "default", "":
		res = newNeatThing(vres.Projected)
	case "full":
		res = newNeatThingFull(vres.Projected)
	case "name":
		res = newNeatThingName(vres.Projected)
	case "name+definition":
		res = newNeatThingNameDefinition(vres.Projected)
	case "name+link":
		res = newNeatThingNameLink(vres.Projected)
	}
	return res
}

// NewViewedNeatThing initializes viewed result type NeatThing from result type
// NeatThing using the given view.
func NewViewedNeatThing(res *NeatThing, view string) *neatthingviews.NeatThing {
	var vres *neatthingviews.NeatThing
	switch view {
	case "default", "":
		p := newNeatThingView(res)
		vres = &neatthingviews.NeatThing{p, "default"}
	case "full":
		p := newNeatThingViewFull(res)
		vres = &neatthingviews.NeatThing{p, "full"}
	case "name":
		p := newNeatThingViewName(res)
		vres = &neatthingviews.NeatThing{p, "name"}
	case "name+definition":
		p := newNeatThingViewNameDefinition(res)
		vres = &neatthingviews.NeatThing{p, "name+definition"}
	case "name+link":
		p := newNeatThingViewNameLink(res)
		vres = &neatthingviews.NeatThing{p, "name+link"}
	}
	return vres
}

// newNeatThing converts projected type NeatThing to service type NeatThing.
func newNeatThing(vres *neatthingviews.NeatThingView) *NeatThing {
	res := &NeatThing{
		Name:       vres.Name,
		Definition: vres.Definition,
		Link:       vres.Link,
	}
	return res
}

// newNeatThingFull converts projected type NeatThing to service type NeatThing.
func newNeatThingFull(vres *neatthingviews.NeatThingView) *NeatThing {
	res := &NeatThing{
		Name:       vres.Name,
		Definition: vres.Definition,
		Link:       vres.Link,
		Date:       vres.Date,
	}
	if vres.Bibliography != nil {
		res.Bibliography = make([]string, len(vres.Bibliography))
		for i, val := range vres.Bibliography {
			res.Bibliography[i] = val
		}
	}
	return res
}

// newNeatThingName converts projected type NeatThing to service type NeatThing.
func newNeatThingName(vres *neatthingviews.NeatThingView) *NeatThing {
	res := &NeatThing{
		Name: vres.Name,
	}
	return res
}

// newNeatThingNameDefinition converts projected type NeatThing to service type
// NeatThing.
func newNeatThingNameDefinition(vres *neatthingviews.NeatThingView) *NeatThing {
	res := &NeatThing{
		Name:       vres.Name,
		Definition: vres.Definition,
	}
	return res
}

// newNeatThingNameLink converts projected type NeatThing to service type
// NeatThing.
func newNeatThingNameLink(vres *neatthingviews.NeatThingView) *NeatThing {
	res := &NeatThing{
		Name: vres.Name,
		Link: vres.Link,
	}
	return res
}

// newNeatThingView projects result type NeatThing to projected type
// NeatThingView using the "default" view.
func newNeatThingView(res *NeatThing) *neatthingviews.NeatThingView {
	vres := &neatthingviews.NeatThingView{
		Name:       res.Name,
		Definition: res.Definition,
		Link:       res.Link,
	}
	return vres
}

// newNeatThingViewFull projects result type NeatThing to projected type
// NeatThingView using the "full" view.
func newNeatThingViewFull(res *NeatThing) *neatthingviews.NeatThingView {
	vres := &neatthingviews.NeatThingView{
		Name:       res.Name,
		Definition: res.Definition,
		Link:       res.Link,
		Date:       res.Date,
	}
	if res.Bibliography != nil {
		vres.Bibliography = make([]string, len(res.Bibliography))
		for i, val := range res.Bibliography {
			vres.Bibliography[i] = val
		}
	}
	return vres
}

// newNeatThingViewName projects result type NeatThing to projected type
// NeatThingView using the "name" view.
func newNeatThingViewName(res *NeatThing) *neatthingviews.NeatThingView {
	vres := &neatthingviews.NeatThingView{
		Name: res.Name,
	}
	return vres
}

// newNeatThingViewNameDefinition projects result type NeatThing to projected
// type NeatThingView using the "name+definition" view.
func newNeatThingViewNameDefinition(res *NeatThing) *neatthingviews.NeatThingView {
	vres := &neatthingviews.NeatThingView{
		Name:       res.Name,
		Definition: res.Definition,
	}
	return vres
}

// newNeatThingViewNameLink projects result type NeatThing to projected type
// NeatThingView using the "name+link" view.
func newNeatThingViewNameLink(res *NeatThing) *neatthingviews.NeatThingView {
	vres := &neatthingviews.NeatThingView{
		Name: res.Name,
		Link: res.Link,
	}
	return vres
}
