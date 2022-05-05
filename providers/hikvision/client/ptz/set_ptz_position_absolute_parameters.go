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

// NewSetPtzPositionAbsoluteParams creates a new SetPtzPositionAbsoluteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetPtzPositionAbsoluteParams() *SetPtzPositionAbsoluteParams {
	return &SetPtzPositionAbsoluteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetPtzPositionAbsoluteParamsWithTimeout creates a new SetPtzPositionAbsoluteParams object
// with the ability to set a timeout on a request.
func NewSetPtzPositionAbsoluteParamsWithTimeout(timeout time.Duration) *SetPtzPositionAbsoluteParams {
	return &SetPtzPositionAbsoluteParams{
		timeout: timeout,
	}
}

// NewSetPtzPositionAbsoluteParamsWithContext creates a new SetPtzPositionAbsoluteParams object
// with the ability to set a context for a request.
func NewSetPtzPositionAbsoluteParamsWithContext(ctx context.Context) *SetPtzPositionAbsoluteParams {
	return &SetPtzPositionAbsoluteParams{
		Context: ctx,
	}
}

// NewSetPtzPositionAbsoluteParamsWithHTTPClient creates a new SetPtzPositionAbsoluteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetPtzPositionAbsoluteParamsWithHTTPClient(client *http.Client) *SetPtzPositionAbsoluteParams {
	return &SetPtzPositionAbsoluteParams{
		HTTPClient: client,
	}
}

/* SetPtzPositionAbsoluteParams contains all the parameters to send to the API endpoint
   for the set ptz position absolute operation.

   Typically these are written to a http.Request.
*/
type SetPtzPositionAbsoluteParams struct {

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

// WithDefaults hydrates default values in the set ptz position absolute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetPtzPositionAbsoluteParams) WithDefaults() *SetPtzPositionAbsoluteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set ptz position absolute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetPtzPositionAbsoluteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set ptz position absolute params
func (o *SetPtzPositionAbsoluteParams) WithTimeout(timeout time.Duration) *SetPtzPositionAbsoluteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set ptz position absolute params
func (o *SetPtzPositionAbsoluteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set ptz position absolute params
func (o *SetPtzPositionAbsoluteParams) WithContext(ctx context.Context) *SetPtzPositionAbsoluteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set ptz position absolute params
func (o *SetPtzPositionAbsoluteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set ptz position absolute params
func (o *SetPtzPositionAbsoluteParams) WithHTTPClient(client *http.Client) *SetPtzPositionAbsoluteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set ptz position absolute params
func (o *SetPtzPositionAbsoluteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPTZData adds the pTZData to the set ptz position absolute params
func (o *SetPtzPositionAbsoluteParams) WithPTZData(pTZData *models.PTZData) *SetPtzPositionAbsoluteParams {
	o.SetPTZData(pTZData)
	return o
}

// SetPTZData adds the pTZData to the set ptz position absolute params
func (o *SetPtzPositionAbsoluteParams) SetPTZData(pTZData *models.PTZData) {
	o.PTZData = pTZData
}

// WithChannel adds the channel to the set ptz position absolute params
func (o *SetPtzPositionAbsoluteParams) WithChannel(channel uint64) *SetPtzPositionAbsoluteParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the set ptz position absolute params
func (o *SetPtzPositionAbsoluteParams) SetChannel(channel uint64) {
	o.Channel = channel
}

// WriteToRequest writes these params to a swagger request
func (o *SetPtzPositionAbsoluteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
