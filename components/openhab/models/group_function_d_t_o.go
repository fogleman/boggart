// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GroupFunctionDTO group function d t o
// swagger:model GroupFunctionDTO
type GroupFunctionDTO struct {

	// name
	Name string `json:"name,omitempty"`

	// params
	Params []string `json:"params"`
}

// Validate validates this group function d t o
func (m *GroupFunctionDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupFunctionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupFunctionDTO) UnmarshalBinary(b []byte) error {
	var res GroupFunctionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
