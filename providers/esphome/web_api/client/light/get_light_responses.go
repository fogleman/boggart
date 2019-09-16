// Code generated by go-swagger; DO NOT EDIT.

package light

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/providers/esphome/web_api/models"
)

// GetLightReader is a Reader for the GetLight structure.
type GetLightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLightOK creates a GetLightOK with default headers values
func NewGetLightOK() *GetLightOK {
	return &GetLightOK{}
}

/*GetLightOK handles this case with default header values.

Successful operation
*/
type GetLightOK struct {
	Payload *models.Light
}

func (o *GetLightOK) Error() string {
	return fmt.Sprintf("[GET /light/{id}][%d] getLightOK  %+v", 200, o.Payload)
}

func (o *GetLightOK) GetPayload() *models.Light {
	return o.Payload
}

func (o *GetLightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Light)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}