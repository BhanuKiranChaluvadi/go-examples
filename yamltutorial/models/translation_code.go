// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TranslationCode translation code
//
// swagger:model TranslationCode
type TranslationCode struct {

	// API context
	Major int64 `json:"major,omitempty" yaml:"major,omitempty"`

	// error context
	Minor int64 `json:"minor,omitempty" yaml:"minor,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *TranslationCode) UnmarshalJSON(data []byte) error {
	var props struct {

		// API context
		Major int64 `json:"major,omitempty" yaml:"major,omitempty"`

		// error context
		Minor int64 `json:"minor,omitempty" yaml:"minor,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Major = props.Major
	m.Minor = props.Minor
	return nil
}

// Validate validates this translation code
func (m *TranslationCode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this translation code based on context it is used
func (m *TranslationCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TranslationCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TranslationCode) UnmarshalBinary(b []byte) error {
	var res TranslationCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
