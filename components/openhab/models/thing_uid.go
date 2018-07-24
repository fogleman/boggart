// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ThingUID thing UID
// swagger:model ThingUID
type ThingUID struct {

	// as string
	AsString string `json:"asString,omitempty"`

	// binding Id
	BindingID string `json:"bindingId,omitempty"`

	// bridge ids
	BridgeIds []string `json:"bridgeIds"`

	// id
	ID string `json:"id,omitempty"`

	// thing type Id
	ThingTypeID string `json:"thingTypeId,omitempty"`

	// thing type UID
	ThingTypeUID *ThingTypeUID `json:"thingTypeUID,omitempty"`
}

// Validate validates this thing UID
func (m *ThingUID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateThingTypeUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThingUID) validateThingTypeUID(formats strfmt.Registry) error {

	if swag.IsZero(m.ThingTypeUID) { // not required
		return nil
	}

	if m.ThingTypeUID != nil {
		if err := m.ThingTypeUID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thingTypeUID")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThingUID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThingUID) UnmarshalBinary(b []byte) error {
	var res ThingUID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
