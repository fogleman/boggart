// Code generated by go-swagger; DO NOT EDIT.

package connection

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

// NewSendConnectionCommandParams creates a new SendConnectionCommandParams object
// with the default values initialized.
func NewSendConnectionCommandParams() *SendConnectionCommandParams {
	var ()
	return &SendConnectionCommandParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendConnectionCommandParamsWithTimeout creates a new SendConnectionCommandParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendConnectionCommandParamsWithTimeout(timeout time.Duration) *SendConnectionCommandParams {
	var ()
	return &SendConnectionCommandParams{

		timeout: timeout,
	}
}

// NewSendConnectionCommandParamsWithContext creates a new SendConnectionCommandParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendConnectionCommandParamsWithContext(ctx context.Context) *SendConnectionCommandParams {
	var ()
	return &SendConnectionCommandParams{

		Context: ctx,
	}
}

// NewSendConnectionCommandParamsWithHTTPClient creates a new SendConnectionCommandParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendConnectionCommandParamsWithHTTPClient(client *http.Client) *SendConnectionCommandParams {
	var ()
	return &SendConnectionCommandParams{
		HTTPClient: client,
	}
}

/*SendConnectionCommandParams contains all the parameters to send to the API endpoint
for the send connection command operation typically these are written to a http.Request
*/
type SendConnectionCommandParams struct {

	/*Body*/
	Body SendConnectionCommandBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send connection command params
func (o *SendConnectionCommandParams) WithTimeout(timeout time.Duration) *SendConnectionCommandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send connection command params
func (o *SendConnectionCommandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send connection command params
func (o *SendConnectionCommandParams) WithContext(ctx context.Context) *SendConnectionCommandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send connection command params
func (o *SendConnectionCommandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send connection command params
func (o *SendConnectionCommandParams) WithHTTPClient(client *http.Client) *SendConnectionCommandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send connection command params
func (o *SendConnectionCommandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send connection command params
func (o *SendConnectionCommandParams) WithBody(body SendConnectionCommandBody) *SendConnectionCommandParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send connection command params
func (o *SendConnectionCommandParams) SetBody(body SendConnectionCommandBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SendConnectionCommandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
