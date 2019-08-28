// Code generated by go-swagger; DO NOT EDIT.

package net

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/components/boggart/providers/hilink/models"
)

// GetCurrentPLMNReader is a Reader for the GetCurrentPLMN structure.
type GetCurrentPLMNReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentPLMNReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCurrentPLMNOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCurrentPLMNOK creates a GetCurrentPLMNOK with default headers values
func NewGetCurrentPLMNOK() *GetCurrentPLMNOK {
	return &GetCurrentPLMNOK{}
}

/*GetCurrentPLMNOK handles this case with default header values.

Successful operation
*/
type GetCurrentPLMNOK struct {
	Payload *models.CurrentPLMN
}

func (o *GetCurrentPLMNOK) Error() string {
	return fmt.Sprintf("[GET /net/current-plmn][%d] getCurrentPLMNOK  %+v", 200, o.Payload)
}

func (o *GetCurrentPLMNOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CurrentPLMN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}