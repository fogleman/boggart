// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StateObject state object
//
// swagger:model StateObject
type StateObject struct {

	// alarms
	Alarms []*StateObjectAlarmsItems0 `json:"alarms"`

	// device flags
	DeviceFlags int64 `json:"deviceFlags,omitempty"`

	// device severity level
	DeviceSeverity int64 `json:"deviceSeverity,omitempty"`

	// heating mode
	HMode int64 `json:"hMode,omitempty"`

	// h modes
	HModes []*StateObjectHModesItems0 `json:"hModes"`

	// scheds
	Scheds []*StateObjectSchedsItems0 `json:"scheds"`

	// Security mode enabled or disabled. Activate in control panel
	SecurityArmed *bool `json:"securityArmed,omitempty"`

	// sim balance
	SimBalance float64 `json:"simBalance,omitempty"`

	// sim signal
	SimSignal int64 `json:"simSignal,omitempty"`
}

// Validate validates this state object
func (m *StateObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlarms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHModes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateObject) validateAlarms(formats strfmt.Registry) error {
	if swag.IsZero(m.Alarms) { // not required
		return nil
	}

	for i := 0; i < len(m.Alarms); i++ {
		if swag.IsZero(m.Alarms[i]) { // not required
			continue
		}

		if m.Alarms[i] != nil {
			if err := m.Alarms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alarms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alarms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StateObject) validateHModes(formats strfmt.Registry) error {
	if swag.IsZero(m.HModes) { // not required
		return nil
	}

	for i := 0; i < len(m.HModes); i++ {
		if swag.IsZero(m.HModes[i]) { // not required
			continue
		}

		if m.HModes[i] != nil {
			if err := m.HModes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hModes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hModes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StateObject) validateScheds(formats strfmt.Registry) error {
	if swag.IsZero(m.Scheds) { // not required
		return nil
	}

	for i := 0; i < len(m.Scheds); i++ {
		if swag.IsZero(m.Scheds[i]) { // not required
			continue
		}

		if m.Scheds[i] != nil {
			if err := m.Scheds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scheds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scheds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this state object based on the context it is used
func (m *StateObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlarms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHModes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScheds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateObject) contextValidateAlarms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Alarms); i++ {

		if m.Alarms[i] != nil {
			if err := m.Alarms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alarms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alarms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StateObject) contextValidateHModes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HModes); i++ {

		if m.HModes[i] != nil {
			if err := m.HModes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hModes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hModes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StateObject) contextValidateScheds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Scheds); i++ {

		if m.Scheds[i] != nil {
			if err := m.Scheds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scheds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scheds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StateObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateObject) UnmarshalBinary(b []byte) error {
	var res StateObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StateObjectAlarmsItems0 state object alarms items0
//
// swagger:model StateObjectAlarmsItems0
type StateObjectAlarmsItems0 struct {

	// f
	F int64 `json:"f,omitempty"`

	// id
	I int64 `json:"i,omitempty"`

	// name
	N string `json:"n,omitempty"`

	// device severity level
	// Enum: [1 32 64]
	Sev int64 `json:"sev,omitempty"`

	// object type
	T int64 `json:"t,omitempty"`
}

// Validate validates this state object alarms items0
func (m *StateObjectAlarmsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSev(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var stateObjectAlarmsItems0TypeSevPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,32,64]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stateObjectAlarmsItems0TypeSevPropEnum = append(stateObjectAlarmsItems0TypeSevPropEnum, v)
	}
}

// prop value enum
func (m *StateObjectAlarmsItems0) validateSevEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, stateObjectAlarmsItems0TypeSevPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StateObjectAlarmsItems0) validateSev(formats strfmt.Registry) error {
	if swag.IsZero(m.Sev) { // not required
		return nil
	}

	// value enum
	if err := m.validateSevEnum("sev", "body", m.Sev); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this state object alarms items0 based on context it is used
func (m *StateObjectAlarmsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StateObjectAlarmsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateObjectAlarmsItems0) UnmarshalBinary(b []byte) error {
	var res StateObjectAlarmsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StateObjectHModesItems0 state object h modes items0
//
// swagger:model StateObjectHModesItems0
type StateObjectHModesItems0 struct {

	// id
	I int64 `json:"i,omitempty"`

	// name
	N string `json:"n,omitempty"`
}

// Validate validates this state object h modes items0
func (m *StateObjectHModesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this state object h modes items0 based on context it is used
func (m *StateObjectHModesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StateObjectHModesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateObjectHModesItems0) UnmarshalBinary(b []byte) error {
	var res StateObjectHModesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StateObjectSchedsItems0 state object scheds items0
//
// swagger:model StateObjectSchedsItems0
type StateObjectSchedsItems0 struct {

	// id
	I int64 `json:"i,omitempty"`

	// name
	N string `json:"n,omitempty"`
}

// Validate validates this state object scheds items0
func (m *StateObjectSchedsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this state object scheds items0 based on context it is used
func (m *StateObjectSchedsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StateObjectSchedsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateObjectSchedsItems0) UnmarshalBinary(b []byte) error {
	var res StateObjectSchedsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
