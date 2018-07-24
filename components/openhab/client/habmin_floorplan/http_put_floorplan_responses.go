// Code generated by go-swagger; DO NOT EDIT.

package habmin_floorplan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HTTPPutFloorplanReader is a Reader for the HTTPPutFloorplan structure.
type HTTPPutFloorplanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPPutFloorplanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewHTTPPutFloorplanDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewHTTPPutFloorplanDefault creates a HTTPPutFloorplanDefault with default headers values
func NewHTTPPutFloorplanDefault(code int) *HTTPPutFloorplanDefault {
	return &HTTPPutFloorplanDefault{
		_statusCode: code,
	}
}

/*HTTPPutFloorplanDefault handles this case with default header values.

successful operation
*/
type HTTPPutFloorplanDefault struct {
	_statusCode int
}

// Code gets the status code for the http put floorplan default response
func (o *HTTPPutFloorplanDefault) Code() int {
	return o._statusCode
}

func (o *HTTPPutFloorplanDefault) Error() string {
	return fmt.Sprintf("[PUT /habmin/floorplan/{floorplanID}][%d] httpPutFloorplan default ", o._statusCode)
}

func (o *HTTPPutFloorplanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
