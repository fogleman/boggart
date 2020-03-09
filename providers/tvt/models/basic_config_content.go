// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BasicConfigContent basic config content
// swagger:model BasicConfigContent
type BasicConfigContent struct {

	// product model
	ProductModel string `json:"productModel,omitempty" xml:"productModel"`

	// sn
	Sn string `json:"sn,omitempty" xml:"sn"`

	// software version
	SoftwareVersion string `json:"softwareVersion,omitempty" xml:"softwareVersion"`
}

// Validate validates this basic config content
func (m *BasicConfigContent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BasicConfigContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BasicConfigContent) UnmarshalBinary(b []byte) error {
	var res BasicConfigContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
