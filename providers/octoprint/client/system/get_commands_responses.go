// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/providers/octoprint/models"
)

// GetCommandsReader is a Reader for the GetCommands structure.
type GetCommandsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCommandsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCommandsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCommandsOK creates a GetCommandsOK with default headers values
func NewGetCommandsOK() *GetCommandsOK {
	return &GetCommandsOK{}
}

/*GetCommandsOK handles this case with default header values.

Successful operation
*/
type GetCommandsOK struct {
	Payload *GetCommandsOKBody
}

func (o *GetCommandsOK) Error() string {
	return fmt.Sprintf("[GET /api/system/commands][%d] getCommandsOK  %+v", 200, o.Payload)
}

func (o *GetCommandsOK) GetPayload() *GetCommandsOKBody {
	return o.Payload
}

func (o *GetCommandsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCommandsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCommandsOKBody get commands o k body
swagger:model GetCommandsOKBody
*/
type GetCommandsOKBody struct {

	// core
	Core []*models.Command `json:"core"`

	// custom
	Custom []*models.Command `json:"custom"`
}

// Validate validates this get commands o k body
func (o *GetCommandsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCore(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCustom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCommandsOKBody) validateCore(formats strfmt.Registry) error {

	if swag.IsZero(o.Core) { // not required
		return nil
	}

	for i := 0; i < len(o.Core); i++ {
		if swag.IsZero(o.Core[i]) { // not required
			continue
		}

		if o.Core[i] != nil {
			if err := o.Core[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCommandsOK" + "." + "core" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetCommandsOKBody) validateCustom(formats strfmt.Registry) error {

	if swag.IsZero(o.Custom) { // not required
		return nil
	}

	for i := 0; i < len(o.Custom); i++ {
		if swag.IsZero(o.Custom[i]) { // not required
			continue
		}

		if o.Custom[i] != nil {
			if err := o.Custom[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCommandsOK" + "." + "custom" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCommandsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCommandsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCommandsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
