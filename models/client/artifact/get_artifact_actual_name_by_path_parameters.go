// Code generated by go-swagger; DO NOT EDIT.

package artifact

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

// NewGetArtifactActualNameByPathParams creates a new GetArtifactActualNameByPathParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetArtifactActualNameByPathParams() *GetArtifactActualNameByPathParams {
	return &GetArtifactActualNameByPathParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetArtifactActualNameByPathParamsWithTimeout creates a new GetArtifactActualNameByPathParams object
// with the ability to set a timeout on a request.
func NewGetArtifactActualNameByPathParamsWithTimeout(timeout time.Duration) *GetArtifactActualNameByPathParams {
	return &GetArtifactActualNameByPathParams{
		timeout: timeout,
	}
}

// NewGetArtifactActualNameByPathParamsWithContext creates a new GetArtifactActualNameByPathParams object
// with the ability to set a context for a request.
func NewGetArtifactActualNameByPathParamsWithContext(ctx context.Context) *GetArtifactActualNameByPathParams {
	return &GetArtifactActualNameByPathParams{
		Context: ctx,
	}
}

// NewGetArtifactActualNameByPathParamsWithHTTPClient creates a new GetArtifactActualNameByPathParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetArtifactActualNameByPathParamsWithHTTPClient(client *http.Client) *GetArtifactActualNameByPathParams {
	return &GetArtifactActualNameByPathParams{
		HTTPClient: client,
	}
}

/*
GetArtifactActualNameByPathParams contains all the parameters to send to the API endpoint

	for the get artifact actual name by path operation.

	Typically these are written to a http.Request.
*/
type GetArtifactActualNameByPathParams struct {

	/* ArtifactPath.

	   Path to the resource/artifact on host machine
	*/
	ArtifactPath string

	/* ArtifactType.

	   The type of the artifact
	*/
	ArtifactType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get artifact actual name by path params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArtifactActualNameByPathParams) WithDefaults() *GetArtifactActualNameByPathParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get artifact actual name by path params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArtifactActualNameByPathParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get artifact actual name by path params
func (o *GetArtifactActualNameByPathParams) WithTimeout(timeout time.Duration) *GetArtifactActualNameByPathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get artifact actual name by path params
func (o *GetArtifactActualNameByPathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get artifact actual name by path params
func (o *GetArtifactActualNameByPathParams) WithContext(ctx context.Context) *GetArtifactActualNameByPathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get artifact actual name by path params
func (o *GetArtifactActualNameByPathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get artifact actual name by path params
func (o *GetArtifactActualNameByPathParams) WithHTTPClient(client *http.Client) *GetArtifactActualNameByPathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get artifact actual name by path params
func (o *GetArtifactActualNameByPathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactPath adds the artifactPath to the get artifact actual name by path params
func (o *GetArtifactActualNameByPathParams) WithArtifactPath(artifactPath string) *GetArtifactActualNameByPathParams {
	o.SetArtifactPath(artifactPath)
	return o
}

// SetArtifactPath adds the artifactPath to the get artifact actual name by path params
func (o *GetArtifactActualNameByPathParams) SetArtifactPath(artifactPath string) {
	o.ArtifactPath = artifactPath
}

// WithArtifactType adds the artifactType to the get artifact actual name by path params
func (o *GetArtifactActualNameByPathParams) WithArtifactType(artifactType string) *GetArtifactActualNameByPathParams {
	o.SetArtifactType(artifactType)
	return o
}

// SetArtifactType adds the artifactType to the get artifact actual name by path params
func (o *GetArtifactActualNameByPathParams) SetArtifactType(artifactType string) {
	o.ArtifactType = artifactType
}

// WriteToRequest writes these params to a swagger request
func (o *GetArtifactActualNameByPathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param artifactPath
	qrArtifactPath := o.ArtifactPath
	qArtifactPath := qrArtifactPath
	if qArtifactPath != "" {

		if err := r.SetQueryParam("artifactPath", qArtifactPath); err != nil {
			return err
		}
	}

	// path param artifactType
	if err := r.SetPathParam("artifactType", o.ArtifactType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}