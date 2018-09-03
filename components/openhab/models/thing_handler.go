// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ThingHandler thing handler
// swagger:model ThingHandler
type ThingHandler struct {

	// thing
	Thing *Thing `json:"thing,omitempty"`
}

// Validate validates this thing handler
func (m *ThingHandler) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateThing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThingHandler) validateThing(formats strfmt.Registry) error {

	if swag.IsZero(m.Thing) { // not required
		return nil
	}

	if m.Thing != nil {
		if err := m.Thing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThingHandler) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThingHandler) UnmarshalBinary(b []byte) error {
	var res ThingHandler
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}