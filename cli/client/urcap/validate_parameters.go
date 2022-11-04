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

// NewValidateParams creates a new ValidateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateParams() *ValidateParams {
	return &ValidateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateParamsWithTimeout creates a new ValidateParams object
// with the ability to set a timeout on a request.
func NewValidateParamsWithTimeout(timeout time.Duration) *ValidateParams {
	return &ValidateParams{
		timeout: timeout,
	}
}

// NewValidateParamsWithContext creates a new ValidateParams object
// with the ability to set a context for a request.
func NewValidateParamsWithContext(ctx context.Context) *ValidateParams {
	return &ValidateParams{
		Context: ctx,
	}
}

// NewValidateParamsWithHTTPClient creates a new ValidateParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateParamsWithHTTPClient(client *http.Client) *ValidateParams {
	return &ValidateParams{
		HTTPClient: client,
	}
}

/*
ValidateParams contains all the parameters to send to the API endpoint

	for the validate operation.

	Typically these are written to a http.Request.
*/
type ValidateParams struct {

	/* Path.

	   path to built urcap file <urcap_name>.urcapx
	*/
	Path *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateParams) WithDefaults() *ValidateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate params
func (o *ValidateParams) WithTimeout(timeout time.Duration) *ValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate params
func (o *ValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate params
func (o *ValidateParams) WithContext(ctx context.Context) *ValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate params
func (o *ValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate params
func (o *ValidateParams) WithHTTPClient(client *http.Client) *ValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate params
func (o *ValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the validate params
func (o *ValidateParams) WithPath(path *string) *ValidateParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the validate params
func (o *ValidateParams) SetPath(path *string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Path != nil {

		// query param path
		var qrPath string

		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}