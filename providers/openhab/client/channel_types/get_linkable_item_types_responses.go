// Code generated by go-swagger; DO NOT EDIT.

package channel_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLinkableItemTypesReader is a Reader for the GetLinkableItemTypes structure.
type GetLinkableItemTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLinkableItemTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLinkableItemTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetLinkableItemTypesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLinkableItemTypesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLinkableItemTypesOK creates a GetLinkableItemTypesOK with default headers values
func NewGetLinkableItemTypesOK() *GetLinkableItemTypesOK {
	return &GetLinkableItemTypesOK{}
}

/*GetLinkableItemTypesOK handles this case with default header values.

OK
*/
type GetLinkableItemTypesOK struct {
	Payload []string
}

func (o *GetLinkableItemTypesOK) Error() string {
	return fmt.Sprintf("[GET /channel-types/{channelTypeUID}/linkableItemTypes][%d] getLinkableItemTypesOK  %+v", 200, o.Payload)
}

func (o *GetLinkableItemTypesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetLinkableItemTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLinkableItemTypesNoContent creates a GetLinkableItemTypesNoContent with default headers values
func NewGetLinkableItemTypesNoContent() *GetLinkableItemTypesNoContent {
	return &GetLinkableItemTypesNoContent{}
}

/*GetLinkableItemTypesNoContent handles this case with default header values.

No content: channel type has no linkable items or is no trigger channel.
*/
type GetLinkableItemTypesNoContent struct {
}

func (o *GetLinkableItemTypesNoContent) Error() string {
	return fmt.Sprintf("[GET /channel-types/{channelTypeUID}/linkableItemTypes][%d] getLinkableItemTypesNoContent ", 204)
}

func (o *GetLinkableItemTypesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLinkableItemTypesNotFound creates a GetLinkableItemTypesNotFound with default headers values
func NewGetLinkableItemTypesNotFound() *GetLinkableItemTypesNotFound {
	return &GetLinkableItemTypesNotFound{}
}

/*GetLinkableItemTypesNotFound handles this case with default header values.

Given channel type UID not found.
*/
type GetLinkableItemTypesNotFound struct {
}

func (o *GetLinkableItemTypesNotFound) Error() string {
	return fmt.Sprintf("[GET /channel-types/{channelTypeUID}/linkableItemTypes][%d] getLinkableItemTypesNotFound ", 404)
}

func (o *GetLinkableItemTypesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}