// Code generated by go-swagger; DO NOT EDIT.

package habmin_floorplan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HTTPGetFloorplanImageReader is a Reader for the HTTPGetFloorplanImage structure.
type HTTPGetFloorplanImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPGetFloorplanImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewHTTPGetFloorplanImageDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewHTTPGetFloorplanImageDefault creates a HTTPGetFloorplanImageDefault with default headers values
func NewHTTPGetFloorplanImageDefault(code int) *HTTPGetFloorplanImageDefault {
	return &HTTPGetFloorplanImageDefault{
		_statusCode: code,
	}
}

/*HTTPGetFloorplanImageDefault handles this case with default header values.

successful operation
*/
type HTTPGetFloorplanImageDefault struct {
	_statusCode int
}

// Code gets the status code for the http get floorplan image default response
func (o *HTTPGetFloorplanImageDefault) Code() int {
	return o._statusCode
}

func (o *HTTPGetFloorplanImageDefault) Error() string {
	return fmt.Sprintf("[GET /habmin/floorplan/{floorplanID}/image][%d] httpGetFloorplanImage default ", o._statusCode)
}

func (o *HTTPGetFloorplanImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
