// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ParameterOptionDTO parameter option d t o
// swagger:model ParameterOptionDTO
type ParameterOptionDTO struct {

	// label
	Label string `json:"label,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this parameter option d t o
func (m *ParameterOptionDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParameterOptionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterOptionDTO) UnmarshalBinary(b []byte) error {
	var res ParameterOptionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
