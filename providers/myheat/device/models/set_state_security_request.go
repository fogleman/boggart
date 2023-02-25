// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SetStateSecurityRequest set state security request
//
// swagger:model SetStateSecurityRequest
type SetStateSecurityRequest struct {

	// action
	// Enum: [armSecurity disarmSecurity]
	Action string `json:"action,omitempty"`
}

// Validate validates this set state security request
func (m *SetStateSecurityRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var setStateSecurityRequestTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["armSecurity","disarmSecurity"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		setStateSecurityRequestTypeActionPropEnum = append(setStateSecurityRequestTypeActionPropEnum, v)
	}
}

const (

	// SetStateSecurityRequestActionArmSecurity captures enum value "armSecurity"
	SetStateSecurityRequestActionArmSecurity string = "armSecurity"

	// SetStateSecurityRequestActionDisarmSecurity captures enum value "disarmSecurity"
	SetStateSecurityRequestActionDisarmSecurity string = "disarmSecurity"
)

// prop value enum
func (m *SetStateSecurityRequest) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, setStateSecurityRequestTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SetStateSecurityRequest) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this set state security request based on context it is used
func (m *SetStateSecurityRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetStateSecurityRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetStateSecurityRequest) UnmarshalBinary(b []byte) error {
	var res SetStateSecurityRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
