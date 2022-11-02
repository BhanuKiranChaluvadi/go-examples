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

// Metadata Metadata of the urcap
//
// swagger:model Metadata
type Metadata struct {

	// Specifies copy rights for the urcap
	// Example: Copyright (c) 2009-2021 Universal Robots. All rights reserved.
	Copyright string `json:"copyright,omitempty" yaml:"copyright,omitempty"`

	// Short description of urcap
	// Example: Sample gripper URCap
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// License of various software/hardware used while developing the urcap
	// Example: TODO
	License string `json:"license,omitempty" yaml:"license,omitempty"`

	// Urcap id for this application. The ID must be unique for each urcap application developed by a company (vendorID).
	//
	// Example: dockerdaemon
	// Required: true
	// Max Length: 27
	// Min Length: 2
	// Pattern: ^[a-zA-Z0-9._-]+$
	UrcapID *string `json:"urcapID" yaml:"urcapID"`

	// Urcap name of this application. This Will be displayed on user interface.
	//
	// Example: Docker Daemon
	// Required: true
	// Max Length: 20
	// Min Length: 3
	// Pattern: ^[a-zA-Z0-9_\-\s]+$
	UrcapName *string `json:"urcapName" yaml:"urcapName"`

	// Company urcap developer ID. The ID is shared among all urcap applications developed by company.
	//
	// Example: universal-robots
	// Required: true
	// Max Length: 27
	// Min Length: 2
	// Pattern: ^[a-zA-Z0-9._-]+$
	VendorID *string `json:"vendorID" yaml:"vendorID"`

	// Urcap name of this application. This Will be displayed on user interface.
	//
	// Example: Universal Robots
	// Required: true
	// Max Length: 27
	// Min Length: 2
	// Pattern: ^[a-zA-Z0-9_\-\s]+$
	VendorName *string `json:"vendorName" yaml:"vendorName"`

	// Urcap version (major.minor.patch)
	//
	// Example: 1.0.0
	// Required: true
	// Pattern: ^\d{1}.\d{1}.\d{1}$
	Version *string `json:"version" yaml:"version"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Metadata) UnmarshalJSON(data []byte) error {
	var props struct {

		// Specifies copy rights for the urcap
		// Example: Copyright (c) 2009-2021 Universal Robots. All rights reserved.
		Copyright string `json:"copyright,omitempty" yaml:"copyright,omitempty"`

		// Short description of urcap
		// Example: Sample gripper URCap
		Description string `json:"description,omitempty" yaml:"description,omitempty"`

		// License of various software/hardware used while developing the urcap
		// Example: TODO
		License string `json:"license,omitempty" yaml:"license,omitempty"`

		// Urcap id for this application. The ID must be unique for each urcap application developed by a company (vendorID).
		//
		// Example: dockerdaemon
		// Required: true
		// Max Length: 27
		// Min Length: 2
		// Pattern: ^[a-zA-Z0-9._-]+$
		UrcapID *string `json:"urcapID" yaml:"urcapID"`

		// Urcap name of this application. This Will be displayed on user interface.
		//
		// Example: Docker Daemon
		// Required: true
		// Max Length: 20
		// Min Length: 3
		// Pattern: ^[a-zA-Z0-9_\-\s]+$
		UrcapName *string `json:"urcapName" yaml:"urcapName"`

		// Company urcap developer ID. The ID is shared among all urcap applications developed by company.
		//
		// Example: universal-robots
		// Required: true
		// Max Length: 27
		// Min Length: 2
		// Pattern: ^[a-zA-Z0-9._-]+$
		VendorID *string `json:"vendorID" yaml:"vendorID"`

		// Urcap name of this application. This Will be displayed on user interface.
		//
		// Example: Universal Robots
		// Required: true
		// Max Length: 27
		// Min Length: 2
		// Pattern: ^[a-zA-Z0-9_\-\s]+$
		VendorName *string `json:"vendorName" yaml:"vendorName"`

		// Urcap version (major.minor.patch)
		//
		// Example: 1.0.0
		// Required: true
		// Pattern: ^\d{1}.\d{1}.\d{1}$
		Version *string `json:"version" yaml:"version"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Copyright = props.Copyright
	m.Description = props.Description
	m.License = props.License
	m.UrcapID = props.UrcapID
	m.UrcapName = props.UrcapName
	m.VendorID = props.VendorID
	m.VendorName = props.VendorName
	m.Version = props.Version
	return nil
}

// Validate validates this metadata
func (m *Metadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUrcapID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUrcapName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendorName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Metadata) validateUrcapID(formats strfmt.Registry) error {

	if err := validate.Required("urcapID", "body", m.UrcapID); err != nil {
		return err
	}

	if err := validate.MinLength("urcapID", "body", *m.UrcapID, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("urcapID", "body", *m.UrcapID, 27); err != nil {
		return err
	}

	if err := validate.Pattern("urcapID", "body", *m.UrcapID, `^[a-zA-Z0-9._-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Metadata) validateUrcapName(formats strfmt.Registry) error {

	if err := validate.Required("urcapName", "body", m.UrcapName); err != nil {
		return err
	}

	if err := validate.MinLength("urcapName", "body", *m.UrcapName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("urcapName", "body", *m.UrcapName, 20); err != nil {
		return err
	}

	if err := validate.Pattern("urcapName", "body", *m.UrcapName, `^[a-zA-Z0-9_\-\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Metadata) validateVendorID(formats strfmt.Registry) error {

	if err := validate.Required("vendorID", "body", m.VendorID); err != nil {
		return err
	}

	if err := validate.MinLength("vendorID", "body", *m.VendorID, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("vendorID", "body", *m.VendorID, 27); err != nil {
		return err
	}

	if err := validate.Pattern("vendorID", "body", *m.VendorID, `^[a-zA-Z0-9._-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Metadata) validateVendorName(formats strfmt.Registry) error {

	if err := validate.Required("vendorName", "body", m.VendorName); err != nil {
		return err
	}

	if err := validate.MinLength("vendorName", "body", *m.VendorName, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("vendorName", "body", *m.VendorName, 27); err != nil {
		return err
	}

	if err := validate.Pattern("vendorName", "body", *m.VendorName, `^[a-zA-Z0-9_\-\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Metadata) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.Pattern("version", "body", *m.Version, `^\d{1}.\d{1}.\d{1}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this metadata based on context it is used
func (m *Metadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metadata) UnmarshalBinary(b []byte) error {
	var res Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
