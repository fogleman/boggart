// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Timezone timezone
// swagger:model Timezone
type Timezone struct {

	// display string
	// Required: true
	DisplayString *string `json:"display_string"`

	// offset
	// Required: true
	Offset *int64 `json:"offset"`

	// timezone
	Timezone string `json:"timezone,omitempty"`
}

// Validate validates this timezone
func (m *Timezone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayString(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffset(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Timezone) validateDisplayString(formats strfmt.Registry) error {

	if err := validate.Required("display_string", "body", m.DisplayString); err != nil {
		return err
	}

	return nil
}

func (m *Timezone) validateOffset(formats strfmt.Registry) error {

	if err := validate.Required("offset", "body", m.Offset); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Timezone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Timezone) UnmarshalBinary(b []byte) error {
	var res Timezone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}