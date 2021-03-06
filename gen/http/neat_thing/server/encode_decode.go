// Code generated by goa v3.0.2, DO NOT EDIT.
//
// neatThing HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/taothit/one-neat-thing-today/design

package server

import (
	"context"
	"io"
	"net/http"

	neatthingviews "github.com/taothit/one-neat-thing-today/gen/neat_thing/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeNeatThingTodayResponse returns an encoder for responses returned by
// the neatThing neatThingToday endpoint.
func EncodeNeatThingTodayResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*neatthingviews.NeatThing)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewNeatThingTodayResponseBody(res.Projected)
		case "full":
			body = NewNeatThingTodayResponseBodyFull(res.Projected)
		case "name":
			body = NewNeatThingTodayResponseBodyName(res.Projected)
		case "name+definition":
			body = NewNeatThingTodayResponseBodyNameDefinition(res.Projected)
		case "name+link":
			body = NewNeatThingTodayResponseBodyNameLink(res.Projected)
		}
		w.WriteHeader(http.StatusNoContent)
		return enc.Encode(body)
	}
}

// EncodeNewNeatThingResponse returns an encoder for responses returned by the
// neatThing newNeatThing endpoint.
func EncodeNewNeatThingResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*neatthingviews.NeatThing)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		body := NewNewNeatThingResponseBodyFull(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeNewNeatThingRequest returns a decoder for requests sent to the
// neatThing newNeatThing endpoint.
func DecodeNewNeatThingRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body NewNeatThingRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateNewNeatThingRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewNewNeatThingNeatThing(&body)

		return payload, nil
	}
}
