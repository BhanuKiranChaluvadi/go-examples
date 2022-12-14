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

// NewUpdateUrcapByIDParams creates a new UpdateUrcapByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUrcapByIDParams() *UpdateUrcapByIDParams {
	return &UpdateUrcapByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUrcapByIDParamsWithTimeout creates a new UpdateUrcapByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateUrcapByIDParamsWithTimeout(timeout time.Duration) *UpdateUrcapByIDParams {
	return &UpdateUrcapByIDParams{
		timeout: timeout,
	}
}

// NewUpdateUrcapByIDParamsWithContext creates a new UpdateUrcapByIDParams object
// with the ability to set a context for a request.
func NewUpdateUrcapByIDParamsWithContext(ctx context.Context) *UpdateUrcapByIDParams {
	return &UpdateUrcapByIDParams{
		Context: ctx,
	}
}

// NewUpdateUrcapByIDParamsWithHTTPClient creates a new UpdateUrcapByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUrcapByIDParamsWithHTTPClient(client *http.Client) *UpdateUrcapByIDParams {
	return &UpdateUrcapByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateUrcapByIDParams contains all the parameters to send to the API endpoint

	for the update urcap by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateUrcapByIDParams struct {

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

// WithDefaults hydrates default values in the update urcap by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUrcapByIDParams) WithDefaults() *UpdateUrcapByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update urcap by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUrcapByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update urcap by Id params
func (o *UpdateUrcapByIDParams) WithTimeout(timeout time.Duration) *UpdateUrcapByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update urcap by Id params
func (o *UpdateUrcapByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update urcap by Id params
func (o *UpdateUrcapByIDParams) WithContext(ctx context.Context) *UpdateUrcapByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update urcap by Id params
func (o *UpdateUrcapByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update urcap by Id params
func (o *UpdateUrcapByIDParams) WithHTTPClient(client *http.Client) *UpdateUrcapByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update urcap by Id params
func (o *UpdateUrcapByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUrcapID adds the urcapID to the update urcap by Id params
func (o *UpdateUrcapByIDParams) WithUrcapID(urcapID string) *UpdateUrcapByIDParams {
	o.SetUrcapID(urcapID)
	return o
}

// SetUrcapID adds the urcapId to the update urcap by Id params
func (o *UpdateUrcapByIDParams) SetUrcapID(urcapID string) {
	o.UrcapID = urcapID
}

// WithVendorID adds the vendorID to the update urcap by Id params
func (o *UpdateUrcapByIDParams) WithVendorID(vendorID string) *UpdateUrcapByIDParams {
	o.SetVendorID(vendorID)
	return o
}

// SetVendorID adds the vendorId to the update urcap by Id params
func (o *UpdateUrcapByIDParams) SetVendorID(vendorID string) {
	o.VendorID = vendorID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUrcapByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
