// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/kihamo/boggart/providers/hilink/models"
)

// DeviceControlReader is a Reader for the DeviceControl structure.
type DeviceControlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeviceControlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeviceControlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeviceControlDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeviceControlOK creates a DeviceControlOK with default headers values
func NewDeviceControlOK() *DeviceControlOK {
	return &DeviceControlOK{}
}

/* DeviceControlOK describes a response with status code 200, with default header values.

Successful operation
*/
type DeviceControlOK struct {
	Payload string
}

func (o *DeviceControlOK) Error() string {
	return fmt.Sprintf("[POST /api/device/control][%d] deviceControlOK  %+v", 200, o.Payload)
}
func (o *DeviceControlOK) GetPayload() string {
	return o.Payload
}

func (o *DeviceControlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeviceControlDefault creates a DeviceControlDefault with default headers values
func NewDeviceControlDefault(code int) *DeviceControlDefault {
	return &DeviceControlDefault{
		_statusCode: code,
	}
}

/* DeviceControlDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type DeviceControlDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the device control default response
func (o *DeviceControlDefault) Code() int {
	return o._statusCode
}

func (o *DeviceControlDefault) Error() string {
	return fmt.Sprintf("[POST /api/device/control][%d] DeviceControl default  %+v", o._statusCode, o.Payload)
}
func (o *DeviceControlDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeviceControlDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeviceControlBody device control body
swagger:model DeviceControlBody
*/
type DeviceControlBody struct {

	// control
	// Enum: [1 2 3 4]
	Control int64 `json:"Control,omitempty" xml:"Control,omitempty"`
}

// Validate validates this device control body
func (o *DeviceControlBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateControl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deviceControlBodyTypeControlPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3,4]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceControlBodyTypeControlPropEnum = append(deviceControlBodyTypeControlPropEnum, v)
	}
}

// prop value enum
func (o *DeviceControlBody) validateControlEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, deviceControlBodyTypeControlPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *DeviceControlBody) validateControl(formats strfmt.Registry) error {
	if swag.IsZero(o.Control) { // not required
		return nil
	}

	// value enum
	if err := o.validateControlEnum("request"+"."+"Control", "body", o.Control); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device control body based on context it is used
func (o *DeviceControlBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeviceControlBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeviceControlBody) UnmarshalBinary(b []byte) error {
	var res DeviceControlBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
