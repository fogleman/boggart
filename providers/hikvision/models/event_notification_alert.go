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

// EventNotificationAlert event notification alert
//
// swagger:model EventNotificationAlert
type EventNotificationAlert struct {

	// active post count
	ActivePostCount uint64 `json:"activePostCount,omitempty" xml:"activePostCount,omitempty"`

	// channel ID
	ChannelID uint64 `json:"channelID,omitempty" xml:"channelID,omitempty"`

	// date time
	// Format: date-time
	DateTime strfmt.DateTime `json:"dateTime,omitempty" xml:"dateTime,omitempty"`

	// dyn channel ID
	DynChannelID uint64 `json:"dynChannelID,omitempty" xml:"dynChannelID,omitempty"`

	// dyn input i o port ID
	DynInputIOPortID string `json:"dynInputIOPortID,omitempty" xml:"dynInputIOPortID,omitempty"`

	// event description
	EventDescription string `json:"eventDescription,omitempty" xml:"eventDescription,omitempty"`

	// event push
	EventPush string `json:"eventPush,omitempty" xml:"Extensions>eventPush,omitempty"`

	// event state
	// Enum: [active inactive]
	EventState string `json:"eventState,omitempty" xml:"eventState,omitempty"`

	// event type
	// Enum: [IO VMD videoloss shelteralarm facedetection defocus audioexception scenechangedetection fielddetection linedetection regionEntrance regionExiting loitering group rapidMove parking unattendedBaggage attendedBaggage PIR peopleDetection]
	EventType string `json:"eventType,omitempty" xml:"eventType,omitempty"`

	// input i o port ID
	InputIOPortID uint64 `json:"inputIOPortID,omitempty" xml:"inputIOPortID,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`

	// ipv6 address
	IPV6Address string `json:"ipv6Address,omitempty" xml:"ipv6Address,omitempty"`

	// mac address
	MacAddress string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`

	// port
	Port uint64 `json:"port,omitempty" xml:"portNo,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty" xml:"protocol,omitempty"`

	// region ID
	RegionID string `json:"regionID,omitempty" xml:"DetectionRegionList>DetectionRegionEntry>regionID,omitempty"`

	// sensitivity level
	SensitivityLevel uint64 `json:"sensitivityLevel,omitempty" xml:"DetectionRegionList>DetectionRegionEntry>sensitivityLevel,omitempty"`

	// serial number
	SerialNumber string `json:"serialNumber,omitempty" xml:"Extensions>serialNumber,omitempty"`
}

// Validate validates this event notification alert
func (m *EventNotificationAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventNotificationAlert) validateDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("dateTime", "body", "date-time", m.DateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var eventNotificationAlertTypeEventStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventNotificationAlertTypeEventStatePropEnum = append(eventNotificationAlertTypeEventStatePropEnum, v)
	}
}

const (

	// EventNotificationAlertEventStateActive captures enum value "active"
	EventNotificationAlertEventStateActive string = "active"

	// EventNotificationAlertEventStateInactive captures enum value "inactive"
	EventNotificationAlertEventStateInactive string = "inactive"
)

// prop value enum
func (m *EventNotificationAlert) validateEventStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, eventNotificationAlertTypeEventStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EventNotificationAlert) validateEventState(formats strfmt.Registry) error {
	if swag.IsZero(m.EventState) { // not required
		return nil
	}

	// value enum
	if err := m.validateEventStateEnum("eventState", "body", m.EventState); err != nil {
		return err
	}

	return nil
}

var eventNotificationAlertTypeEventTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IO","VMD","videoloss","shelteralarm","facedetection","defocus","audioexception","scenechangedetection","fielddetection","linedetection","regionEntrance","regionExiting","loitering","group","rapidMove","parking","unattendedBaggage","attendedBaggage","PIR","peopleDetection"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventNotificationAlertTypeEventTypePropEnum = append(eventNotificationAlertTypeEventTypePropEnum, v)
	}
}

const (

	// EventNotificationAlertEventTypeIO captures enum value "IO"
	EventNotificationAlertEventTypeIO string = "IO"

	// EventNotificationAlertEventTypeVMD captures enum value "VMD"
	EventNotificationAlertEventTypeVMD string = "VMD"

	// EventNotificationAlertEventTypeVideoloss captures enum value "videoloss"
	EventNotificationAlertEventTypeVideoloss string = "videoloss"

	// EventNotificationAlertEventTypeShelteralarm captures enum value "shelteralarm"
	EventNotificationAlertEventTypeShelteralarm string = "shelteralarm"

	// EventNotificationAlertEventTypeFacedetection captures enum value "facedetection"
	EventNotificationAlertEventTypeFacedetection string = "facedetection"

	// EventNotificationAlertEventTypeDefocus captures enum value "defocus"
	EventNotificationAlertEventTypeDefocus string = "defocus"

	// EventNotificationAlertEventTypeAudioexception captures enum value "audioexception"
	EventNotificationAlertEventTypeAudioexception string = "audioexception"

	// EventNotificationAlertEventTypeScenechangedetection captures enum value "scenechangedetection"
	EventNotificationAlertEventTypeScenechangedetection string = "scenechangedetection"

	// EventNotificationAlertEventTypeFielddetection captures enum value "fielddetection"
	EventNotificationAlertEventTypeFielddetection string = "fielddetection"

	// EventNotificationAlertEventTypeLinedetection captures enum value "linedetection"
	EventNotificationAlertEventTypeLinedetection string = "linedetection"

	// EventNotificationAlertEventTypeRegionEntrance captures enum value "regionEntrance"
	EventNotificationAlertEventTypeRegionEntrance string = "regionEntrance"

	// EventNotificationAlertEventTypeRegionExiting captures enum value "regionExiting"
	EventNotificationAlertEventTypeRegionExiting string = "regionExiting"

	// EventNotificationAlertEventTypeLoitering captures enum value "loitering"
	EventNotificationAlertEventTypeLoitering string = "loitering"

	// EventNotificationAlertEventTypeGroup captures enum value "group"
	EventNotificationAlertEventTypeGroup string = "group"

	// EventNotificationAlertEventTypeRapidMove captures enum value "rapidMove"
	EventNotificationAlertEventTypeRapidMove string = "rapidMove"

	// EventNotificationAlertEventTypeParking captures enum value "parking"
	EventNotificationAlertEventTypeParking string = "parking"

	// EventNotificationAlertEventTypeUnattendedBaggage captures enum value "unattendedBaggage"
	EventNotificationAlertEventTypeUnattendedBaggage string = "unattendedBaggage"

	// EventNotificationAlertEventTypeAttendedBaggage captures enum value "attendedBaggage"
	EventNotificationAlertEventTypeAttendedBaggage string = "attendedBaggage"

	// EventNotificationAlertEventTypePIR captures enum value "PIR"
	EventNotificationAlertEventTypePIR string = "PIR"

	// EventNotificationAlertEventTypePeopleDetection captures enum value "peopleDetection"
	EventNotificationAlertEventTypePeopleDetection string = "peopleDetection"
)

// prop value enum
func (m *EventNotificationAlert) validateEventTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, eventNotificationAlertTypeEventTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EventNotificationAlert) validateEventType(formats strfmt.Registry) error {
	if swag.IsZero(m.EventType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEventTypeEnum("eventType", "body", m.EventType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event notification alert based on context it is used
func (m *EventNotificationAlert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventNotificationAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventNotificationAlert) UnmarshalBinary(b []byte) error {
	var res EventNotificationAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
