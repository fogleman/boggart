// Code generated by go-swagger; DO NOT EDIT.

package links

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

// NewGetLinkParams creates a new GetLinkParams object
// with the default values initialized.
func NewGetLinkParams() *GetLinkParams {
	var ()
	return &GetLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLinkParamsWithTimeout creates a new GetLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLinkParamsWithTimeout(timeout time.Duration) *GetLinkParams {
	var ()
	return &GetLinkParams{

		timeout: timeout,
	}
}

// NewGetLinkParamsWithContext creates a new GetLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLinkParamsWithContext(ctx context.Context) *GetLinkParams {
	var ()
	return &GetLinkParams{

		Context: ctx,
	}
}

// NewGetLinkParamsWithHTTPClient creates a new GetLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLinkParamsWithHTTPClient(client *http.Client) *GetLinkParams {
	var ()
	return &GetLinkParams{
		HTTPClient: client,
	}
}

/*GetLinkParams contains all the parameters to send to the API endpoint
for the get link operation typically these are written to a http.Request
*/
type GetLinkParams struct {

	/*ChannelUID
	  channelUID

	*/
	ChannelUID string
	/*ItemName
	  itemName

	*/
	ItemName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get link params
func (o *GetLinkParams) WithTimeout(timeout time.Duration) *GetLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get link params
func (o *GetLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get link params
func (o *GetLinkParams) WithContext(ctx context.Context) *GetLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get link params
func (o *GetLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get link params
func (o *GetLinkParams) WithHTTPClient(client *http.Client) *GetLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get link params
func (o *GetLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelUID adds the channelUID to the get link params
func (o *GetLinkParams) WithChannelUID(channelUID string) *GetLinkParams {
	o.SetChannelUID(channelUID)
	return o
}

// SetChannelUID adds the channelUid to the get link params
func (o *GetLinkParams) SetChannelUID(channelUID string) {
	o.ChannelUID = channelUID
}

// WithItemName adds the itemName to the get link params
func (o *GetLinkParams) WithItemName(itemName string) *GetLinkParams {
	o.SetItemName(itemName)
	return o
}

// SetItemName adds the itemName to the get link params
func (o *GetLinkParams) SetItemName(itemName string) {
	o.ItemName = itemName
}

// WriteToRequest writes these params to a swagger request
func (o *GetLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channelUID
	if err := r.SetPathParam("channelUID", o.ChannelUID); err != nil {
		return err
	}

	// path param itemName
	if err := r.SetPathParam("itemName", o.ItemName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
