// Code generated by go-swagger; DO NOT EDIT.

package audio

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDefaultSinkParams creates a new GetDefaultSinkParams object
// with the default values initialized.
func NewGetDefaultSinkParams() *GetDefaultSinkParams {
	var ()
	return &GetDefaultSinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDefaultSinkParamsWithTimeout creates a new GetDefaultSinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDefaultSinkParamsWithTimeout(timeout time.Duration) *GetDefaultSinkParams {
	var ()
	return &GetDefaultSinkParams{

		timeout: timeout,
	}
}

// NewGetDefaultSinkParamsWithContext creates a new GetDefaultSinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDefaultSinkParamsWithContext(ctx context.Context) *GetDefaultSinkParams {
	var ()
	return &GetDefaultSinkParams{

		Context: ctx,
	}
}

// NewGetDefaultSinkParamsWithHTTPClient creates a new GetDefaultSinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDefaultSinkParamsWithHTTPClient(client *http.Client) *GetDefaultSinkParams {
	var ()
	return &GetDefaultSinkParams{
		HTTPClient: client,
	}
}

/*GetDefaultSinkParams contains all the parameters to send to the API endpoint
for the get default sink operation typically these are written to a http.Request
*/
type GetDefaultSinkParams struct {

	/*AcceptLanguage
	  language

	*/
	AcceptLanguage *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get default sink params
func (o *GetDefaultSinkParams) WithTimeout(timeout time.Duration) *GetDefaultSinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get default sink params
func (o *GetDefaultSinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get default sink params
func (o *GetDefaultSinkParams) WithContext(ctx context.Context) *GetDefaultSinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get default sink params
func (o *GetDefaultSinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get default sink params
func (o *GetDefaultSinkParams) WithHTTPClient(client *http.Client) *GetDefaultSinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get default sink params
func (o *GetDefaultSinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get default sink params
func (o *GetDefaultSinkParams) WithAcceptLanguage(acceptLanguage *string) *GetDefaultSinkParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get default sink params
func (o *GetDefaultSinkParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WriteToRequest writes these params to a swagger request
func (o *GetDefaultSinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
