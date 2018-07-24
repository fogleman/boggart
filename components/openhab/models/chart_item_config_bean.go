// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ChartItemConfigBean chart item config bean
// swagger:model ChartItemConfigBean
type ChartItemConfigBean struct {

	// axis
	Axis string `json:"axis,omitempty"`

	// chart
	Chart string `json:"chart,omitempty"`

	// fill
	Fill *bool `json:"fill,omitempty"`

	// item
	Item string `json:"item,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// legend
	Legend *bool `json:"legend,omitempty"`

	// line color
	LineColor string `json:"lineColor,omitempty"`

	// line style
	LineStyle string `json:"lineStyle,omitempty"`

	// line width
	LineWidth string `json:"lineWidth,omitempty"`

	// marker color
	MarkerColor string `json:"markerColor,omitempty"`

	// marker symbol
	MarkerSymbol string `json:"markerSymbol,omitempty"`

	// repeat time
	RepeatTime int32 `json:"repeatTime,omitempty"`
}

// Validate validates this chart item config bean
func (m *ChartItemConfigBean) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChartItemConfigBean) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChartItemConfigBean) UnmarshalBinary(b []byte) error {
	var res ChartItemConfigBean
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
