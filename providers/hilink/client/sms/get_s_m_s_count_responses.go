// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/providers/hilink/models"
)

// GetSMSCountReader is a Reader for the GetSMSCount structure.
type GetSMSCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSMSCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSMSCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSMSCountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSMSCountOK creates a GetSMSCountOK with default headers values
func NewGetSMSCountOK() *GetSMSCountOK {
	return &GetSMSCountOK{}
}

/*GetSMSCountOK handles this case with default header values.

Successful operation
*/
type GetSMSCountOK struct {
	Payload *models.SMSCount
}

func (o *GetSMSCountOK) Error() string {
	return fmt.Sprintf("[GET /sms/sms-count][%d] getSMSCountOK  %+v", 200, o.Payload)
}

func (o *GetSMSCountOK) GetPayload() *models.SMSCount {
	return o.Payload
}

func (o *GetSMSCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SMSCount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSMSCountDefault creates a GetSMSCountDefault with default headers values
func NewGetSMSCountDefault(code int) *GetSMSCountDefault {
	return &GetSMSCountDefault{
		_statusCode: code,
	}
}

/*GetSMSCountDefault handles this case with default header values.

Unexpected error
*/
type GetSMSCountDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get s m s count default response
func (o *GetSMSCountDefault) Code() int {
	return o._statusCode
}

func (o *GetSMSCountDefault) Error() string {
	return fmt.Sprintf("[GET /sms/sms-count][%d] getSMSCount default  %+v", o._statusCode, o.Payload)
}

func (o *GetSMSCountDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSMSCountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}