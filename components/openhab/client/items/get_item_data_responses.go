// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/components/openhab/models"
)

// GetItemDataReader is a Reader for the GetItemData structure.
type GetItemDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetItemDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetItemDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetItemDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetItemDataOK creates a GetItemDataOK with default headers values
func NewGetItemDataOK() *GetItemDataOK {
	return &GetItemDataOK{}
}

/*GetItemDataOK handles this case with default header values.

OK
*/
type GetItemDataOK struct {
	Payload *models.EnrichedItemDTO
}

func (o *GetItemDataOK) Error() string {
	return fmt.Sprintf("[GET /items/{itemname}][%d] getItemDataOK  %+v", 200, o.Payload)
}

func (o *GetItemDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnrichedItemDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemDataNotFound creates a GetItemDataNotFound with default headers values
func NewGetItemDataNotFound() *GetItemDataNotFound {
	return &GetItemDataNotFound{}
}

/*GetItemDataNotFound handles this case with default header values.

Item not found
*/
type GetItemDataNotFound struct {
}

func (o *GetItemDataNotFound) Error() string {
	return fmt.Sprintf("[GET /items/{itemname}][%d] getItemDataNotFound ", 404)
}

func (o *GetItemDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
