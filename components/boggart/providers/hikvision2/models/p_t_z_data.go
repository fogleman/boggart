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

// PTZData p t z data
// swagger:model PTZData
type PTZData struct {

	// absolute high
	AbsoluteHigh *PtzAbsoluteHigh `json:"absoluteHigh,omitempty" xml:"AbsoluteHigh"`

	// duration
	// Format: duration
	Duration strfmt.Duration `json:"duration,omitempty" xml:"Momentary>duration"`

	// pan
	// Maximum: 100
	// Minimum: -100
	Pan *int64 `json:"pan,omitempty" xml:"pan"`

	// relative
	Relative *PtzRelative `json:"relative,omitempty" xml:"Relative"`

	// tilt
	// Maximum: 100
	// Minimum: -100
	Tilt *int64 `json:"tilt,omitempty" xml:"tilt"`

	// zoom
	// Maximum: 100
	// Minimum: -100
	Zoom *int64 `json:"zoom,omitempty" xml:"zoom"`
}

// Validate validates this p t z data
func (m *PTZData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbsoluteHigh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTilt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTZData) validateAbsoluteHigh(formats strfmt.Registry) error {

	if swag.IsZero(m.AbsoluteHigh) { // not required
		return nil
	}

	if m.AbsoluteHigh != nil {
		if err := m.AbsoluteHigh.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("absoluteHigh")
			}
			return err
		}
	}

	return nil
}

func (m *PTZData) validateDuration(formats strfmt.Registry) error {

	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	if err := validate.FormatOf("duration", "body", "duration", m.Duration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTZData) validatePan(formats strfmt.Registry) error {

	if swag.IsZero(m.Pan) { // not required
		return nil
	}

	if err := validate.MinimumInt("pan", "body", int64(*m.Pan), -100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pan", "body", int64(*m.Pan), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *PTZData) validateRelative(formats strfmt.Registry) error {

	if swag.IsZero(m.Relative) { // not required
		return nil
	}

	if m.Relative != nil {
		if err := m.Relative.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relative")
			}
			return err
		}
	}

	return nil
}

func (m *PTZData) validateTilt(formats strfmt.Registry) error {

	if swag.IsZero(m.Tilt) { // not required
		return nil
	}

	if err := validate.MinimumInt("tilt", "body", int64(*m.Tilt), -100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("tilt", "body", int64(*m.Tilt), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *PTZData) validateZoom(formats strfmt.Registry) error {

	if swag.IsZero(m.Zoom) { // not required
		return nil
	}

	if err := validate.MinimumInt("zoom", "body", int64(*m.Zoom), -100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("zoom", "body", int64(*m.Zoom), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTZData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTZData) UnmarshalBinary(b []byte) error {
	var res PTZData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
