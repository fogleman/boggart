// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/components/boggart/providers/hilink/models"
)

// GetDeviceSignalReader is a Reader for the GetDeviceSignal structure.
type GetDeviceSignalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceSignalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceSignalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDeviceSignalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceSignalOK creates a GetDeviceSignalOK with default headers values
func NewGetDeviceSignalOK() *GetDeviceSignalOK {
	return &GetDeviceSignalOK{}
}

/*GetDeviceSignalOK handles this case with default header values.

Successful operation
*/
type GetDeviceSignalOK struct {
	Payload *models.DeviceSignal
}

func (o *GetDeviceSignalOK) Error() string {
	return fmt.Sprintf("[GET /device/signal][%d] getDeviceSignalOK  %+v", 200, o.Payload)
}

func (o *GetDeviceSignalOK) GetPayload() *models.DeviceSignal {
	return o.Payload
}

func (o *GetDeviceSignalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceSignal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceSignalDefault creates a GetDeviceSignalDefault with default headers values
func NewGetDeviceSignalDefault(code int) *GetDeviceSignalDefault {
	return &GetDeviceSignalDefault{
		_statusCode: code,
	}
}

/*GetDeviceSignalDefault handles this case with default header values.

Unexpected error
*/
type GetDeviceSignalDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get device signal default response
func (o *GetDeviceSignalDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceSignalDefault) Error() string {
	return fmt.Sprintf("[GET /device/signal][%d] getDeviceSignal default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceSignalDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeviceSignalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
