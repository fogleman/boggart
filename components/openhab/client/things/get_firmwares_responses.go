// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetFirmwaresReader is a Reader for the GetFirmwares structure.
type GetFirmwaresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFirmwaresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFirmwaresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetFirmwaresNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFirmwaresOK creates a GetFirmwaresOK with default headers values
func NewGetFirmwaresOK() *GetFirmwaresOK {
	return &GetFirmwaresOK{}
}

/*GetFirmwaresOK handles this case with default header values.

OK
*/
type GetFirmwaresOK struct {
}

func (o *GetFirmwaresOK) Error() string {
	return fmt.Sprintf("[GET /things/{thingUID}/firmwares][%d] getFirmwaresOK ", 200)
}

func (o *GetFirmwaresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFirmwaresNoContent creates a GetFirmwaresNoContent with default headers values
func NewGetFirmwaresNoContent() *GetFirmwaresNoContent {
	return &GetFirmwaresNoContent{}
}

/*GetFirmwaresNoContent handles this case with default header values.

No firmwares found.
*/
type GetFirmwaresNoContent struct {
}

func (o *GetFirmwaresNoContent) Error() string {
	return fmt.Sprintf("[GET /things/{thingUID}/firmwares][%d] getFirmwaresNoContent ", 204)
}

func (o *GetFirmwaresNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
