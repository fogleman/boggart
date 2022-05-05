// Code generated by go-swagger; DO NOT EDIT.

package ptz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/hikvision/models"
)

// SetPtzPositionRelativeReader is a Reader for the SetPtzPositionRelative structure.
type SetPtzPositionRelativeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetPtzPositionRelativeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetPtzPositionRelativeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetPtzPositionRelativeOK creates a SetPtzPositionRelativeOK with default headers values
func NewSetPtzPositionRelativeOK() *SetPtzPositionRelativeOK {
	return &SetPtzPositionRelativeOK{}
}

/* SetPtzPositionRelativeOK describes a response with status code 200, with default header values.

Successful operation
*/
type SetPtzPositionRelativeOK struct {
	Payload *models.Status
}

func (o *SetPtzPositionRelativeOK) Error() string {
	return fmt.Sprintf("[PUT /PTZCtrl/channels/{channel}/relative][%d] setPtzPositionRelativeOK  %+v", 200, o.Payload)
}
func (o *SetPtzPositionRelativeOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *SetPtzPositionRelativeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
