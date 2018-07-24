// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetByUIDParams creates a new GetByUIDParams object
// with the default values initialized.
func NewGetByUIDParams() *GetByUIDParams {
	var ()
	return &GetByUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetByUIDParamsWithTimeout creates a new GetByUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetByUIDParamsWithTimeout(timeout time.Duration) *GetByUIDParams {
	var ()
	return &GetByUIDParams{

		timeout: timeout,
	}
}

// NewGetByUIDParamsWithContext creates a new GetByUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetByUIDParamsWithContext(ctx context.Context) *GetByUIDParams {
	var ()
	return &GetByUIDParams{

		Context: ctx,
	}
}

// NewGetByUIDParamsWithHTTPClient creates a new GetByUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetByUIDParamsWithHTTPClient(client *http.Client) *GetByUIDParams {
	var ()
	return &GetByUIDParams{
		HTTPClient: client,
	}
}

/*GetByUIDParams contains all the parameters to send to the API endpoint
for the get by UID operation typically these are written to a http.Request
*/
type GetByUIDParams struct {

	/*AcceptLanguage
	  language

	*/
	AcceptLanguage *string
	/*ThingUID
	  thingUID

	*/
	ThingUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get by UID params
func (o *GetByUIDParams) WithTimeout(timeout time.Duration) *GetByUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by UID params
func (o *GetByUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by UID params
func (o *GetByUIDParams) WithContext(ctx context.Context) *GetByUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by UID params
func (o *GetByUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by UID params
func (o *GetByUIDParams) WithHTTPClient(client *http.Client) *GetByUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by UID params
func (o *GetByUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get by UID params
func (o *GetByUIDParams) WithAcceptLanguage(acceptLanguage *string) *GetByUIDParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get by UID params
func (o *GetByUIDParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithThingUID adds the thingUID to the get by UID params
func (o *GetByUIDParams) WithThingUID(thingUID string) *GetByUIDParams {
	o.SetThingUID(thingUID)
	return o
}

// SetThingUID adds the thingUid to the get by UID params
func (o *GetByUIDParams) SetThingUID(thingUID string) {
	o.ThingUID = thingUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetByUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}

	}

	// path param thingUID
	if err := r.SetPathParam("thingUID", o.ThingUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
