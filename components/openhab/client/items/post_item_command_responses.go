// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostItemCommandReader is a Reader for the PostItemCommand structure.
type PostItemCommandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostItemCommandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostItemCommandOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostItemCommandBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostItemCommandNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostItemCommandOK creates a PostItemCommandOK with default headers values
func NewPostItemCommandOK() *PostItemCommandOK {
	return &PostItemCommandOK{}
}

/*PostItemCommandOK handles this case with default header values.

OK
*/
type PostItemCommandOK struct {
}

func (o *PostItemCommandOK) Error() string {
	return fmt.Sprintf("[POST /items/{itemname}][%d] postItemCommandOK ", 200)
}

func (o *PostItemCommandOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostItemCommandBadRequest creates a PostItemCommandBadRequest with default headers values
func NewPostItemCommandBadRequest() *PostItemCommandBadRequest {
	return &PostItemCommandBadRequest{}
}

/*PostItemCommandBadRequest handles this case with default header values.

Item command null
*/
type PostItemCommandBadRequest struct {
}

func (o *PostItemCommandBadRequest) Error() string {
	return fmt.Sprintf("[POST /items/{itemname}][%d] postItemCommandBadRequest ", 400)
}

func (o *PostItemCommandBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostItemCommandNotFound creates a PostItemCommandNotFound with default headers values
func NewPostItemCommandNotFound() *PostItemCommandNotFound {
	return &PostItemCommandNotFound{}
}

/*PostItemCommandNotFound handles this case with default header values.

Item not found
*/
type PostItemCommandNotFound struct {
}

func (o *PostItemCommandNotFound) Error() string {
	return fmt.Sprintf("[POST /items/{itemname}][%d] postItemCommandNotFound ", 404)
}

func (o *PostItemCommandNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
