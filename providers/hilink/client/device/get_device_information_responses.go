// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/hilink/models"
)

// GetDeviceInformationReader is a Reader for the GetDeviceInformation structure.
type GetDeviceInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDeviceInformationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceInformationOK creates a GetDeviceInformationOK with default headers values
func NewGetDeviceInformationOK() *GetDeviceInformationOK {
	return &GetDeviceInformationOK{}
}

/* GetDeviceInformationOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceInformationOK struct {
	Payload *models.DeviceInformation
}

func (o *GetDeviceInformationOK) Error() string {
	return fmt.Sprintf("[GET /api/device/information][%d] getDeviceInformationOK  %+v", 200, o.Payload)
}
func (o *GetDeviceInformationOK) GetPayload() *models.DeviceInformation {
	return o.Payload
}

func (o *GetDeviceInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceInformation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceInformationDefault creates a GetDeviceInformationDefault with default headers values
func NewGetDeviceInformationDefault(code int) *GetDeviceInformationDefault {
	return &GetDeviceInformationDefault{
		_statusCode: code,
	}
}

/* GetDeviceInformationDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type GetDeviceInformationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get device information default response
func (o *GetDeviceInformationDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceInformationDefault) Error() string {
	return fmt.Sprintf("[GET /api/device/information][%d] getDeviceInformation default  %+v", o._statusCode, o.Payload)
}
func (o *GetDeviceInformationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeviceInformationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
