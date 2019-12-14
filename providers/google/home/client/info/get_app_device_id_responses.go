// Code generated by go-swagger; DO NOT EDIT.

package info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAppDeviceIDReader is a Reader for the GetAppDeviceID structure.
type GetAppDeviceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppDeviceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppDeviceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppDeviceIDOK creates a GetAppDeviceIDOK with default headers values
func NewGetAppDeviceIDOK() *GetAppDeviceIDOK {
	return &GetAppDeviceIDOK{}
}

/*GetAppDeviceIDOK handles this case with default header values.

successful operation
*/
type GetAppDeviceIDOK struct {
	Payload *GetAppDeviceIDOKBody
}

func (o *GetAppDeviceIDOK) Error() string {
	return fmt.Sprintf("[POST /setup/get_app_device_id][%d] getAppDeviceIdOK  %+v", 200, o.Payload)
}

func (o *GetAppDeviceIDOK) GetPayload() *GetAppDeviceIDOKBody {
	return o.Payload
}

func (o *GetAppDeviceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAppDeviceIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAppDeviceIDBody get app device ID body
swagger:model GetAppDeviceIDBody
*/
type GetAppDeviceIDBody struct {

	// app id
	// Required: true
	// Enum: [E8C28D3C]
	AppID *string `json:"app_id"`
}

// Validate validates this get app device ID body
func (o *GetAppDeviceIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAppID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getAppDeviceIdBodyTypeAppIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["E8C28D3C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getAppDeviceIdBodyTypeAppIDPropEnum = append(getAppDeviceIdBodyTypeAppIDPropEnum, v)
	}
}

const (

	// GetAppDeviceIDBodyAppIDE8C28D3C captures enum value "E8C28D3C"
	GetAppDeviceIDBodyAppIDE8C28D3C string = "E8C28D3C"
)

// prop value enum
func (o *GetAppDeviceIDBody) validateAppIDEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getAppDeviceIdBodyTypeAppIDPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetAppDeviceIDBody) validateAppID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"app_id", "body", o.AppID); err != nil {
		return err
	}

	// value enum
	if err := o.validateAppIDEnum("body"+"."+"app_id", "body", *o.AppID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAppDeviceIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAppDeviceIDBody) UnmarshalBinary(b []byte) error {
	var res GetAppDeviceIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAppDeviceIDOKBody get app device ID o k body
swagger:model GetAppDeviceIDOKBody
*/
type GetAppDeviceIDOKBody struct {

	// app device id
	// Required: true
	AppDeviceID *string `json:"app_device_id"`

	// certificate
	// Required: true
	Certificate *string `json:"certificate"`

	// signed data
	// Required: true
	SignedData *string `json:"signed_data"`
}

// Validate validates this get app device ID o k body
func (o *GetAppDeviceIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAppDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSignedData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAppDeviceIDOKBody) validateAppDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("getAppDeviceIdOK"+"."+"app_device_id", "body", o.AppDeviceID); err != nil {
		return err
	}

	return nil
}

func (o *GetAppDeviceIDOKBody) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("getAppDeviceIdOK"+"."+"certificate", "body", o.Certificate); err != nil {
		return err
	}

	return nil
}

func (o *GetAppDeviceIDOKBody) validateSignedData(formats strfmt.Registry) error {

	if err := validate.Required("getAppDeviceIdOK"+"."+"signed_data", "body", o.SignedData); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAppDeviceIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAppDeviceIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetAppDeviceIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
