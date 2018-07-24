// Code generated by go-swagger; DO NOT EDIT.

package config_descriptions

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

// NewGetByURIParams creates a new GetByURIParams object
// with the default values initialized.
func NewGetByURIParams() *GetByURIParams {
	var ()
	return &GetByURIParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetByURIParamsWithTimeout creates a new GetByURIParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetByURIParamsWithTimeout(timeout time.Duration) *GetByURIParams {
	var ()
	return &GetByURIParams{

		timeout: timeout,
	}
}

// NewGetByURIParamsWithContext creates a new GetByURIParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetByURIParamsWithContext(ctx context.Context) *GetByURIParams {
	var ()
	return &GetByURIParams{

		Context: ctx,
	}
}

// NewGetByURIParamsWithHTTPClient creates a new GetByURIParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetByURIParamsWithHTTPClient(client *http.Client) *GetByURIParams {
	var ()
	return &GetByURIParams{
		HTTPClient: client,
	}
}

/*GetByURIParams contains all the parameters to send to the API endpoint
for the get by URI operation typically these are written to a http.Request
*/
type GetByURIParams struct {

	/*AcceptLanguage
	  Accept-Language

	*/
	AcceptLanguage *string
	/*URI
	  uri

	*/
	URI string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get by URI params
func (o *GetByURIParams) WithTimeout(timeout time.Duration) *GetByURIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by URI params
func (o *GetByURIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by URI params
func (o *GetByURIParams) WithContext(ctx context.Context) *GetByURIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by URI params
func (o *GetByURIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by URI params
func (o *GetByURIParams) WithHTTPClient(client *http.Client) *GetByURIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by URI params
func (o *GetByURIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get by URI params
func (o *GetByURIParams) WithAcceptLanguage(acceptLanguage *string) *GetByURIParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get by URI params
func (o *GetByURIParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithURI adds the uri to the get by URI params
func (o *GetByURIParams) WithURI(uri string) *GetByURIParams {
	o.SetURI(uri)
	return o
}

// SetURI adds the uri to the get by URI params
func (o *GetByURIParams) SetURI(uri string) {
	o.URI = uri
}

// WriteToRequest writes these params to a swagger request
func (o *GetByURIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uri
	if err := r.SetPathParam("uri", o.URI); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
