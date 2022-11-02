// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIError API error
//
// swagger:model APIError
type APIError struct {

	// The UUID to uniquely identifying the request
	// Example: 9daee671-916a-4678-850b-10b911f0236d
	Trace string `json:"trace,omitempty" yaml:"trace,omitempty"`

	// HTTP status code used for the response
	// Example: 500
	StatusCode uint16 `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`

	// translation code
	TranslationCode *TranslationCode `json:"translationCode,omitempty" yaml:"translationCode,omitempty"`

	// This field contains a snake case string succiently identifying the problem from api context.
	// Example: installation_failed
	// Required: true
	APIContext string `json:"apiContext" yaml:"apiContext"`

	// This field contains description of the problem
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// This field explains a possible solution to the problem
	// Example: Please download urcapx file again and try installing again or contact urcap developer
	Resolve string `json:"resolve,omitempty" yaml:"resolve,omitempty"`

	// This filed contains a publicly-accessible URL where information about the error can be read in a web browser
	// Example: https://docs.api.example.com/v2/urcaps/installation#corrupted_file
	MoreInfo string `json:"moreInfo,omitempty" yaml:"moreInfo,omitempty"`

	// List of errors
	// Required: true
	Errors []*Error `json:"errors" yaml:"errors"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *APIError) UnmarshalJSON(data []byte) error {
	var props struct {

		// The UUID to uniquely identifying the request
		// Example: 9daee671-916a-4678-850b-10b911f0236d
		Trace string `json:"trace,omitempty" yaml:"trace,omitempty"`

		// HTTP status code used for the response
		// Example: 500
		StatusCode uint16 `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`

		// translation code
		TranslationCode *TranslationCode `json:"translationCode,omitempty" yaml:"translationCode,omitempty"`

		// This field contains a snake case string succiently identifying the problem from api context.
		// Example: installation_failed
		// Required: true
		APIContext string `json:"apiContext" yaml:"apiContext"`

		// This field contains description of the problem
		Description string `json:"description,omitempty" yaml:"description,omitempty"`

		// This field explains a possible solution to the problem
		// Example: Please download urcapx file again and try installing again or contact urcap developer
		Resolve string `json:"resolve,omitempty" yaml:"resolve,omitempty"`

		// This filed contains a publicly-accessible URL where information about the error can be read in a web browser
		// Example: https://docs.api.example.com/v2/urcaps/installation#corrupted_file
		MoreInfo string `json:"moreInfo,omitempty" yaml:"moreInfo,omitempty"`

		// List of errors
		// Required: true
		Errors []*Error `json:"errors" yaml:"errors"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Trace = props.Trace
	m.StatusCode = props.StatusCode
	m.TranslationCode = props.TranslationCode
	m.APIContext = props.APIContext
	m.Description = props.Description
	m.Resolve = props.Resolve
	m.MoreInfo = props.MoreInfo
	m.Errors = props.Errors
	return nil
}

// Validate validates this API error
func (m *APIError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTranslationCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIError) validateTranslationCode(formats strfmt.Registry) error {
	if swag.IsZero(m.TranslationCode) { // not required
		return nil
	}

	if m.TranslationCode != nil {
		if err := m.TranslationCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("translationCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("translationCode")
			}
			return err
		}
	}

	return nil
}

func (m *APIError) validateAPIContext(formats strfmt.Registry) error {

	if err := validate.RequiredString("apiContext", "body", m.APIContext); err != nil {
		return err
	}

	return nil
}

func (m *APIError) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this API error based on the context it is used
func (m *APIError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTranslationCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIError) contextValidateTranslationCode(ctx context.Context, formats strfmt.Registry) error {

	if m.TranslationCode != nil {
		if err := m.TranslationCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("translationCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("translationCode")
			}
			return err
		}
	}

	return nil
}

func (m *APIError) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIError) UnmarshalBinary(b []byte) error {
	var res APIError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
