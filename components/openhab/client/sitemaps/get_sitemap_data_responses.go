// Code generated by go-swagger; DO NOT EDIT.

package sitemaps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetSitemapDataReader is a Reader for the GetSitemapData structure.
type GetSitemapDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitemapDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSitemapDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSitemapDataOK creates a GetSitemapDataOK with default headers values
func NewGetSitemapDataOK() *GetSitemapDataOK {
	return &GetSitemapDataOK{}
}

/*GetSitemapDataOK handles this case with default header values.

OK
*/
type GetSitemapDataOK struct {
}

func (o *GetSitemapDataOK) Error() string {
	return fmt.Sprintf("[GET /sitemaps/{sitemapname}][%d] getSitemapDataOK ", 200)
}

func (o *GetSitemapDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
