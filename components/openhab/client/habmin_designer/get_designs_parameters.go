// Code generated by go-swagger; DO NOT EDIT.

package habmin_designer

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

// NewGetDesignsParams creates a new GetDesignsParams object
// with the default values initialized.
func NewGetDesignsParams() *GetDesignsParams {

	return &GetDesignsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDesignsParamsWithTimeout creates a new GetDesignsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDesignsParamsWithTimeout(timeout time.Duration) *GetDesignsParams {

	return &GetDesignsParams{

		timeout: timeout,
	}
}

// NewGetDesignsParamsWithContext creates a new GetDesignsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDesignsParamsWithContext(ctx context.Context) *GetDesignsParams {

	return &GetDesignsParams{

		Context: ctx,
	}
}

// NewGetDesignsParamsWithHTTPClient creates a new GetDesignsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDesignsParamsWithHTTPClient(client *http.Client) *GetDesignsParams {

	return &GetDesignsParams{
		HTTPClient: client,
	}
}

/*GetDesignsParams contains all the parameters to send to the API endpoint
for the get designs operation typically these are written to a http.Request
*/
type GetDesignsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get designs params
func (o *GetDesignsParams) WithTimeout(timeout time.Duration) *GetDesignsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get designs params
func (o *GetDesignsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get designs params
func (o *GetDesignsParams) WithContext(ctx context.Context) *GetDesignsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get designs params
func (o *GetDesignsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get designs params
func (o *GetDesignsParams) WithHTTPClient(client *http.Client) *GetDesignsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get designs params
func (o *GetDesignsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDesignsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}