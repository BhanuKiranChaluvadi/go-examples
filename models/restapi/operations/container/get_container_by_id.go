// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetContainerByIDHandlerFunc turns a function with the right signature into a get container by Id handler
type GetContainerByIDHandlerFunc func(GetContainerByIDParams) GetContainerByIDResponder

// Handle executing the request and returning a response
func (fn GetContainerByIDHandlerFunc) Handle(params GetContainerByIDParams) GetContainerByIDResponder {
	return fn(params)
}

// GetContainerByIDHandler interface for that can handle valid get container by Id params
type GetContainerByIDHandler interface {
	Handle(GetContainerByIDParams) GetContainerByIDResponder
}

// NewGetContainerByID creates a new http.Handler for the get container by Id operation
func NewGetContainerByID(ctx *middleware.Context, handler GetContainerByIDHandler) *GetContainerByID {
	return &GetContainerByID{Context: ctx, Handler: handler}
}

/*
	GetContainerByID swagger:route GET /containers/{vendorID}/{urcapID}/{artifactName} container getContainerById

# Gets a container info by urcap ID's and artifactName

TODO
*/
type GetContainerByID struct {
	Context *middleware.Context
	Handler GetContainerByIDHandler
}

func (o *GetContainerByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetContainerByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}