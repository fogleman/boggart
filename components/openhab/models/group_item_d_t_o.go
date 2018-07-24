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

// GroupItemDTO group item d t o
// swagger:model GroupItemDTO
type GroupItemDTO struct {

	// category
	Category string `json:"category,omitempty"`

	// function
	Function *GroupFunctionDTO `json:"function,omitempty"`

	// group names
	GroupNames []string `json:"groupNames"`

	// group type
	GroupType string `json:"groupType,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// tags
	// Unique: true
	Tags []string `json:"tags"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this group item d t o
func (m *GroupItemDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunction(formats); err != nil {
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

func (m *GroupItemDTO) validateFunction(formats strfmt.Registry) error {

	if swag.IsZero(m.Function) { // not required
		return nil
	}

	if m.Function != nil {
		if err := m.Function.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("function")
			}
			return err
		}
	}

	return nil
}

func (m *GroupItemDTO) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupItemDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupItemDTO) UnmarshalBinary(b []byte) error {
	var res GroupItemDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
