// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Manifest URCap Manifest Specification. The Manifest file is a YAML file defining a multi-artifacts based urcap application. The Manifest.yaml file is located on the top level directory/folder.
//
// swagger:model Manifest
type Manifest struct {

	// Version (major.minor.patch) of the Manifest specification used. Tools not implementing required version MUST reject the configuration file.
	//
	// Example: 1.0.0
	// Required: true
	// Pattern: ^[a-zA-Z0-9._-]+$
	APIVersion string `json:"apiVersion" yaml:"apiVersion"`

	// metadata
	// Required: true
	Metadata *Metadata `json:"metadata" yaml:"metadata"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Manifest) UnmarshalJSON(data []byte) error {
	var props struct {

		// Version (major.minor.patch) of the Manifest specification used. Tools not implementing required version MUST reject the configuration file.
		//
		// Example: 1.0.0
		// Required: true
		// Pattern: ^[a-zA-Z0-9._-]+$
		APIVersion string `json:"apiVersion" yaml:"apiVersion"`

		// metadata
		// Required: true
		Metadata *Metadata `json:"metadata" yaml:"metadata"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.APIVersion = props.APIVersion
	m.Metadata = props.Metadata
	return nil
}

// Validate validates this manifest
func (m *Manifest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Manifest) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.RequiredString("apiVersion", "body", m.APIVersion); err != nil {
		return err
	}

	if err := validate.Pattern("apiVersion", "body", m.APIVersion, `^[a-zA-Z0-9._-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Manifest) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this manifest based on the context it is used
func (m *Manifest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Manifest) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Manifest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Manifest) UnmarshalBinary(b []byte) error {
	var res Manifest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
