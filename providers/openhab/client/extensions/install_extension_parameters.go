// Code generated by go-swagger; DO NOT EDIT.

package extensions

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

// NewInstallExtensionParams creates a new InstallExtensionParams object
// with the default values initialized.
func NewInstallExtensionParams() *InstallExtensionParams {
	var ()
	return &InstallExtensionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstallExtensionParamsWithTimeout creates a new InstallExtensionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstallExtensionParamsWithTimeout(timeout time.Duration) *InstallExtensionParams {
	var ()
	return &InstallExtensionParams{

		timeout: timeout,
	}
}

// NewInstallExtensionParamsWithContext creates a new InstallExtensionParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstallExtensionParamsWithContext(ctx context.Context) *InstallExtensionParams {
	var ()
	return &InstallExtensionParams{

		Context: ctx,
	}
}

// NewInstallExtensionParamsWithHTTPClient creates a new InstallExtensionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstallExtensionParamsWithHTTPClient(client *http.Client) *InstallExtensionParams {
	var ()
	return &InstallExtensionParams{
		HTTPClient: client,
	}
}

/*InstallExtensionParams contains all the parameters to send to the API endpoint
for the install extension operation typically these are written to a http.Request
*/
type InstallExtensionParams struct {

	/*ExtensionID
	  extension ID

	*/
	ExtensionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the install extension params
func (o *InstallExtensionParams) WithTimeout(timeout time.Duration) *InstallExtensionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the install extension params
func (o *InstallExtensionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the install extension params
func (o *InstallExtensionParams) WithContext(ctx context.Context) *InstallExtensionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the install extension params
func (o *InstallExtensionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the install extension params
func (o *InstallExtensionParams) WithHTTPClient(client *http.Client) *InstallExtensionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the install extension params
func (o *InstallExtensionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtensionID adds the extensionID to the install extension params
func (o *InstallExtensionParams) WithExtensionID(extensionID string) *InstallExtensionParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the install extension params
func (o *InstallExtensionParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *InstallExtensionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extensionId
	if err := r.SetPathParam("extensionId", o.ExtensionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}