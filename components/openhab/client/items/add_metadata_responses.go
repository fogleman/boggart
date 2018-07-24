// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddMetadataReader is a Reader for the AddMetadata structure.
type AddMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewAddMetadataCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewAddMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewAddMetadataMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddMetadataOK creates a AddMetadataOK with default headers values
func NewAddMetadataOK() *AddMetadataOK {
	return &AddMetadataOK{}
}

/*AddMetadataOK handles this case with default header values.

OK
*/
type AddMetadataOK struct {
}

func (o *AddMetadataOK) Error() string {
	return fmt.Sprintf("[PUT /items/{itemname}/metadata/{namespace}][%d] addMetadataOK ", 200)
}

func (o *AddMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddMetadataCreated creates a AddMetadataCreated with default headers values
func NewAddMetadataCreated() *AddMetadataCreated {
	return &AddMetadataCreated{}
}

/*AddMetadataCreated handles this case with default header values.

Created
*/
type AddMetadataCreated struct {
}

func (o *AddMetadataCreated) Error() string {
	return fmt.Sprintf("[PUT /items/{itemname}/metadata/{namespace}][%d] addMetadataCreated ", 201)
}

func (o *AddMetadataCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddMetadataNotFound creates a AddMetadataNotFound with default headers values
func NewAddMetadataNotFound() *AddMetadataNotFound {
	return &AddMetadataNotFound{}
}

/*AddMetadataNotFound handles this case with default header values.

Item not found.
*/
type AddMetadataNotFound struct {
}

func (o *AddMetadataNotFound) Error() string {
	return fmt.Sprintf("[PUT /items/{itemname}/metadata/{namespace}][%d] addMetadataNotFound ", 404)
}

func (o *AddMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddMetadataMethodNotAllowed creates a AddMetadataMethodNotAllowed with default headers values
func NewAddMetadataMethodNotAllowed() *AddMetadataMethodNotAllowed {
	return &AddMetadataMethodNotAllowed{}
}

/*AddMetadataMethodNotAllowed handles this case with default header values.

Metadata not editable.
*/
type AddMetadataMethodNotAllowed struct {
}

func (o *AddMetadataMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /items/{itemname}/metadata/{namespace}][%d] addMetadataMethodNotAllowed ", 405)
}

func (o *AddMetadataMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
