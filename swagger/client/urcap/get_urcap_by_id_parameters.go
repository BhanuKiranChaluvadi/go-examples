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

// NewGetUrcapByIDParams creates a new GetUrcapByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUrcapByIDParams() *GetUrcapByIDParams {
	return &GetUrcapByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUrcapByIDParamsWithTimeout creates a new GetUrcapByIDParams object
// with the ability to set a timeout on a request.
func NewGetUrcapByIDParamsWithTimeout(timeout time.Duration) *GetUrcapByIDParams {
	return &GetUrcapByIDParams{
		timeout: timeout,
	}
}

// NewGetUrcapByIDParamsWithContext creates a new GetUrcapByIDParams object
// with the ability to set a context for a request.
func NewGetUrcapByIDParamsWithContext(ctx context.Context) *GetUrcapByIDParams {
	return &GetUrcapByIDParams{
		Context: ctx,
	}
}

// NewGetUrcapByIDParamsWithHTTPClient creates a new GetUrcapByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUrcapByIDParamsWithHTTPClient(client *http.Client) *GetUrcapByIDParams {
	return &GetUrcapByIDParams{
		HTTPClient: client,
	}
}

/*
GetUrcapByIDParams contains all the parameters to send to the API endpoint

	for the get urcap by Id operation.

	Typically these are written to a http.Request.
*/
type GetUrcapByIDParams struct {

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

// WithDefaults hydrates default values in the get urcap by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUrcapByIDParams) WithDefaults() *GetUrcapByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get urcap by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUrcapByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get urcap by Id params
func (o *GetUrcapByIDParams) WithTimeout(timeout time.Duration) *GetUrcapByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get urcap by Id params
func (o *GetUrcapByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get urcap by Id params
func (o *GetUrcapByIDParams) WithContext(ctx context.Context) *GetUrcapByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get urcap by Id params
func (o *GetUrcapByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get urcap by Id params
func (o *GetUrcapByIDParams) WithHTTPClient(client *http.Client) *GetUrcapByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get urcap by Id params
func (o *GetUrcapByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUrcapID adds the urcapID to the get urcap by Id params
func (o *GetUrcapByIDParams) WithUrcapID(urcapID string) *GetUrcapByIDParams {
	o.SetUrcapID(urcapID)
	return o
}

// SetUrcapID adds the urcapId to the get urcap by Id params
func (o *GetUrcapByIDParams) SetUrcapID(urcapID string) {
	o.UrcapID = urcapID
}

// WithVendorID adds the vendorID to the get urcap by Id params
func (o *GetUrcapByIDParams) WithVendorID(vendorID string) *GetUrcapByIDParams {
	o.SetVendorID(vendorID)
	return o
}

// SetVendorID adds the vendorId to the get urcap by Id params
func (o *GetUrcapByIDParams) SetVendorID(vendorID string) {
	o.VendorID = vendorID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUrcapByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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