// Code generated by go-swagger; DO NOT EDIT.

package sitemaps

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

// NewCreateEventSubscriptionParams creates a new CreateEventSubscriptionParams object
// with the default values initialized.
func NewCreateEventSubscriptionParams() *CreateEventSubscriptionParams {

	return &CreateEventSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEventSubscriptionParamsWithTimeout creates a new CreateEventSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateEventSubscriptionParamsWithTimeout(timeout time.Duration) *CreateEventSubscriptionParams {

	return &CreateEventSubscriptionParams{

		timeout: timeout,
	}
}

// NewCreateEventSubscriptionParamsWithContext creates a new CreateEventSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateEventSubscriptionParamsWithContext(ctx context.Context) *CreateEventSubscriptionParams {

	return &CreateEventSubscriptionParams{

		Context: ctx,
	}
}

// NewCreateEventSubscriptionParamsWithHTTPClient creates a new CreateEventSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateEventSubscriptionParamsWithHTTPClient(client *http.Client) *CreateEventSubscriptionParams {

	return &CreateEventSubscriptionParams{
		HTTPClient: client,
	}
}

/*CreateEventSubscriptionParams contains all the parameters to send to the API endpoint
for the create event subscription operation typically these are written to a http.Request
*/
type CreateEventSubscriptionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create event subscription params
func (o *CreateEventSubscriptionParams) WithTimeout(timeout time.Duration) *CreateEventSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create event subscription params
func (o *CreateEventSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create event subscription params
func (o *CreateEventSubscriptionParams) WithContext(ctx context.Context) *CreateEventSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create event subscription params
func (o *CreateEventSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create event subscription params
func (o *CreateEventSubscriptionParams) WithHTTPClient(client *http.Client) *CreateEventSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create event subscription params
func (o *CreateEventSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEventSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
