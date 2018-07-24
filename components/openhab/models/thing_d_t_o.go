// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ThingDTO thing d t o
// swagger:model ThingDTO
type ThingDTO struct {

	// UID
	UID string `json:"UID,omitempty"`

	// bridge UID
	BridgeUID string `json:"bridgeUID,omitempty"`

	// channels
	Channels []*ChannelDTO `json:"channels"`

	// configuration
	Configuration map[string]interface{} `json:"configuration,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// properties
	Properties map[string]string `json:"properties,omitempty"`

	// thing type UID
	ThingTypeUID string `json:"thingTypeUID,omitempty"`
}

// Validate validates this thing d t o
func (m *ThingDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThingDTO) validateChannels(formats strfmt.Registry) error {

	if swag.IsZero(m.Channels) { // not required
		return nil
	}

	for i := 0; i < len(m.Channels); i++ {
		if swag.IsZero(m.Channels[i]) { // not required
			continue
		}

		if m.Channels[i] != nil {
			if err := m.Channels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("channels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThingDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThingDTO) UnmarshalBinary(b []byte) error {
	var res ThingDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
