// Code generated by go-swagger; DO NOT EDIT.

package urcap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateUrcapByIDHandlerFunc turns a function with the right signature into a update urcap by Id handler
type UpdateUrcapByIDHandlerFunc func(UpdateUrcapByIDParams) UpdateUrcapByIDResponder

// Handle executing the request and returning a response
func (fn UpdateUrcapByIDHandlerFunc) Handle(params UpdateUrcapByIDParams) UpdateUrcapByIDResponder {
	return fn(params)
}

// UpdateUrcapByIDHandler interface for that can handle valid update urcap by Id params
type UpdateUrcapByIDHandler interface {
	Handle(UpdateUrcapByIDParams) UpdateUrcapByIDResponder
}

// NewUpdateUrcapByID creates a new http.Handler for the update urcap by Id operation
func NewUpdateUrcapByID(ctx *middleware.Context, handler UpdateUrcapByIDHandler) *UpdateUrcapByID {
	return &UpdateUrcapByID{Context: ctx, Handler: handler}
}

/*
	UpdateUrcapByID swagger:route PUT /urcaps/{vendorID}/{urcapID} urcap updateUrcapById

# Updates a urcap info by ID's

This operation updates an exsisting urcap.
*/
type UpdateUrcapByID struct {
	Context *middleware.Context
	Handler UpdateUrcapByIDHandler
}

func (o *UpdateUrcapByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateUrcapByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
