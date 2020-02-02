// Code generated by go-swagger; DO NOT EDIT.

package voice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// InterpretReader is a Reader for the Interpret structure.
type InterpretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterpretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterpretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInterpretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInterpretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterpretOK creates a InterpretOK with default headers values
func NewInterpretOK() *InterpretOK {
	return &InterpretOK{}
}

/*InterpretOK handles this case with default header values.

OK
*/
type InterpretOK struct {
}

func (o *InterpretOK) Error() string {
	return fmt.Sprintf("[POST /voice/interpreters/{id}][%d] interpretOK ", 200)
}

func (o *InterpretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInterpretBadRequest creates a InterpretBadRequest with default headers values
func NewInterpretBadRequest() *InterpretBadRequest {
	return &InterpretBadRequest{}
}

/*InterpretBadRequest handles this case with default header values.

interpretation exception occurs
*/
type InterpretBadRequest struct {
}

func (o *InterpretBadRequest) Error() string {
	return fmt.Sprintf("[POST /voice/interpreters/{id}][%d] interpretBadRequest ", 400)
}

func (o *InterpretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInterpretNotFound creates a InterpretNotFound with default headers values
func NewInterpretNotFound() *InterpretNotFound {
	return &InterpretNotFound{}
}

/*InterpretNotFound handles this case with default header values.

No human language interpreter was found.
*/
type InterpretNotFound struct {
}

func (o *InterpretNotFound) Error() string {
	return fmt.Sprintf("[POST /voice/interpreters/{id}][%d] interpretNotFound ", 404)
}

func (o *InterpretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}