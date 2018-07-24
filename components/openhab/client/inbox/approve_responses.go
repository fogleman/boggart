// Code generated by go-swagger; DO NOT EDIT.

package inbox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ApproveReader is a Reader for the Approve structure.
type ApproveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApproveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewApproveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewApproveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewApproveConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewApproveOK creates a ApproveOK with default headers values
func NewApproveOK() *ApproveOK {
	return &ApproveOK{}
}

/*ApproveOK handles this case with default header values.

OK
*/
type ApproveOK struct {
}

func (o *ApproveOK) Error() string {
	return fmt.Sprintf("[POST /inbox/{thingUID}/approve][%d] approveOK ", 200)
}

func (o *ApproveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApproveNotFound creates a ApproveNotFound with default headers values
func NewApproveNotFound() *ApproveNotFound {
	return &ApproveNotFound{}
}

/*ApproveNotFound handles this case with default header values.

Thing not found in the inbox.
*/
type ApproveNotFound struct {
}

func (o *ApproveNotFound) Error() string {
	return fmt.Sprintf("[POST /inbox/{thingUID}/approve][%d] approveNotFound ", 404)
}

func (o *ApproveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApproveConflict creates a ApproveConflict with default headers values
func NewApproveConflict() *ApproveConflict {
	return &ApproveConflict{}
}

/*ApproveConflict handles this case with default header values.

No binding found that supports this thing.
*/
type ApproveConflict struct {
}

func (o *ApproveConflict) Error() string {
	return fmt.Sprintf("[POST /inbox/{thingUID}/approve][%d] approveConflict ", 409)
}

func (o *ApproveConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
