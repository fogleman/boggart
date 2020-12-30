// Code generated by go-swagger; DO NOT EDIT.

package passes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/pass24online/models"
)

// GetPassesReader is a Reader for the GetPasses structure.
type GetPassesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPassesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPassesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPassesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPassesOK creates a GetPassesOK with default headers values
func NewGetPassesOK() *GetPassesOK {
	return &GetPassesOK{}
}

/*GetPassesOK handles this case with default header values.

Successful operation
*/
type GetPassesOK struct {
	Payload *models.PassesResponse
}

func (o *GetPassesOK) Error() string {
	return fmt.Sprintf("[GET /passes][%d] getPassesOK  %+v", 200, o.Payload)
}

func (o *GetPassesOK) GetPayload() *models.PassesResponse {
	return o.Payload
}

func (o *GetPassesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PassesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPassesDefault creates a GetPassesDefault with default headers values
func NewGetPassesDefault(code int) *GetPassesDefault {
	return &GetPassesDefault{
		_statusCode: code,
	}
}

/*GetPassesDefault handles this case with default header values.

Unexpected error
*/
type GetPassesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get passes default response
func (o *GetPassesDefault) Code() int {
	return o._statusCode
}

func (o *GetPassesDefault) Error() string {
	return fmt.Sprintf("[GET /passes][%d] getPasses default  %+v", o._statusCode, o.Payload)
}

func (o *GetPassesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPassesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}