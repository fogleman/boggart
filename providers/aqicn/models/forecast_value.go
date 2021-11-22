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

// ForecastValue forecast value
// swagger:model ForecastValue
type ForecastValue struct {

	// Avg value
	Avg float64 `json:"avg,omitempty"`

	// Day date
	// Format: date
	Day strfmt.Date `json:"day,omitempty"`

	// Max value
	Max float64 `json:"max,omitempty"`

	// Min value
	Min float64 `json:"min,omitempty"`
}

// Validate validates this forecast value
func (m *ForecastValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ForecastValue) validateDay(formats strfmt.Registry) error {

	if swag.IsZero(m.Day) { // not required
		return nil
	}

	if err := validate.FormatOf("day", "body", "date", m.Day.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ForecastValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ForecastValue) UnmarshalBinary(b []byte) error {
	var res ForecastValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
