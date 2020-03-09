// Code generated by go-swagger; DO NOT EDIT.

package information

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

	models "github.com/kihamo/boggart/providers/tvt/models"
)

// NewGetBasicConfigParams creates a new GetBasicConfigParams object
// with the default values initialized.
func NewGetBasicConfigParams() *GetBasicConfigParams {
	var ()
	return &GetBasicConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBasicConfigParamsWithTimeout creates a new GetBasicConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBasicConfigParamsWithTimeout(timeout time.Duration) *GetBasicConfigParams {
	var ()
	return &GetBasicConfigParams{

		timeout: timeout,
	}
}

// NewGetBasicConfigParamsWithContext creates a new GetBasicConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBasicConfigParamsWithContext(ctx context.Context) *GetBasicConfigParams {
	var ()
	return &GetBasicConfigParams{

		Context: ctx,
	}
}

// NewGetBasicConfigParamsWithHTTPClient creates a new GetBasicConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBasicConfigParamsWithHTTPClient(client *http.Client) *GetBasicConfigParams {
	var ()
	return &GetBasicConfigParams{
		HTTPClient: client,
	}
}

/*GetBasicConfigParams contains all the parameters to send to the API endpoint
for the get basic config operation typically these are written to a http.Request
*/
type GetBasicConfigParams struct {

	/*Request*/
	Request models.Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get basic config params
func (o *GetBasicConfigParams) WithTimeout(timeout time.Duration) *GetBasicConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get basic config params
func (o *GetBasicConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get basic config params
func (o *GetBasicConfigParams) WithContext(ctx context.Context) *GetBasicConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get basic config params
func (o *GetBasicConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get basic config params
func (o *GetBasicConfigParams) WithHTTPClient(client *http.Client) *GetBasicConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get basic config params
func (o *GetBasicConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the get basic config params
func (o *GetBasicConfigParams) WithRequest(request models.Request) *GetBasicConfigParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the get basic config params
func (o *GetBasicConfigParams) SetRequest(request models.Request) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *GetBasicConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
