// Code generated by go-swagger; DO NOT EDIT.

package habmin_designer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutDesignRefReader is a Reader for the PutDesignRef structure.
type PutDesignRefReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDesignRefReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutDesignRefDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutDesignRefDefault creates a PutDesignRefDefault with default headers values
func NewPutDesignRefDefault(code int) *PutDesignRefDefault {
	return &PutDesignRefDefault{
		_statusCode: code,
	}
}

/*PutDesignRefDefault handles this case with default header values.

successful operation
*/
type PutDesignRefDefault struct {
	_statusCode int
}

// Code gets the status code for the put design ref default response
func (o *PutDesignRefDefault) Code() int {
	return o._statusCode
}

func (o *PutDesignRefDefault) Error() string {
	return fmt.Sprintf("[PUT /habmin/designer/{designref}][%d] putDesignRef default ", o._statusCode)
}

func (o *PutDesignRefDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
