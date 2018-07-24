// Code generated by go-swagger; DO NOT EDIT.

package habmin_charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewHTTPDeleteChartsParams creates a new HTTPDeleteChartsParams object
// with the default values initialized.
func NewHTTPDeleteChartsParams() *HTTPDeleteChartsParams {
	var ()
	return &HTTPDeleteChartsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHTTPDeleteChartsParamsWithTimeout creates a new HTTPDeleteChartsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHTTPDeleteChartsParamsWithTimeout(timeout time.Duration) *HTTPDeleteChartsParams {
	var ()
	return &HTTPDeleteChartsParams{

		timeout: timeout,
	}
}

// NewHTTPDeleteChartsParamsWithContext creates a new HTTPDeleteChartsParams object
// with the default values initialized, and the ability to set a context for a request
func NewHTTPDeleteChartsParamsWithContext(ctx context.Context) *HTTPDeleteChartsParams {
	var ()
	return &HTTPDeleteChartsParams{

		Context: ctx,
	}
}

// NewHTTPDeleteChartsParamsWithHTTPClient creates a new HTTPDeleteChartsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHTTPDeleteChartsParamsWithHTTPClient(client *http.Client) *HTTPDeleteChartsParams {
	var ()
	return &HTTPDeleteChartsParams{
		HTTPClient: client,
	}
}

/*HTTPDeleteChartsParams contains all the parameters to send to the API endpoint
for the http delete charts operation typically these are written to a http.Request
*/
type HTTPDeleteChartsParams struct {

	/*ChartID*/
	ChartID int32
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the http delete charts params
func (o *HTTPDeleteChartsParams) WithTimeout(timeout time.Duration) *HTTPDeleteChartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the http delete charts params
func (o *HTTPDeleteChartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the http delete charts params
func (o *HTTPDeleteChartsParams) WithContext(ctx context.Context) *HTTPDeleteChartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the http delete charts params
func (o *HTTPDeleteChartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the http delete charts params
func (o *HTTPDeleteChartsParams) WithHTTPClient(client *http.Client) *HTTPDeleteChartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the http delete charts params
func (o *HTTPDeleteChartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChartID adds the chartID to the http delete charts params
func (o *HTTPDeleteChartsParams) WithChartID(chartID int32) *HTTPDeleteChartsParams {
	o.SetChartID(chartID)
	return o
}

// SetChartID adds the chartId to the http delete charts params
func (o *HTTPDeleteChartsParams) SetChartID(chartID int32) {
	o.ChartID = chartID
}

// WithType adds the typeVar to the http delete charts params
func (o *HTTPDeleteChartsParams) WithType(typeVar *string) *HTTPDeleteChartsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the http delete charts params
func (o *HTTPDeleteChartsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *HTTPDeleteChartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param chartId
	if err := r.SetPathParam("chartId", swag.FormatInt32(o.ChartID)); err != nil {
		return err
	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
