// Code generated by go-swagger; DO NOT EDIT.

package urcap

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

// NewDeleteUrcapByIDParams creates a new DeleteUrcapByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUrcapByIDParams() *DeleteUrcapByIDParams {
	return &DeleteUrcapByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUrcapByIDParamsWithTimeout creates a new DeleteUrcapByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteUrcapByIDParamsWithTimeout(timeout time.Duration) *DeleteUrcapByIDParams {
	return &DeleteUrcapByIDParams{
		timeout: timeout,
	}
}

// NewDeleteUrcapByIDParamsWithContext creates a new DeleteUrcapByIDParams object
// with the ability to set a context for a request.
func NewDeleteUrcapByIDParamsWithContext(ctx context.Context) *DeleteUrcapByIDParams {
	return &DeleteUrcapByIDParams{
		Context: ctx,
	}
}

// NewDeleteUrcapByIDParamsWithHTTPClient creates a new DeleteUrcapByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUrcapByIDParamsWithHTTPClient(client *http.Client) *DeleteUrcapByIDParams {
	return &DeleteUrcapByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteUrcapByIDParams contains all the parameters to send to the API endpoint

	for the delete urcap by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteUrcapByIDParams struct {

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

// WithDefaults hydrates default values in the delete urcap by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUrcapByIDParams) WithDefaults() *DeleteUrcapByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete urcap by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUrcapByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete urcap by Id params
func (o *DeleteUrcapByIDParams) WithTimeout(timeout time.Duration) *DeleteUrcapByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete urcap by Id params
func (o *DeleteUrcapByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete urcap by Id params
func (o *DeleteUrcapByIDParams) WithContext(ctx context.Context) *DeleteUrcapByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete urcap by Id params
func (o *DeleteUrcapByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete urcap by Id params
func (o *DeleteUrcapByIDParams) WithHTTPClient(client *http.Client) *DeleteUrcapByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete urcap by Id params
func (o *DeleteUrcapByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUrcapID adds the urcapID to the delete urcap by Id params
func (o *DeleteUrcapByIDParams) WithUrcapID(urcapID string) *DeleteUrcapByIDParams {
	o.SetUrcapID(urcapID)
	return o
}

// SetUrcapID adds the urcapId to the delete urcap by Id params
func (o *DeleteUrcapByIDParams) SetUrcapID(urcapID string) {
	o.UrcapID = urcapID
}

// WithVendorID adds the vendorID to the delete urcap by Id params
func (o *DeleteUrcapByIDParams) WithVendorID(vendorID string) *DeleteUrcapByIDParams {
	o.SetVendorID(vendorID)
	return o
}

// SetVendorID adds the vendorId to the delete urcap by Id params
func (o *DeleteUrcapByIDParams) SetVendorID(vendorID string) {
	o.VendorID = vendorID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUrcapByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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