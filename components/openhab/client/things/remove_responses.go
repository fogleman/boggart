// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveReader is a Reader for the Remove structure.
type RemoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewRemoveAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewRemoveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRemoveConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveOK creates a RemoveOK with default headers values
func NewRemoveOK() *RemoveOK {
	return &RemoveOK{}
}

/*RemoveOK handles this case with default header values.

OK, was deleted.
*/
type RemoveOK struct {
}

func (o *RemoveOK) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingUID}][%d] removeOK ", 200)
}

func (o *RemoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAccepted creates a RemoveAccepted with default headers values
func NewRemoveAccepted() *RemoveAccepted {
	return &RemoveAccepted{}
}

/*RemoveAccepted handles this case with default header values.

ACCEPTED for asynchronous deletion.
*/
type RemoveAccepted struct {
}

func (o *RemoveAccepted) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingUID}][%d] removeAccepted ", 202)
}

func (o *RemoveAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveNotFound creates a RemoveNotFound with default headers values
func NewRemoveNotFound() *RemoveNotFound {
	return &RemoveNotFound{}
}

/*RemoveNotFound handles this case with default header values.

Thing not found.
*/
type RemoveNotFound struct {
}

func (o *RemoveNotFound) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingUID}][%d] removeNotFound ", 404)
}

func (o *RemoveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveConflict creates a RemoveConflict with default headers values
func NewRemoveConflict() *RemoveConflict {
	return &RemoveConflict{}
}

/*RemoveConflict handles this case with default header values.

Thing could not be deleted because it's not editable.
*/
type RemoveConflict struct {
}

func (o *RemoveConflict) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingUID}][%d] removeConflict ", 409)
}

func (o *RemoveConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}