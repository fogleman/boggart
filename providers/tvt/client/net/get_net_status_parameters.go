// Code generated by go-swagger; DO NOT EDIT.

package net

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

	models "github.com/kihamo/boggart/providers/tvt/models"
)

// NewGetNetStatusParams creates a new GetNetStatusParams object
// with the default values initialized.
func NewGetNetStatusParams() *GetNetStatusParams {
	var ()
	return &GetNetStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetStatusParamsWithTimeout creates a new GetNetStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetStatusParamsWithTimeout(timeout time.Duration) *GetNetStatusParams {
	var ()
	return &GetNetStatusParams{

		timeout: timeout,
	}
}

// NewGetNetStatusParamsWithContext creates a new GetNetStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetStatusParamsWithContext(ctx context.Context) *GetNetStatusParams {
	var ()
	return &GetNetStatusParams{

		Context: ctx,
	}
}

// NewGetNetStatusParamsWithHTTPClient creates a new GetNetStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetStatusParamsWithHTTPClient(client *http.Client) *GetNetStatusParams {
	var ()
	return &GetNetStatusParams{
		HTTPClient: client,
	}
}

/*GetNetStatusParams contains all the parameters to send to the API endpoint
for the get net status operation typically these are written to a http.Request
*/
type GetNetStatusParams struct {

	/*Request*/
	Request models.Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get net status params
func (o *GetNetStatusParams) WithTimeout(timeout time.Duration) *GetNetStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get net status params
func (o *GetNetStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get net status params
func (o *GetNetStatusParams) WithContext(ctx context.Context) *GetNetStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get net status params
func (o *GetNetStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get net status params
func (o *GetNetStatusParams) WithHTTPClient(client *http.Client) *GetNetStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get net status params
func (o *GetNetStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the get net status params
func (o *GetNetStatusParams) WithRequest(request models.Request) *GetNetStatusParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the get net status params
func (o *GetNetStatusParams) SetRequest(request models.Request) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
