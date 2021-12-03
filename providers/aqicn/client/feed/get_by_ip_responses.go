// Code generated by go-swagger; DO NOT EDIT.

package feed

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/providers/aqicn/models"
)

// GetByIPReader is a Reader for the GetByIP structure.
type GetByIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetByIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetByIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetByIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetByIPOK creates a GetByIPOK with default headers values
func NewGetByIPOK() *GetByIPOK {
	return &GetByIPOK{}
}

/*GetByIPOK handles this case with default header values.

Successful operation
*/
type GetByIPOK struct {
	Payload *models.Feed
}

func (o *GetByIPOK) Error() string {
	return fmt.Sprintf("[GET /feed/here/][%d] getByIpOK  %+v", 200, o.Payload)
}

func (o *GetByIPOK) GetPayload() *models.Feed {
	return o.Payload
}

func (o *GetByIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Feed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByIPDefault creates a GetByIPDefault with default headers values
func NewGetByIPDefault(code int) *GetByIPDefault {
	return &GetByIPDefault{
		_statusCode: code,
	}
}

/*GetByIPDefault handles this case with default header values.

Unexpected error
*/
type GetByIPDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get by IP default response
func (o *GetByIPDefault) Code() int {
	return o._statusCode
}

func (o *GetByIPDefault) Error() string {
	return fmt.Sprintf("[GET /feed/here/][%d] getByIP default  %+v", o._statusCode, o.Payload)
}

func (o *GetByIPDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetByIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}