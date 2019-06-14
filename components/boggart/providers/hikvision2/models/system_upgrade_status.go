// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SystemUpgradeStatus system upgrade status
// swagger:model SystemUpgradeStatus
type SystemUpgradeStatus struct {

	// percent
	Percent int64 `json:"percent,omitempty" xml:"percent"`

	// upgrading
	Upgrading bool `json:"upgrading,omitempty" xml:"upgrading"`
}

// Validate validates this system upgrade status
func (m *SystemUpgradeStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemUpgradeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemUpgradeStatus) UnmarshalBinary(b []byte) error {
	var res SystemUpgradeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
