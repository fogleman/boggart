// Code generated by go-swagger; DO NOT EDIT.

package items

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

// NewPostItemCommandParams creates a new PostItemCommandParams object
// with the default values initialized.
func NewPostItemCommandParams() *PostItemCommandParams {
	var ()
	return &PostItemCommandParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostItemCommandParamsWithTimeout creates a new PostItemCommandParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostItemCommandParamsWithTimeout(timeout time.Duration) *PostItemCommandParams {
	var ()
	return &PostItemCommandParams{

		timeout: timeout,
	}
}

// NewPostItemCommandParamsWithContext creates a new PostItemCommandParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostItemCommandParamsWithContext(ctx context.Context) *PostItemCommandParams {
	var ()
	return &PostItemCommandParams{

		Context: ctx,
	}
}

// NewPostItemCommandParamsWithHTTPClient creates a new PostItemCommandParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostItemCommandParamsWithHTTPClient(client *http.Client) *PostItemCommandParams {
	var ()
	return &PostItemCommandParams{
		HTTPClient: client,
	}
}

/*PostItemCommandParams contains all the parameters to send to the API endpoint
for the post item command operation typically these are written to a http.Request
*/
type PostItemCommandParams struct {

	/*Body
	  valid item command (e.g. ON, OFF, UP, DOWN, REFRESH)

	*/
	Body string
	/*Itemname
	  item name

	*/
	Itemname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post item command params
func (o *PostItemCommandParams) WithTimeout(timeout time.Duration) *PostItemCommandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post item command params
func (o *PostItemCommandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post item command params
func (o *PostItemCommandParams) WithContext(ctx context.Context) *PostItemCommandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post item command params
func (o *PostItemCommandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post item command params
func (o *PostItemCommandParams) WithHTTPClient(client *http.Client) *PostItemCommandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post item command params
func (o *PostItemCommandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post item command params
func (o *PostItemCommandParams) WithBody(body string) *PostItemCommandParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post item command params
func (o *PostItemCommandParams) SetBody(body string) {
	o.Body = body
}

// WithItemname adds the itemname to the post item command params
func (o *PostItemCommandParams) WithItemname(itemname string) *PostItemCommandParams {
	o.SetItemname(itemname)
	return o
}

// SetItemname adds the itemname to the post item command params
func (o *PostItemCommandParams) SetItemname(itemname string) {
	o.Itemname = itemname
}

// WriteToRequest writes these params to a swagger request
func (o *PostItemCommandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param itemname
	if err := r.SetPathParam("itemname", o.Itemname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
