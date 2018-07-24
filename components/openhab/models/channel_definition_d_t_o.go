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

// ChannelDefinitionDTO channel definition d t o
// swagger:model ChannelDefinitionDTO
type ChannelDefinitionDTO struct {

	// advanced
	Advanced *bool `json:"advanced,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// properties
	Properties map[string]string `json:"properties,omitempty"`

	// state description
	StateDescription *StateDescription `json:"stateDescription,omitempty"`

	// tags
	// Unique: true
	Tags []string `json:"tags"`

	// type UID
	TypeUID string `json:"typeUID,omitempty"`
}

// Validate validates this channel definition d t o
func (m *ChannelDefinitionDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelDefinitionDTO) validateStateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.StateDescription) { // not required
		return nil
	}

	if m.StateDescription != nil {
		if err := m.StateDescription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stateDescription")
			}
			return err
		}
	}

	return nil
}

func (m *ChannelDefinitionDTO) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelDefinitionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelDefinitionDTO) UnmarshalBinary(b []byte) error {
	var res ChannelDefinitionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
