// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/components/boggart/providers/nalog.ru/proverkacheka/models"
)

// GetReader is a Reader for the Get structure.
type GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOK creates a GetOK with default headers values
func NewGetOK() *GetOK {
	return &GetOK{}
}

/*GetOK handles this case with default header values.

Successful operation
*/
type GetOK struct {
	Payload *GetOKBody
}

func (o *GetOK) Error() string {
	return fmt.Sprintf("[GET /inns/*/kkts/*/fss/{fiscalDriveNumber}/tickets/{fiscalDocumentNumber}][%d] getOK  %+v", 200, o.Payload)
}

func (o *GetOK) GetPayload() *GetOKBody {
	return o.Payload
}

func (o *GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOKBody get o k body
swagger:model GetOKBody
*/
type GetOKBody struct {

	// document
	Document *models.Document `json:"document,omitempty"`
}

// Validate validates this get o k body
func (o *GetOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDocument(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOKBody) validateDocument(formats strfmt.Registry) error {

	if swag.IsZero(o.Document) { // not required
		return nil
	}

	if o.Document != nil {
		if err := o.Document.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOK" + "." + "document")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOKBody) UnmarshalBinary(b []byte) error {
	var res GetOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
