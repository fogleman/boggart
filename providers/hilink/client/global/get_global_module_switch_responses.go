// Code generated by go-swagger; DO NOT EDIT.

package global

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/hilink/models"
)

// GetGlobalModuleSwitchReader is a Reader for the GetGlobalModuleSwitch structure.
type GetGlobalModuleSwitchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGlobalModuleSwitchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGlobalModuleSwitchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetGlobalModuleSwitchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGlobalModuleSwitchOK creates a GetGlobalModuleSwitchOK with default headers values
func NewGetGlobalModuleSwitchOK() *GetGlobalModuleSwitchOK {
	return &GetGlobalModuleSwitchOK{}
}

/* GetGlobalModuleSwitchOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetGlobalModuleSwitchOK struct {
	Payload *models.ModuleSwitch
}

func (o *GetGlobalModuleSwitchOK) Error() string {
	return fmt.Sprintf("[GET /api/global/module-switch][%d] getGlobalModuleSwitchOK  %+v", 200, o.Payload)
}
func (o *GetGlobalModuleSwitchOK) GetPayload() *models.ModuleSwitch {
	return o.Payload
}

func (o *GetGlobalModuleSwitchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleSwitch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalModuleSwitchDefault creates a GetGlobalModuleSwitchDefault with default headers values
func NewGetGlobalModuleSwitchDefault(code int) *GetGlobalModuleSwitchDefault {
	return &GetGlobalModuleSwitchDefault{
		_statusCode: code,
	}
}

/* GetGlobalModuleSwitchDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type GetGlobalModuleSwitchDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get global module switch default response
func (o *GetGlobalModuleSwitchDefault) Code() int {
	return o._statusCode
}

func (o *GetGlobalModuleSwitchDefault) Error() string {
	return fmt.Sprintf("[GET /api/global/module-switch][%d] getGlobalModuleSwitch default  %+v", o._statusCode, o.Payload)
}
func (o *GetGlobalModuleSwitchDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGlobalModuleSwitchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
