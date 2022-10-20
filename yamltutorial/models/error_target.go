// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ErrorTarget Error Target Model. This field MAY contain an error target model. Otherwise, it MUST be omitted.
//
// swagger:model ErrorTarget
type ErrorTarget struct {

	// This field MUST contain field, parameter, or header
	// Example: field
	// Enum: [field parameter header]
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// This field MUST contain the name of the problematic field (with dot-syntax if necessary), query parameter, or header.
	// Example: username
	// Required: true
	Name string `json:"name" yaml:"name"`

	// value
	// Example: universal-robots
	// Required: true
	Value string `json:"value" yaml:"value"`
}

// Validate validates this error target
func (m *ErrorTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var errorTargetTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["field","parameter","header"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorTargetTypeTypePropEnum = append(errorTargetTypeTypePropEnum, v)
	}
}

const (

	// ErrorTargetTypeField captures enum value "field"
	ErrorTargetTypeField string = "field"

	// ErrorTargetTypeParameter captures enum value "parameter"
	ErrorTargetTypeParameter string = "parameter"

	// ErrorTargetTypeHeader captures enum value "header"
	ErrorTargetTypeHeader string = "header"
)

// prop value enum
func (m *ErrorTarget) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, errorTargetTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ErrorTarget) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ErrorTarget) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ErrorTarget) validateValue(formats strfmt.Registry) error {

	if err := validate.RequiredString("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error target based on context it is used
func (m *ErrorTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorTarget) UnmarshalBinary(b []byte) error {
	var res ErrorTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
