// Code generated by go-swagger; DO NOT EDIT.

package services

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

// NewDeleteConfigurationParams creates a new DeleteConfigurationParams object
// with the default values initialized.
func NewDeleteConfigurationParams() *DeleteConfigurationParams {
	var ()
	return &DeleteConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConfigurationParamsWithTimeout creates a new DeleteConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConfigurationParamsWithTimeout(timeout time.Duration) *DeleteConfigurationParams {
	var ()
	return &DeleteConfigurationParams{

		timeout: timeout,
	}
}

// NewDeleteConfigurationParamsWithContext creates a new DeleteConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConfigurationParamsWithContext(ctx context.Context) *DeleteConfigurationParams {
	var ()
	return &DeleteConfigurationParams{

		Context: ctx,
	}
}

// NewDeleteConfigurationParamsWithHTTPClient creates a new DeleteConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConfigurationParamsWithHTTPClient(client *http.Client) *DeleteConfigurationParams {
	var ()
	return &DeleteConfigurationParams{
		HTTPClient: client,
	}
}

/*DeleteConfigurationParams contains all the parameters to send to the API endpoint
for the delete configuration operation typically these are written to a http.Request
*/
type DeleteConfigurationParams struct {

	/*ServiceID
	  service ID

	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete configuration params
func (o *DeleteConfigurationParams) WithTimeout(timeout time.Duration) *DeleteConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete configuration params
func (o *DeleteConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete configuration params
func (o *DeleteConfigurationParams) WithContext(ctx context.Context) *DeleteConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete configuration params
func (o *DeleteConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete configuration params
func (o *DeleteConfigurationParams) WithHTTPClient(client *http.Client) *DeleteConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete configuration params
func (o *DeleteConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the delete configuration params
func (o *DeleteConfigurationParams) WithServiceID(serviceID string) *DeleteConfigurationParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the delete configuration params
func (o *DeleteConfigurationParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceId
	if err := r.SetPathParam("serviceId", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
