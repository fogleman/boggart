// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Forecast forecast
//
// swagger:model Forecast
type Forecast struct {

	// city
	City *ForecastCity `json:"city,omitempty"`

	// cnt
	Cnt uint64 `json:"cnt,omitempty"`

	// cod
	Cod string `json:"cod,omitempty"`

	// list
	List []*ForecastListItem `json:"list"`

	// message
	Message float64 `json:"message,omitempty"`
}

// Validate validates this forecast
func (m *Forecast) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Forecast) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	if m.City != nil {
		if err := m.City.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("city")
			}
			return err
		}
	}

	return nil
}

func (m *Forecast) validateList(formats strfmt.Registry) error {

	if swag.IsZero(m.List) { // not required
		return nil
	}

	for i := 0; i < len(m.List); i++ {
		if swag.IsZero(m.List[i]) { // not required
			continue
		}

		if m.List[i] != nil {
			if err := m.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Forecast) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Forecast) UnmarshalBinary(b []byte) error {
	var res Forecast
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ForecastCity forecast city
//
// swagger:model ForecastCity
type ForecastCity struct {

	// coord
	Coord *Coord `json:"coord,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// id
	ID uint64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// population
	Population uint64 `json:"population,omitempty"`

	// sunrise
	Sunrise uint64 `json:"sunrise,omitempty"`

	// sunset
	Sunset uint64 `json:"sunset,omitempty"`

	// timezone
	Timezone uint64 `json:"timezone,omitempty"`
}

// Validate validates this forecast city
func (m *ForecastCity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoord(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ForecastCity) validateCoord(formats strfmt.Registry) error {

	if swag.IsZero(m.Coord) { // not required
		return nil
	}

	if m.Coord != nil {
		if err := m.Coord.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("city" + "." + "coord")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ForecastCity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ForecastCity) UnmarshalBinary(b []byte) error {
	var res ForecastCity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}