// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// StartContainerByIDHandlerFunc turns a function with the right signature into a start container by Id handler
type StartContainerByIDHandlerFunc func(StartContainerByIDParams) StartContainerByIDResponder

// Handle executing the request and returning a response
func (fn StartContainerByIDHandlerFunc) Handle(params StartContainerByIDParams) StartContainerByIDResponder {
	return fn(params)
}

// StartContainerByIDHandler interface for that can handle valid start container by Id params
type StartContainerByIDHandler interface {
	Handle(StartContainerByIDParams) StartContainerByIDResponder
}

// NewStartContainerByID creates a new http.Handler for the start container by Id operation
func NewStartContainerByID(ctx *middleware.Context, handler StartContainerByIDHandler) *StartContainerByID {
	return &StartContainerByID{Context: ctx, Handler: handler}
}

/*
	StartContainerByID swagger:route PUT /containers/{vendorID}/{urcapID}/{artifactName}/start container startContainerById

# Start a container

TODO
*/
type StartContainerByID struct {
	Context *middleware.Context
	Handler StartContainerByIDHandler
}

func (o *StartContainerByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewStartContainerByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
