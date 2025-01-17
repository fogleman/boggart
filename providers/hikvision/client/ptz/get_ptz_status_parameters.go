// Code generated by go-swagger; DO NOT EDIT.

package ptz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetPtzStatusParams creates a new GetPtzStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPtzStatusParams() *GetPtzStatusParams {
	return &GetPtzStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPtzStatusParamsWithTimeout creates a new GetPtzStatusParams object
// with the ability to set a timeout on a request.
func NewGetPtzStatusParamsWithTimeout(timeout time.Duration) *GetPtzStatusParams {
	return &GetPtzStatusParams{
		timeout: timeout,
	}
}

// NewGetPtzStatusParamsWithContext creates a new GetPtzStatusParams object
// with the ability to set a context for a request.
func NewGetPtzStatusParamsWithContext(ctx context.Context) *GetPtzStatusParams {
	return &GetPtzStatusParams{
		Context: ctx,
	}
}

// NewGetPtzStatusParamsWithHTTPClient creates a new GetPtzStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPtzStatusParamsWithHTTPClient(client *http.Client) *GetPtzStatusParams {
	return &GetPtzStatusParams{
		HTTPClient: client,
	}
}

/* GetPtzStatusParams contains all the parameters to send to the API endpoint
   for the get ptz status operation.

   Typically these are written to a http.Request.
*/
type GetPtzStatusParams struct {

	/* Channel.

	   Channel ID

	   Format: uint64
	*/
	Channel uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ptz status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPtzStatusParams) WithDefaults() *GetPtzStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ptz status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPtzStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ptz status params
func (o *GetPtzStatusParams) WithTimeout(timeout time.Duration) *GetPtzStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ptz status params
func (o *GetPtzStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ptz status params
func (o *GetPtzStatusParams) WithContext(ctx context.Context) *GetPtzStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ptz status params
func (o *GetPtzStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ptz status params
func (o *GetPtzStatusParams) WithHTTPClient(client *http.Client) *GetPtzStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ptz status params
func (o *GetPtzStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannel adds the channel to the get ptz status params
func (o *GetPtzStatusParams) WithChannel(channel uint64) *GetPtzStatusParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the get ptz status params
func (o *GetPtzStatusParams) SetChannel(channel uint64) {
	o.Channel = channel
}

// WriteToRequest writes these params to a swagger request
func (o *GetPtzStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channel
	if err := r.SetPathParam("channel", swag.FormatUint64(o.Channel)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
