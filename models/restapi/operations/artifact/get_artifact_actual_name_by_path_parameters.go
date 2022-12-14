// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetArtifactActualNameByPathParams creates a new GetArtifactActualNameByPathParams object
//
// There are no default values defined in the spec.
func NewGetArtifactActualNameByPathParams() GetArtifactActualNameByPathParams {

	return GetArtifactActualNameByPathParams{}
}

// GetArtifactActualNameByPathParams contains all the bound params for the get artifact actual name by path operation
// typically these are obtained from a http.Request
//
// swagger:parameters getArtifactActualNameByPath
type GetArtifactActualNameByPathParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Path to the resource/artifact on host machine
	  Required: true
	  In: query
	*/
	ArtifactPath string
	/*The type of the artifact
	  Required: true
	  In: path
	*/
	ArtifactType string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetArtifactActualNameByPathParams() beforehand.
func (o *GetArtifactActualNameByPathParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qArtifactPath, qhkArtifactPath, _ := qs.GetOK("artifactPath")
	if err := o.bindArtifactPath(qArtifactPath, qhkArtifactPath, route.Formats); err != nil {
		res = append(res, err)
	}

	rArtifactType, rhkArtifactType, _ := route.Params.GetOK("artifactType")
	if err := o.bindArtifactType(rArtifactType, rhkArtifactType, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindArtifactPath binds and validates parameter ArtifactPath from query.
func (o *GetArtifactActualNameByPathParams) bindArtifactPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("artifactPath", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("artifactPath", "query", raw); err != nil {
		return err
	}
	o.ArtifactPath = raw

	return nil
}

// bindArtifactType binds and validates parameter ArtifactType from path.
func (o *GetArtifactActualNameByPathParams) bindArtifactType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ArtifactType = raw

	if err := o.validateArtifactType(formats); err != nil {
		return err
	}

	return nil
}

// validateArtifactType carries on validations for parameter ArtifactType
func (o *GetArtifactActualNameByPathParams) validateArtifactType(formats strfmt.Registry) error {

	if err := validate.EnumCase("artifactType", "path", o.ArtifactType, []interface{}{"container", "polyscopeBundle", "webArchieve"}, true); err != nil {
		return err
	}

	return nil
}
