// Code generated by goa v3.0.2, DO NOT EDIT.
//
// neatThing HTTP server types
//
// Command:
// $ goa gen github.com/taothit/one-neat-thing-today/design

package server

import (
	neatthing "github.com/taothit/one-neat-thing-today/gen/neat_thing"
	neatthingviews "github.com/taothit/one-neat-thing-today/gen/neat_thing/views"
	goa "goa.design/goa/v3/pkg"
)

// NewNeatThingRequestBody is the type of the "neatThing" service
// "newNeatThing" endpoint HTTP request body.
type NewNeatThingRequestBody struct {
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// What the neat thing is
	Definition *string `form:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
	// Illustrative link for the neat thing
	Link *string `form:"link,omitempty" json:"link,omitempty" xml:"link,omitempty"`
	// When this was a neat thing
	Date         *string  `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	Bibliography []string `form:"bibliography,omitempty" json:"bibliography,omitempty" xml:"bibliography,omitempty"`
}

// NeatThingTodayResponseBody is the type of the "neatThing" service
// "neatThingToday" endpoint HTTP response body.
type NeatThingTodayResponseBody struct {
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// What the neat thing is
	Definition *string `form:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
	// Illustrative link for the neat thing
	Link *string `form:"link,omitempty" json:"link,omitempty" xml:"link,omitempty"`
}

// NeatThingTodayResponseBodyFull is the type of the "neatThing" service
// "neatThingToday" endpoint HTTP response body.
type NeatThingTodayResponseBodyFull struct {
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// What the neat thing is
	Definition *string `form:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
	// Illustrative link for the neat thing
	Link *string `form:"link,omitempty" json:"link,omitempty" xml:"link,omitempty"`
	// When this was a neat thing
	Date         *string  `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	Bibliography []string `form:"bibliography,omitempty" json:"bibliography,omitempty" xml:"bibliography,omitempty"`
}

// NeatThingTodayResponseBodyName is the type of the "neatThing" service
// "neatThingToday" endpoint HTTP response body.
type NeatThingTodayResponseBodyName struct {
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NeatThingTodayResponseBodyNameDefinition is the type of the "neatThing"
// service "neatThingToday" endpoint HTTP response body.
type NeatThingTodayResponseBodyNameDefinition struct {
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// What the neat thing is
	Definition *string `form:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
}

// NeatThingTodayResponseBodyNameLink is the type of the "neatThing" service
// "neatThingToday" endpoint HTTP response body.
type NeatThingTodayResponseBodyNameLink struct {
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Illustrative link for the neat thing
	Link *string `form:"link,omitempty" json:"link,omitempty" xml:"link,omitempty"`
}

// NewNeatThingResponseBodyFull is the type of the "neatThing" service
// "newNeatThing" endpoint HTTP response body.
type NewNeatThingResponseBodyFull struct {
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// What the neat thing is
	Definition *string `form:"definition,omitempty" json:"definition,omitempty" xml:"definition,omitempty"`
	// Illustrative link for the neat thing
	Link *string `form:"link,omitempty" json:"link,omitempty" xml:"link,omitempty"`
	// When this was a neat thing
	Date         *string  `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	Bibliography []string `form:"bibliography,omitempty" json:"bibliography,omitempty" xml:"bibliography,omitempty"`
}

// NewNeatThingTodayResponseBody builds the HTTP response body from the result
// of the "neatThingToday" endpoint of the "neatThing" service.
func NewNeatThingTodayResponseBody(res *neatthingviews.NeatThingView) *NeatThingTodayResponseBody {
	body := &NeatThingTodayResponseBody{
		Name:       res.Name,
		Definition: res.Definition,
		Link:       res.Link,
	}
	return body
}

// NewNeatThingTodayResponseBodyFull builds the HTTP response body from the
// result of the "neatThingToday" endpoint of the "neatThing" service.
func NewNeatThingTodayResponseBodyFull(res *neatthingviews.NeatThingView) *NeatThingTodayResponseBodyFull {
	body := &NeatThingTodayResponseBodyFull{
		Name:       res.Name,
		Definition: res.Definition,
		Link:       res.Link,
		Date:       res.Date,
	}
	if res.Bibliography != nil {
		body.Bibliography = make([]string, len(res.Bibliography))
		for i, val := range res.Bibliography {
			body.Bibliography[i] = val
		}
	}
	return body
}

// NewNeatThingTodayResponseBodyName builds the HTTP response body from the
// result of the "neatThingToday" endpoint of the "neatThing" service.
func NewNeatThingTodayResponseBodyName(res *neatthingviews.NeatThingView) *NeatThingTodayResponseBodyName {
	body := &NeatThingTodayResponseBodyName{
		Name: res.Name,
	}
	return body
}

// NewNeatThingTodayResponseBodyNameDefinition builds the HTTP response body
// from the result of the "neatThingToday" endpoint of the "neatThing" service.
func NewNeatThingTodayResponseBodyNameDefinition(res *neatthingviews.NeatThingView) *NeatThingTodayResponseBodyNameDefinition {
	body := &NeatThingTodayResponseBodyNameDefinition{
		Name:       res.Name,
		Definition: res.Definition,
	}
	return body
}

// NewNeatThingTodayResponseBodyNameLink builds the HTTP response body from the
// result of the "neatThingToday" endpoint of the "neatThing" service.
func NewNeatThingTodayResponseBodyNameLink(res *neatthingviews.NeatThingView) *NeatThingTodayResponseBodyNameLink {
	body := &NeatThingTodayResponseBodyNameLink{
		Name: res.Name,
		Link: res.Link,
	}
	return body
}

// NewNewNeatThingResponseBodyFull builds the HTTP response body from the
// result of the "newNeatThing" endpoint of the "neatThing" service.
func NewNewNeatThingResponseBodyFull(res *neatthingviews.NeatThingView) *NewNeatThingResponseBodyFull {
	body := &NewNeatThingResponseBodyFull{
		Name:       res.Name,
		Definition: res.Definition,
		Link:       res.Link,
		Date:       res.Date,
	}
	if res.Bibliography != nil {
		body.Bibliography = make([]string, len(res.Bibliography))
		for i, val := range res.Bibliography {
			body.Bibliography[i] = val
		}
	}
	return body
}

// NewNewNeatThingNeatThing builds a neatThing service newNeatThing endpoint
// payload.
func NewNewNeatThingNeatThing(body *NewNeatThingRequestBody) *neatthing.NeatThing {
	v := &neatthing.NeatThing{
		Name:       body.Name,
		Definition: body.Definition,
		Link:       body.Link,
		Date:       body.Date,
	}
	if body.Bibliography != nil {
		v.Bibliography = make([]string, len(body.Bibliography))
		for i, val := range body.Bibliography {
			v.Bibliography[i] = val
		}
	}
	return v
}

// ValidateNewNeatThingRequestBody runs the validations defined on
// NewNeatThingRequestBody
func ValidateNewNeatThingRequestBody(body *NewNeatThingRequestBody) (err error) {
	if body.Link != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.link", *body.Link, goa.FormatURI))
	}
	if body.Date != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date", *body.Date, goa.FormatDateTime))
	}
	return
}
