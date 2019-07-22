// Code generated by go-swagger; DO NOT EDIT.

package ptz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/components/boggart/providers/hikvision/models"
)

// GetPtzStatusReader is a Reader for the GetPtzStatus structure.
type GetPtzStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPtzStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPtzStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPtzStatusOK creates a GetPtzStatusOK with default headers values
func NewGetPtzStatusOK() *GetPtzStatusOK {
	return &GetPtzStatusOK{}
}

/*GetPtzStatusOK handles this case with default header values.

Successful operation
*/
type GetPtzStatusOK struct {
	Payload *models.PTZStatus
}

func (o *GetPtzStatusOK) Error() string {
	return fmt.Sprintf("[GET /PTZCtrl/channels/{channel}/status][%d] getPtzStatusOK  %+v", 200, o.Payload)
}

func (o *GetPtzStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTZStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}