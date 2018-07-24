// Code generated by go-swagger; DO NOT EDIT.

package voice

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

// NewGetInterpreterParams creates a new GetInterpreterParams object
// with the default values initialized.
func NewGetInterpreterParams() *GetInterpreterParams {
	var ()
	return &GetInterpreterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInterpreterParamsWithTimeout creates a new GetInterpreterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInterpreterParamsWithTimeout(timeout time.Duration) *GetInterpreterParams {
	var ()
	return &GetInterpreterParams{

		timeout: timeout,
	}
}

// NewGetInterpreterParamsWithContext creates a new GetInterpreterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInterpreterParamsWithContext(ctx context.Context) *GetInterpreterParams {
	var ()
	return &GetInterpreterParams{

		Context: ctx,
	}
}

// NewGetInterpreterParamsWithHTTPClient creates a new GetInterpreterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInterpreterParamsWithHTTPClient(client *http.Client) *GetInterpreterParams {
	var ()
	return &GetInterpreterParams{
		HTTPClient: client,
	}
}

/*GetInterpreterParams contains all the parameters to send to the API endpoint
for the get interpreter operation typically these are written to a http.Request
*/
type GetInterpreterParams struct {

	/*AcceptLanguage
	  language

	*/
	AcceptLanguage *string
	/*ID
	  interpreter id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get interpreter params
func (o *GetInterpreterParams) WithTimeout(timeout time.Duration) *GetInterpreterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get interpreter params
func (o *GetInterpreterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get interpreter params
func (o *GetInterpreterParams) WithContext(ctx context.Context) *GetInterpreterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get interpreter params
func (o *GetInterpreterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get interpreter params
func (o *GetInterpreterParams) WithHTTPClient(client *http.Client) *GetInterpreterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get interpreter params
func (o *GetInterpreterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get interpreter params
func (o *GetInterpreterParams) WithAcceptLanguage(acceptLanguage *string) *GetInterpreterParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get interpreter params
func (o *GetInterpreterParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithID adds the id to the get interpreter params
func (o *GetInterpreterParams) WithID(id string) *GetInterpreterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get interpreter params
func (o *GetInterpreterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetInterpreterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
