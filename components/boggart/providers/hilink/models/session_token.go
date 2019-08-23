// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SessionToken session token
// swagger:model SessionToken
type SessionToken struct {

	// session
	Session string `json:"Session,omitempty" xml:"SesInfo"`

	// token
	Token string `json:"Token,omitempty" xml:"TokInfo"`
}

// Validate validates this session token
func (m *SessionToken) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SessionToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SessionToken) UnmarshalBinary(b []byte) error {
	var res SessionToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}