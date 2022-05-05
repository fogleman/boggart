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

	"github.com/kihamo/boggart/providers/hikvision/models"
)

// NewSetPtzContinuousParams creates a new SetPtzContinuousParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetPtzContinuousParams() *SetPtzContinuousParams {
	return &SetPtzContinuousParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetPtzContinuousParamsWithTimeout creates a new SetPtzContinuousParams object
// with the ability to set a timeout on a request.
func NewSetPtzContinuousParamsWithTimeout(timeout time.Duration) *SetPtzContinuousParams {
	return &SetPtzContinuousParams{
		timeout: timeout,
	}
}

// NewSetPtzContinuousParamsWithContext creates a new SetPtzContinuousParams object
// with the ability to set a context for a request.
func NewSetPtzContinuousParamsWithContext(ctx context.Context) *SetPtzContinuousParams {
	return &SetPtzContinuousParams{
		Context: ctx,
	}
}

// NewSetPtzContinuousParamsWithHTTPClient creates a new SetPtzContinuousParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetPtzContinuousParamsWithHTTPClient(client *http.Client) *SetPtzContinuousParams {
	return &SetPtzContinuousParams{
		HTTPClient: client,
	}
}

/* SetPtzContinuousParams contains all the parameters to send to the API endpoint
   for the set ptz continuous operation.

   Typically these are written to a http.Request.
*/
type SetPtzContinuousParams struct {

	// PTZData.
	PTZData *models.PTZData

	/* Channel.

	   Channel ID

	   Format: uint64
	*/
	Channel uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set ptz continuous params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetPtzContinuousParams) WithDefaults() *SetPtzContinuousParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set ptz continuous params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetPtzContinuousParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set ptz continuous params
func (o *SetPtzContinuousParams) WithTimeout(timeout time.Duration) *SetPtzContinuousParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set ptz continuous params
func (o *SetPtzContinuousParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set ptz continuous params
func (o *SetPtzContinuousParams) WithContext(ctx context.Context) *SetPtzContinuousParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set ptz continuous params
func (o *SetPtzContinuousParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set ptz continuous params
func (o *SetPtzContinuousParams) WithHTTPClient(client *http.Client) *SetPtzContinuousParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set ptz continuous params
func (o *SetPtzContinuousParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPTZData adds the pTZData to the set ptz continuous params
func (o *SetPtzContinuousParams) WithPTZData(pTZData *models.PTZData) *SetPtzContinuousParams {
	o.SetPTZData(pTZData)
	return o
}

// SetPTZData adds the pTZData to the set ptz continuous params
func (o *SetPtzContinuousParams) SetPTZData(pTZData *models.PTZData) {
	o.PTZData = pTZData
}

// WithChannel adds the channel to the set ptz continuous params
func (o *SetPtzContinuousParams) WithChannel(channel uint64) *SetPtzContinuousParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the set ptz continuous params
func (o *SetPtzContinuousParams) SetChannel(channel uint64) {
	o.Channel = channel
}

// WriteToRequest writes these params to a swagger request
func (o *SetPtzContinuousParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PTZData != nil {
		if err := r.SetBodyParam(o.PTZData); err != nil {
			return err
		}
	}

	// path param channel
	if err := r.SetPathParam("channel", swag.FormatUint64(o.Channel)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
