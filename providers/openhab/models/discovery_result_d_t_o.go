// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DiscoveryResultDTO discovery result d t o
// swagger:model DiscoveryResultDTO
type DiscoveryResultDTO struct {

	// bridge UID
	BridgeUID string `json:"bridgeUID,omitempty"`

	// flag
	// Enum: [NEW IGNORED]
	Flag string `json:"flag,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// properties
	Properties map[string]interface{} `json:"properties,omitempty"`

	// representation property
	RepresentationProperty string `json:"representationProperty,omitempty"`

	// thing type UID
	ThingTypeUID string `json:"thingTypeUID,omitempty"`

	// thing UID
	ThingUID string `json:"thingUID,omitempty"`
}

// Validate validates this discovery result d t o
func (m *DiscoveryResultDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var discoveryResultDTOTypeFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NEW","IGNORED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		discoveryResultDTOTypeFlagPropEnum = append(discoveryResultDTOTypeFlagPropEnum, v)
	}
}

const (

	// DiscoveryResultDTOFlagNEW captures enum value "NEW"
	DiscoveryResultDTOFlagNEW string = "NEW"

	// DiscoveryResultDTOFlagIGNORED captures enum value "IGNORED"
	DiscoveryResultDTOFlagIGNORED string = "IGNORED"
)

// prop value enum
func (m *DiscoveryResultDTO) validateFlagEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, discoveryResultDTOTypeFlagPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DiscoveryResultDTO) validateFlag(formats strfmt.Registry) error {

	if swag.IsZero(m.Flag) { // not required
		return nil
	}

	// value enum
	if err := m.validateFlagEnum("flag", "body", m.Flag); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiscoveryResultDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscoveryResultDTO) UnmarshalBinary(b []byte) error {
	var res DiscoveryResultDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}