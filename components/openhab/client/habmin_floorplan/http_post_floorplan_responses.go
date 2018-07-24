// Code generated by go-swagger; DO NOT EDIT.

package habmin_floorplan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HTTPPostFloorplanReader is a Reader for the HTTPPostFloorplan structure.
type HTTPPostFloorplanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPPostFloorplanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewHTTPPostFloorplanDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewHTTPPostFloorplanDefault creates a HTTPPostFloorplanDefault with default headers values
func NewHTTPPostFloorplanDefault(code int) *HTTPPostFloorplanDefault {
	return &HTTPPostFloorplanDefault{
		_statusCode: code,
	}
}

/*HTTPPostFloorplanDefault handles this case with default header values.

successful operation
*/
type HTTPPostFloorplanDefault struct {
	_statusCode int
}

// Code gets the status code for the http post floorplan default response
func (o *HTTPPostFloorplanDefault) Code() int {
	return o._statusCode
}

func (o *HTTPPostFloorplanDefault) Error() string {
	return fmt.Sprintf("[POST /habmin/floorplan][%d] httpPostFloorplan default ", o._statusCode)
}

func (o *HTTPPostFloorplanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
