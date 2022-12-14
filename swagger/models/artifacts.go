// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Artifacts All artifacts like containers, polyscopeBundle, webArchive that makes up an urcap.
//
// swagger:model artifacts
type Artifacts struct {

	// List of container artifacts
	Containers []*Container `json:"containers" yaml:"containers"`

	// List of osgi polyscope bundles
	PolyscopeBundles []*PolyscopeBundle `json:"polyscopeBundles" yaml:"polyscopeBundles"`

	// List of web archives
	WebArchives []*WebArchive `json:"webArchives" yaml:"webArchives"`
}

// Validate validates this artifacts
func (m *Artifacts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolyscopeBundles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebArchives(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Artifacts) validateContainers(formats strfmt.Registry) error {
	if swag.IsZero(m.Containers) { // not required
		return nil
	}

	for i := 0; i < len(m.Containers); i++ {
		if swag.IsZero(m.Containers[i]) { // not required
			continue
		}

		if m.Containers[i] != nil {
			if err := m.Containers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Artifacts) validatePolyscopeBundles(formats strfmt.Registry) error {
	if swag.IsZero(m.PolyscopeBundles) { // not required
		return nil
	}

	for i := 0; i < len(m.PolyscopeBundles); i++ {
		if swag.IsZero(m.PolyscopeBundles[i]) { // not required
			continue
		}

		if m.PolyscopeBundles[i] != nil {
			if err := m.PolyscopeBundles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("polyscopeBundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("polyscopeBundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Artifacts) validateWebArchives(formats strfmt.Registry) error {
	if swag.IsZero(m.WebArchives) { // not required
		return nil
	}

	for i := 0; i < len(m.WebArchives); i++ {
		if swag.IsZero(m.WebArchives[i]) { // not required
			continue
		}

		if m.WebArchives[i] != nil {
			if err := m.WebArchives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webArchives" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webArchives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this artifacts based on the context it is used
func (m *Artifacts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolyscopeBundles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebArchives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Artifacts) contextValidateContainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Containers); i++ {

		if m.Containers[i] != nil {
			if err := m.Containers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Artifacts) contextValidatePolyscopeBundles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PolyscopeBundles); i++ {

		if m.PolyscopeBundles[i] != nil {
			if err := m.PolyscopeBundles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("polyscopeBundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("polyscopeBundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Artifacts) contextValidateWebArchives(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WebArchives); i++ {

		if m.WebArchives[i] != nil {
			if err := m.WebArchives[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webArchives" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webArchives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Artifacts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Artifacts) UnmarshalBinary(b []byte) error {
	var res Artifacts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
