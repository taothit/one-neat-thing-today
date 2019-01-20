// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Discovery": Application Contexts
//
// Command:
// $ goagen
// --design=douthitlab.edu/one-new-thing-today/design
// --out=$(GOPATH)/src/douthitlab.edu/one-new-thing-today
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// TodayNewThingContext provides the newThing today action context.
type TodayNewThingContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewTodayNewThingContext parses the incoming request URL and body, performs validations and creates the
// context used by the newThing controller today action.
func NewTodayNewThingContext(ctx context.Context, r *http.Request, service *goa.Service) (*TodayNewThingContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := TodayNewThingContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *TodayNewThingContext) OK(r *NewThing) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *TodayNewThingContext) OKFull(r *NewThingFull) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKName sends a HTTP response with status code 200.
func (ctx *TodayNewThingContext) OKName(r *NewThingName) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKNameDefinition sends a HTTP response with status code 200.
func (ctx *TodayNewThingContext) OKNameDefinition(r *NewThingNameDefinition) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKNameLink sends a HTTP response with status code 200.
func (ctx *TodayNewThingContext) OKNameLink(r *NewThingNameLink) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
