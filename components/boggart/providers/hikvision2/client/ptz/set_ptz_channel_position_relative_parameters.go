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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/components/boggart/providers/hikvision2/models"
)

// NewSetPtzChannelPositionRelativeParams creates a new SetPtzChannelPositionRelativeParams object
// with the default values initialized.
func NewSetPtzChannelPositionRelativeParams() *SetPtzChannelPositionRelativeParams {
	var ()
	return &SetPtzChannelPositionRelativeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetPtzChannelPositionRelativeParamsWithTimeout creates a new SetPtzChannelPositionRelativeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetPtzChannelPositionRelativeParamsWithTimeout(timeout time.Duration) *SetPtzChannelPositionRelativeParams {
	var ()
	return &SetPtzChannelPositionRelativeParams{

		timeout: timeout,
	}
}

// NewSetPtzChannelPositionRelativeParamsWithContext creates a new SetPtzChannelPositionRelativeParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetPtzChannelPositionRelativeParamsWithContext(ctx context.Context) *SetPtzChannelPositionRelativeParams {
	var ()
	return &SetPtzChannelPositionRelativeParams{

		Context: ctx,
	}
}

// NewSetPtzChannelPositionRelativeParamsWithHTTPClient creates a new SetPtzChannelPositionRelativeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetPtzChannelPositionRelativeParamsWithHTTPClient(client *http.Client) *SetPtzChannelPositionRelativeParams {
	var ()
	return &SetPtzChannelPositionRelativeParams{
		HTTPClient: client,
	}
}

/*SetPtzChannelPositionRelativeParams contains all the parameters to send to the API endpoint
for the set ptz channel position relative operation typically these are written to a http.Request
*/
type SetPtzChannelPositionRelativeParams struct {

	/*PTZData*/
	PTZData *models.PTZData
	/*Channel
	  Channel ID

	*/
	Channel uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set ptz channel position relative params
func (o *SetPtzChannelPositionRelativeParams) WithTimeout(timeout time.Duration) *SetPtzChannelPositionRelativeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set ptz channel position relative params
func (o *SetPtzChannelPositionRelativeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set ptz channel position relative params
func (o *SetPtzChannelPositionRelativeParams) WithContext(ctx context.Context) *SetPtzChannelPositionRelativeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set ptz channel position relative params
func (o *SetPtzChannelPositionRelativeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set ptz channel position relative params
func (o *SetPtzChannelPositionRelativeParams) WithHTTPClient(client *http.Client) *SetPtzChannelPositionRelativeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set ptz channel position relative params
func (o *SetPtzChannelPositionRelativeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPTZData adds the pTZData to the set ptz channel position relative params
func (o *SetPtzChannelPositionRelativeParams) WithPTZData(pTZData *models.PTZData) *SetPtzChannelPositionRelativeParams {
	o.SetPTZData(pTZData)
	return o
}

// SetPTZData adds the pTZData to the set ptz channel position relative params
func (o *SetPtzChannelPositionRelativeParams) SetPTZData(pTZData *models.PTZData) {
	o.PTZData = pTZData
}

// WithChannel adds the channel to the set ptz channel position relative params
func (o *SetPtzChannelPositionRelativeParams) WithChannel(channel uint64) *SetPtzChannelPositionRelativeParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the set ptz channel position relative params
func (o *SetPtzChannelPositionRelativeParams) SetChannel(channel uint64) {
	o.Channel = channel
}

// WriteToRequest writes these params to a swagger request
func (o *SetPtzChannelPositionRelativeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
