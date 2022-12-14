// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStatusContainerByIDParams creates a new StatusContainerByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatusContainerByIDParams() *StatusContainerByIDParams {
	return &StatusContainerByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatusContainerByIDParamsWithTimeout creates a new StatusContainerByIDParams object
// with the ability to set a timeout on a request.
func NewStatusContainerByIDParamsWithTimeout(timeout time.Duration) *StatusContainerByIDParams {
	return &StatusContainerByIDParams{
		timeout: timeout,
	}
}

// NewStatusContainerByIDParamsWithContext creates a new StatusContainerByIDParams object
// with the ability to set a context for a request.
func NewStatusContainerByIDParamsWithContext(ctx context.Context) *StatusContainerByIDParams {
	return &StatusContainerByIDParams{
		Context: ctx,
	}
}

// NewStatusContainerByIDParamsWithHTTPClient creates a new StatusContainerByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatusContainerByIDParamsWithHTTPClient(client *http.Client) *StatusContainerByIDParams {
	return &StatusContainerByIDParams{
		HTTPClient: client,
	}
}

/*
StatusContainerByIDParams contains all the parameters to send to the API endpoint

	for the status container by Id operation.

	Typically these are written to a http.Request.
*/
type StatusContainerByIDParams struct {

	/* ArtifactName.

	   artifact name of container defined in manifest
	*/
	ArtifactName string

	/* UrcapID.

	   urcapID of urcap defined in manifest
	*/
	UrcapID string

	/* VendorID.

	   vendorID of urcap defined in manifest
	*/
	VendorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the status container by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusContainerByIDParams) WithDefaults() *StatusContainerByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the status container by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusContainerByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the status container by Id params
func (o *StatusContainerByIDParams) WithTimeout(timeout time.Duration) *StatusContainerByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status container by Id params
func (o *StatusContainerByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status container by Id params
func (o *StatusContainerByIDParams) WithContext(ctx context.Context) *StatusContainerByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status container by Id params
func (o *StatusContainerByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status container by Id params
func (o *StatusContainerByIDParams) WithHTTPClient(client *http.Client) *StatusContainerByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status container by Id params
func (o *StatusContainerByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactName adds the artifactName to the status container by Id params
func (o *StatusContainerByIDParams) WithArtifactName(artifactName string) *StatusContainerByIDParams {
	o.SetArtifactName(artifactName)
	return o
}

// SetArtifactName adds the artifactName to the status container by Id params
func (o *StatusContainerByIDParams) SetArtifactName(artifactName string) {
	o.ArtifactName = artifactName
}

// WithUrcapID adds the urcapID to the status container by Id params
func (o *StatusContainerByIDParams) WithUrcapID(urcapID string) *StatusContainerByIDParams {
	o.SetUrcapID(urcapID)
	return o
}

// SetUrcapID adds the urcapId to the status container by Id params
func (o *StatusContainerByIDParams) SetUrcapID(urcapID string) {
	o.UrcapID = urcapID
}

// WithVendorID adds the vendorID to the status container by Id params
func (o *StatusContainerByIDParams) WithVendorID(vendorID string) *StatusContainerByIDParams {
	o.SetVendorID(vendorID)
	return o
}

// SetVendorID adds the vendorId to the status container by Id params
func (o *StatusContainerByIDParams) SetVendorID(vendorID string) {
	o.VendorID = vendorID
}

// WriteToRequest writes these params to a swagger request
func (o *StatusContainerByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param artifactName
	if err := r.SetPathParam("artifactName", o.ArtifactName); err != nil {
		return err
	}

	// path param urcapID
	if err := r.SetPathParam("urcapID", o.UrcapID); err != nil {
		return err
	}

	// path param vendorID
	if err := r.SetPathParam("vendorID", o.VendorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
