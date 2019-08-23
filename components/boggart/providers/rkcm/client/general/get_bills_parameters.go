// Code generated by go-swagger; DO NOT EDIT.

package general

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

// NewGetBillsParams creates a new GetBillsParams object
// with the default values initialized.
func NewGetBillsParams() *GetBillsParams {
	var ()
	return &GetBillsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBillsParamsWithTimeout creates a new GetBillsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBillsParamsWithTimeout(timeout time.Duration) *GetBillsParams {
	var ()
	return &GetBillsParams{

		timeout: timeout,
	}
}

// NewGetBillsParamsWithContext creates a new GetBillsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBillsParamsWithContext(ctx context.Context) *GetBillsParams {
	var ()
	return &GetBillsParams{

		Context: ctx,
	}
}

// NewGetBillsParamsWithHTTPClient creates a new GetBillsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBillsParamsWithHTTPClient(client *http.Client) *GetBillsParams {
	var ()
	return &GetBillsParams{
		HTTPClient: client,
	}
}

/*GetBillsParams contains all the parameters to send to the API endpoint
for the get bills operation typically these are written to a http.Request
*/
type GetBillsParams struct {

	/*Login
	  Login

	*/
	Login string
	/*Pwd
	  Password

	*/
	Pwd string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bills params
func (o *GetBillsParams) WithTimeout(timeout time.Duration) *GetBillsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bills params
func (o *GetBillsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bills params
func (o *GetBillsParams) WithContext(ctx context.Context) *GetBillsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bills params
func (o *GetBillsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bills params
func (o *GetBillsParams) WithHTTPClient(client *http.Client) *GetBillsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bills params
func (o *GetBillsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogin adds the login to the get bills params
func (o *GetBillsParams) WithLogin(login string) *GetBillsParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the get bills params
func (o *GetBillsParams) SetLogin(login string) {
	o.Login = login
}

// WithPwd adds the pwd to the get bills params
func (o *GetBillsParams) WithPwd(pwd string) *GetBillsParams {
	o.SetPwd(pwd)
	return o
}

// SetPwd adds the pwd to the get bills params
func (o *GetBillsParams) SetPwd(pwd string) {
	o.Pwd = pwd
}

// WriteToRequest writes these params to a swagger request
func (o *GetBillsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param login
	qrLogin := o.Login
	qLogin := qrLogin
	if qLogin != "" {
		if err := r.SetQueryParam("login", qLogin); err != nil {
			return err
		}
	}

	// query param pwd
	qrPwd := o.Pwd
	qPwd := qrPwd
	if qPwd != "" {
		if err := r.SetQueryParam("pwd", qPwd); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}