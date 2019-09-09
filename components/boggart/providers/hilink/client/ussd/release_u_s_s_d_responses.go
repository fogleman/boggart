// Code generated by go-swagger; DO NOT EDIT.

package ussd

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/components/boggart/providers/hilink/models"
)

// ReleaseUSSDReader is a Reader for the ReleaseUSSD structure.
type ReleaseUSSDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReleaseUSSDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReleaseUSSDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReleaseUSSDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReleaseUSSDOK creates a ReleaseUSSDOK with default headers values
func NewReleaseUSSDOK() *ReleaseUSSDOK {
	return &ReleaseUSSDOK{}
}

/*ReleaseUSSDOK handles this case with default header values.

Successful operation
*/
type ReleaseUSSDOK struct {
	Payload string
}

func (o *ReleaseUSSDOK) Error() string {
	return fmt.Sprintf("[GET /ussd/release][%d] releaseUSSDOK  %+v", 200, o.Payload)
}

func (o *ReleaseUSSDOK) GetPayload() string {
	return o.Payload
}

func (o *ReleaseUSSDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReleaseUSSDDefault creates a ReleaseUSSDDefault with default headers values
func NewReleaseUSSDDefault(code int) *ReleaseUSSDDefault {
	return &ReleaseUSSDDefault{
		_statusCode: code,
	}
}

/*ReleaseUSSDDefault handles this case with default header values.

Unexpected error
*/
type ReleaseUSSDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the release u s s d default response
func (o *ReleaseUSSDDefault) Code() int {
	return o._statusCode
}

func (o *ReleaseUSSDDefault) Error() string {
	return fmt.Sprintf("[GET /ussd/release][%d] releaseUSSD default  %+v", o._statusCode, o.Payload)
}

func (o *ReleaseUSSDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReleaseUSSDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
