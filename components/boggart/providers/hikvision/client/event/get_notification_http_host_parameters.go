// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNotificationHttpHostParams creates a new GetNotificationHttpHostParams object
// with the default values initialized.
func NewGetNotificationHttpHostParams() *GetNotificationHttpHostParams {
	var ()
	return &GetNotificationHttpHostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNotificationHttpHostParamsWithTimeout creates a new GetNotificationHttpHostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNotificationHttpHostParamsWithTimeout(timeout time.Duration) *GetNotificationHttpHostParams {
	var ()
	return &GetNotificationHttpHostParams{

		timeout: timeout,
	}
}

// NewGetNotificationHttpHostParamsWithContext creates a new GetNotificationHttpHostParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNotificationHttpHostParamsWithContext(ctx context.Context) *GetNotificationHttpHostParams {
	var ()
	return &GetNotificationHttpHostParams{

		Context: ctx,
	}
}

// NewGetNotificationHttpHostParamsWithHTTPClient creates a new GetNotificationHttpHostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNotificationHttpHostParamsWithHTTPClient(client *http.Client) *GetNotificationHttpHostParams {
	var ()
	return &GetNotificationHttpHostParams{
		HTTPClient: client,
	}
}

/*GetNotificationHttpHostParams contains all the parameters to send to the API endpoint
for the get notification http host operation typically these are written to a http.Request
*/
type GetNotificationHttpHostParams struct {

	/*HttpHost
	  HTTP host ID

	*/
	HttpHost uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get notification http host params
func (o *GetNotificationHttpHostParams) WithTimeout(timeout time.Duration) *GetNotificationHttpHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get notification http host params
func (o *GetNotificationHttpHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get notification http host params
func (o *GetNotificationHttpHostParams) WithContext(ctx context.Context) *GetNotificationHttpHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get notification http host params
func (o *GetNotificationHttpHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get notification http host params
func (o *GetNotificationHttpHostParams) WithHTTPClient(client *http.Client) *GetNotificationHttpHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get notification http host params
func (o *GetNotificationHttpHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHttpHost adds the httpHost to the get notification http host params
func (o *GetNotificationHttpHostParams) WithHttpHost(httpHost uint64) *GetNotificationHttpHostParams {
	o.SetHttpHost(httpHost)
	return o
}

// SetHttpHost adds the httpHost to the get notification http host params
func (o *GetNotificationHttpHostParams) SetHttpHost(httpHost uint64) {
	o.HttpHost = httpHost
}

// WriteToRequest writes these params to a swagger request
func (o *GetNotificationHttpHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param httpHost
	if err := r.SetPathParam("httpHost", swag.FormatUint64(o.HttpHost)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
