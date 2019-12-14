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

// GetToolStateReader is a Reader for the GetToolState structure.
type GetToolStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetToolStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetToolStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewGetToolStateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetToolStateOK creates a GetToolStateOK with default headers values
func NewGetToolStateOK() *GetToolStateOK {
	return &GetToolStateOK{}
}

/*GetToolStateOK handles this case with default header values.

Successful operation
*/
type GetToolStateOK struct {
	Payload *models.ToolState
}

func (o *GetToolStateOK) Error() string {
	return fmt.Sprintf("[GET /api/printer/tool][%d] getToolStateOK  %+v", 200, o.Payload)
}

func (o *GetToolStateOK) GetPayload() *models.ToolState {
	return o.Payload
}

func (o *GetToolStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ToolState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetToolStateConflict creates a GetToolStateConflict with default headers values
func NewGetToolStateConflict() *GetToolStateConflict {
	return &GetToolStateConflict{}
}

/*GetToolStateConflict handles this case with default header values.

If the printer is not operational
*/
type GetToolStateConflict struct {
}

func (o *GetToolStateConflict) Error() string {
	return fmt.Sprintf("[GET /api/printer/tool][%d] getToolStateConflict ", 409)
}

func (o *GetToolStateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}