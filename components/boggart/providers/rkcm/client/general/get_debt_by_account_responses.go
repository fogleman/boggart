// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/components/boggart/providers/rkcm/models"
)

// GetDebtByAccountReader is a Reader for the GetDebtByAccount structure.
type GetDebtByAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDebtByAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDebtByAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDebtByAccountOK creates a GetDebtByAccountOK with default headers values
func NewGetDebtByAccountOK() *GetDebtByAccountOK {
	return &GetDebtByAccountOK{}
}

/*GetDebtByAccountOK handles this case with default header values.

Successful operation
*/
type GetDebtByAccountOK struct {
	Payload *GetDebtByAccountOKBody
}

func (o *GetDebtByAccountOK) Error() string {
	return fmt.Sprintf("[GET /GetDebtByAccount.ashx][%d] getDebtByAccountOK  %+v", 200, o.Payload)
}

func (o *GetDebtByAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDebtByAccountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDebtByAccountOKBody get debt by account o k body
swagger:model GetDebtByAccountOKBody
*/
type GetDebtByAccountOKBody struct {

	// data
	Data *models.DebtForIdent `json:"data,omitempty"`
}

// Validate validates this get debt by account o k body
func (o *GetDebtByAccountOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDebtByAccountOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDebtByAccountOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDebtByAccountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDebtByAccountOKBody) UnmarshalBinary(b []byte) error {
	var res GetDebtByAccountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
