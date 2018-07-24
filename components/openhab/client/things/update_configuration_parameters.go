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

// NewUpdateConfigurationParams creates a new UpdateConfigurationParams object
// with the default values initialized.
func NewUpdateConfigurationParams() *UpdateConfigurationParams {
	var ()
	return &UpdateConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfigurationParamsWithTimeout creates a new UpdateConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateConfigurationParamsWithTimeout(timeout time.Duration) *UpdateConfigurationParams {
	var ()
	return &UpdateConfigurationParams{

		timeout: timeout,
	}
}

// NewUpdateConfigurationParamsWithContext creates a new UpdateConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateConfigurationParamsWithContext(ctx context.Context) *UpdateConfigurationParams {
	var ()
	return &UpdateConfigurationParams{

		Context: ctx,
	}
}

// NewUpdateConfigurationParamsWithHTTPClient creates a new UpdateConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateConfigurationParamsWithHTTPClient(client *http.Client) *UpdateConfigurationParams {
	var ()
	return &UpdateConfigurationParams{
		HTTPClient: client,
	}
}

/*UpdateConfigurationParams contains all the parameters to send to the API endpoint
for the update configuration operation typically these are written to a http.Request
*/
type UpdateConfigurationParams struct {

	/*AcceptLanguage*/
	AcceptLanguage *string
	/*Body
	  configuration parameters

	*/
	Body map[string]interface{}
	/*ThingUID
	  thing

	*/
	ThingUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update configuration params
func (o *UpdateConfigurationParams) WithTimeout(timeout time.Duration) *UpdateConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update configuration params
func (o *UpdateConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update configuration params
func (o *UpdateConfigurationParams) WithContext(ctx context.Context) *UpdateConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update configuration params
func (o *UpdateConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update configuration params
func (o *UpdateConfigurationParams) WithHTTPClient(client *http.Client) *UpdateConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update configuration params
func (o *UpdateConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the update configuration params
func (o *UpdateConfigurationParams) WithAcceptLanguage(acceptLanguage *string) *UpdateConfigurationParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the update configuration params
func (o *UpdateConfigurationParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithBody adds the body to the update configuration params
func (o *UpdateConfigurationParams) WithBody(body map[string]interface{}) *UpdateConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update configuration params
func (o *UpdateConfigurationParams) SetBody(body map[string]interface{}) {
	o.Body = body
}

// WithThingUID adds the thingUID to the update configuration params
func (o *UpdateConfigurationParams) WithThingUID(thingUID string) *UpdateConfigurationParams {
	o.SetThingUID(thingUID)
	return o
}

// SetThingUID adds the thingUid to the update configuration params
func (o *UpdateConfigurationParams) SetThingUID(thingUID string) {
	o.ThingUID = thingUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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
