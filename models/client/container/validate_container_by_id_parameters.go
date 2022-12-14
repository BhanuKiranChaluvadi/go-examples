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

// NewValidateContainerByIDParams creates a new ValidateContainerByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateContainerByIDParams() *ValidateContainerByIDParams {
	return &ValidateContainerByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateContainerByIDParamsWithTimeout creates a new ValidateContainerByIDParams object
// with the ability to set a timeout on a request.
func NewValidateContainerByIDParamsWithTimeout(timeout time.Duration) *ValidateContainerByIDParams {
	return &ValidateContainerByIDParams{
		timeout: timeout,
	}
}

// NewValidateContainerByIDParamsWithContext creates a new ValidateContainerByIDParams object
// with the ability to set a context for a request.
func NewValidateContainerByIDParamsWithContext(ctx context.Context) *ValidateContainerByIDParams {
	return &ValidateContainerByIDParams{
		Context: ctx,
	}
}

// NewValidateContainerByIDParamsWithHTTPClient creates a new ValidateContainerByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateContainerByIDParamsWithHTTPClient(client *http.Client) *ValidateContainerByIDParams {
	return &ValidateContainerByIDParams{
		HTTPClient: client,
	}
}

/*
ValidateContainerByIDParams contains all the parameters to send to the API endpoint

	for the validate container by Id operation.

	Typically these are written to a http.Request.
*/
type ValidateContainerByIDParams struct {

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

// WithDefaults hydrates default values in the validate container by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateContainerByIDParams) WithDefaults() *ValidateContainerByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate container by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateContainerByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate container by Id params
func (o *ValidateContainerByIDParams) WithTimeout(timeout time.Duration) *ValidateContainerByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate container by Id params
func (o *ValidateContainerByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate container by Id params
func (o *ValidateContainerByIDParams) WithContext(ctx context.Context) *ValidateContainerByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate container by Id params
func (o *ValidateContainerByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate container by Id params
func (o *ValidateContainerByIDParams) WithHTTPClient(client *http.Client) *ValidateContainerByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate container by Id params
func (o *ValidateContainerByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactName adds the artifactName to the validate container by Id params
func (o *ValidateContainerByIDParams) WithArtifactName(artifactName string) *ValidateContainerByIDParams {
	o.SetArtifactName(artifactName)
	return o
}

// SetArtifactName adds the artifactName to the validate container by Id params
func (o *ValidateContainerByIDParams) SetArtifactName(artifactName string) {
	o.ArtifactName = artifactName
}

// WithUrcapID adds the urcapID to the validate container by Id params
func (o *ValidateContainerByIDParams) WithUrcapID(urcapID string) *ValidateContainerByIDParams {
	o.SetUrcapID(urcapID)
	return o
}

// SetUrcapID adds the urcapId to the validate container by Id params
func (o *ValidateContainerByIDParams) SetUrcapID(urcapID string) {
	o.UrcapID = urcapID
}

// WithVendorID adds the vendorID to the validate container by Id params
func (o *ValidateContainerByIDParams) WithVendorID(vendorID string) *ValidateContainerByIDParams {
	o.SetVendorID(vendorID)
	return o
}

// SetVendorID adds the vendorId to the validate container by Id params
func (o *ValidateContainerByIDParams) SetVendorID(vendorID string) {
	o.VendorID = vendorID
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateContainerByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
