// Code generated by go-swagger; DO NOT EDIT.

package streaming

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

// NewGetStreamingPictureParams creates a new GetStreamingPictureParams object
// with the default values initialized.
func NewGetStreamingPictureParams() *GetStreamingPictureParams {
	var ()
	return &GetStreamingPictureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStreamingPictureParamsWithTimeout creates a new GetStreamingPictureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStreamingPictureParamsWithTimeout(timeout time.Duration) *GetStreamingPictureParams {
	var ()
	return &GetStreamingPictureParams{

		timeout: timeout,
	}
}

// NewGetStreamingPictureParamsWithContext creates a new GetStreamingPictureParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStreamingPictureParamsWithContext(ctx context.Context) *GetStreamingPictureParams {
	var ()
	return &GetStreamingPictureParams{

		Context: ctx,
	}
}

// NewGetStreamingPictureParamsWithHTTPClient creates a new GetStreamingPictureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStreamingPictureParamsWithHTTPClient(client *http.Client) *GetStreamingPictureParams {
	var ()
	return &GetStreamingPictureParams{
		HTTPClient: client,
	}
}

/*GetStreamingPictureParams contains all the parameters to send to the API endpoint
for the get streaming picture operation typically these are written to a http.Request
*/
type GetStreamingPictureParams struct {

	/*Channel
	  Channel ID

	*/
	Channel uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get streaming picture params
func (o *GetStreamingPictureParams) WithTimeout(timeout time.Duration) *GetStreamingPictureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get streaming picture params
func (o *GetStreamingPictureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get streaming picture params
func (o *GetStreamingPictureParams) WithContext(ctx context.Context) *GetStreamingPictureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get streaming picture params
func (o *GetStreamingPictureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get streaming picture params
func (o *GetStreamingPictureParams) WithHTTPClient(client *http.Client) *GetStreamingPictureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get streaming picture params
func (o *GetStreamingPictureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannel adds the channel to the get streaming picture params
func (o *GetStreamingPictureParams) WithChannel(channel uint64) *GetStreamingPictureParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the get streaming picture params
func (o *GetStreamingPictureParams) SetChannel(channel uint64) {
	o.Channel = channel
}

// WriteToRequest writes these params to a swagger request
func (o *GetStreamingPictureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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