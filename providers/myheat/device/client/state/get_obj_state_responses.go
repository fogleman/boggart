// Code generated by go-swagger; DO NOT EDIT.

package state

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/myheat/device/models"
)

// GetObjStateReader is a Reader for the GetObjState structure.
type GetObjStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetObjStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetObjStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetObjStateOK creates a GetObjStateOK with default headers values
func NewGetObjStateOK() *GetObjStateOK {
	return &GetObjStateOK{}
}

/* GetObjStateOK describes a response with status code 200, with default header values.

Successful login
*/
type GetObjStateOK struct {
	Payload *models.StateObject
}

func (o *GetObjStateOK) Error() string {
	return fmt.Sprintf("[GET /getObjState][%d] getObjStateOK  %+v", 200, o.Payload)
}
func (o *GetObjStateOK) GetPayload() *models.StateObject {
	return o.Payload
}

func (o *GetObjStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StateObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}