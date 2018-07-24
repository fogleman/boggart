// Code generated by go-swagger; DO NOT EDIT.

package uuid

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

// NewGetInstanceUUIDParams creates a new GetInstanceUUIDParams object
// with the default values initialized.
func NewGetInstanceUUIDParams() *GetInstanceUUIDParams {

	return &GetInstanceUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstanceUUIDParamsWithTimeout creates a new GetInstanceUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstanceUUIDParamsWithTimeout(timeout time.Duration) *GetInstanceUUIDParams {

	return &GetInstanceUUIDParams{

		timeout: timeout,
	}
}

// NewGetInstanceUUIDParamsWithContext creates a new GetInstanceUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstanceUUIDParamsWithContext(ctx context.Context) *GetInstanceUUIDParams {

	return &GetInstanceUUIDParams{

		Context: ctx,
	}
}

// NewGetInstanceUUIDParamsWithHTTPClient creates a new GetInstanceUUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstanceUUIDParamsWithHTTPClient(client *http.Client) *GetInstanceUUIDParams {

	return &GetInstanceUUIDParams{
		HTTPClient: client,
	}
}

/*GetInstanceUUIDParams contains all the parameters to send to the API endpoint
for the get instance UUID operation typically these are written to a http.Request
*/
type GetInstanceUUIDParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instance UUID params
func (o *GetInstanceUUIDParams) WithTimeout(timeout time.Duration) *GetInstanceUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instance UUID params
func (o *GetInstanceUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instance UUID params
func (o *GetInstanceUUIDParams) WithContext(ctx context.Context) *GetInstanceUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instance UUID params
func (o *GetInstanceUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instance UUID params
func (o *GetInstanceUUIDParams) WithHTTPClient(client *http.Client) *GetInstanceUUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instance UUID params
func (o *GetInstanceUUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstanceUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
