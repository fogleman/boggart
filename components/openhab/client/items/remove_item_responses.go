// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveItemReader is a Reader for the RemoveItem structure.
type RemoveItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewRemoveItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveItemOK creates a RemoveItemOK with default headers values
func NewRemoveItemOK() *RemoveItemOK {
	return &RemoveItemOK{}
}

/*RemoveItemOK handles this case with default header values.

OK
*/
type RemoveItemOK struct {
}

func (o *RemoveItemOK) Error() string {
	return fmt.Sprintf("[DELETE /items/{itemname}][%d] removeItemOK ", 200)
}

func (o *RemoveItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveItemNotFound creates a RemoveItemNotFound with default headers values
func NewRemoveItemNotFound() *RemoveItemNotFound {
	return &RemoveItemNotFound{}
}

/*RemoveItemNotFound handles this case with default header values.

Item not found or item is not editable.
*/
type RemoveItemNotFound struct {
}

func (o *RemoveItemNotFound) Error() string {
	return fmt.Sprintf("[DELETE /items/{itemname}][%d] removeItemNotFound ", 404)
}

func (o *RemoveItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
