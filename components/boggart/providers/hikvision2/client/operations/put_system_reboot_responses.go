// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutSystemRebootReader is a Reader for the PutSystemReboot structure.
type PutSystemRebootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSystemRebootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutSystemRebootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutSystemRebootOK creates a PutSystemRebootOK with default headers values
func NewPutSystemRebootOK() *PutSystemRebootOK {
	return &PutSystemRebootOK{}
}

/*PutSystemRebootOK handles this case with default header values.

PutSystemRebootOK put system reboot o k
*/
type PutSystemRebootOK struct {
}

func (o *PutSystemRebootOK) Error() string {
	return fmt.Sprintf("[PUT /System/reboot][%d] putSystemRebootOK ", 200)
}

func (o *PutSystemRebootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
