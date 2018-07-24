// Code generated by go-swagger; DO NOT EDIT.

package persistence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HTTPGetPersistenceServiceItemsReader is a Reader for the HTTPGetPersistenceServiceItems structure.
type HTTPGetPersistenceServiceItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPGetPersistenceServiceItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHTTPGetPersistenceServiceItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHTTPGetPersistenceServiceItemsOK creates a HTTPGetPersistenceServiceItemsOK with default headers values
func NewHTTPGetPersistenceServiceItemsOK() *HTTPGetPersistenceServiceItemsOK {
	return &HTTPGetPersistenceServiceItemsOK{}
}

/*HTTPGetPersistenceServiceItemsOK handles this case with default header values.

OK
*/
type HTTPGetPersistenceServiceItemsOK struct {
	Payload []string
}

func (o *HTTPGetPersistenceServiceItemsOK) Error() string {
	return fmt.Sprintf("[GET /persistence/items][%d] httpGetPersistenceServiceItemsOK  %+v", 200, o.Payload)
}

func (o *HTTPGetPersistenceServiceItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
