// Code generated by go-swagger; DO NOT EDIT.

package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/providers/octoprint/models"
)

// GetChamberStateReader is a Reader for the GetChamberState structure.
type GetChamberStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChamberStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChamberStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewGetChamberStateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChamberStateOK creates a GetChamberStateOK with default headers values
func NewGetChamberStateOK() *GetChamberStateOK {
	return &GetChamberStateOK{}
}

/*GetChamberStateOK handles this case with default header values.

Successful operation
*/
type GetChamberStateOK struct {
	Payload *models.ChamberState
}

func (o *GetChamberStateOK) Error() string {
	return fmt.Sprintf("[GET /api/printer/chamber][%d] getChamberStateOK  %+v", 200, o.Payload)
}

func (o *GetChamberStateOK) GetPayload() *models.ChamberState {
	return o.Payload
}

func (o *GetChamberStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChamberState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChamberStateConflict creates a GetChamberStateConflict with default headers values
func NewGetChamberStateConflict() *GetChamberStateConflict {
	return &GetChamberStateConflict{}
}

/*GetChamberStateConflict handles this case with default header values.

If the printer is not operational
*/
type GetChamberStateConflict struct {
}

func (o *GetChamberStateConflict) Error() string {
	return fmt.Sprintf("[GET /api/printer/chamber][%d] getChamberStateConflict ", 409)
}

func (o *GetChamberStateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
