// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewStopContainerByIDParams creates a new StopContainerByIDParams object
//
// There are no default values defined in the spec.
func NewStopContainerByIDParams() StopContainerByIDParams {

	return StopContainerByIDParams{}
}

// StopContainerByIDParams contains all the bound params for the stop container by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters stopContainerById
type StopContainerByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*artifact name of container defined in manifest
	  Required: true
	  In: path
	*/
	ArtifactName string
	/*urcapID of urcap defined in manifest
	  Required: true
	  In: path
	*/
	UrcapID string
	/*vendorID of urcap defined in manifest
	  Required: true
	  In: path
	*/
	VendorID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewStopContainerByIDParams() beforehand.
func (o *StopContainerByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rArtifactName, rhkArtifactName, _ := route.Params.GetOK("artifactName")
	if err := o.bindArtifactName(rArtifactName, rhkArtifactName, route.Formats); err != nil {
		res = append(res, err)
	}

	rUrcapID, rhkUrcapID, _ := route.Params.GetOK("urcapID")
	if err := o.bindUrcapID(rUrcapID, rhkUrcapID, route.Formats); err != nil {
		res = append(res, err)
	}

	rVendorID, rhkVendorID, _ := route.Params.GetOK("vendorID")
	if err := o.bindVendorID(rVendorID, rhkVendorID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindArtifactName binds and validates parameter ArtifactName from path.
func (o *StopContainerByIDParams) bindArtifactName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ArtifactName = raw

	return nil
}

// bindUrcapID binds and validates parameter UrcapID from path.
func (o *StopContainerByIDParams) bindUrcapID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.UrcapID = raw

	return nil
}

// bindVendorID binds and validates parameter VendorID from path.
func (o *StopContainerByIDParams) bindVendorID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.VendorID = raw

	return nil
}