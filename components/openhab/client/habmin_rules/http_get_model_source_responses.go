// Code generated by go-swagger; DO NOT EDIT.

package habmin_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HTTPGetModelSourceReader is a Reader for the HTTPGetModelSource structure.
type HTTPGetModelSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPGetModelSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewHTTPGetModelSourceDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewHTTPGetModelSourceDefault creates a HTTPGetModelSourceDefault with default headers values
func NewHTTPGetModelSourceDefault(code int) *HTTPGetModelSourceDefault {
	return &HTTPGetModelSourceDefault{
		_statusCode: code,
	}
}

/*HTTPGetModelSourceDefault handles this case with default header values.

successful operation
*/
type HTTPGetModelSourceDefault struct {
	_statusCode int
}

// Code gets the status code for the http get model source default response
func (o *HTTPGetModelSourceDefault) Code() int {
	return o._statusCode
}

func (o *HTTPGetModelSourceDefault) Error() string {
	return fmt.Sprintf("[GET /habmin/rules/{modelname}][%d] httpGetModelSource default ", o._statusCode)
}

func (o *HTTPGetModelSourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}