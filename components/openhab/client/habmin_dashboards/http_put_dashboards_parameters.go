// Code generated by go-swagger; DO NOT EDIT.

package habmin_dashboards

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

	models "github.com/kihamo/boggart/components/openhab/models"
)

// NewHTTPPutDashboardsParams creates a new HTTPPutDashboardsParams object
// with the default values initialized.
func NewHTTPPutDashboardsParams() *HTTPPutDashboardsParams {
	var ()
	return &HTTPPutDashboardsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHTTPPutDashboardsParamsWithTimeout creates a new HTTPPutDashboardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHTTPPutDashboardsParamsWithTimeout(timeout time.Duration) *HTTPPutDashboardsParams {
	var ()
	return &HTTPPutDashboardsParams{

		timeout: timeout,
	}
}

// NewHTTPPutDashboardsParamsWithContext creates a new HTTPPutDashboardsParams object
// with the default values initialized, and the ability to set a context for a request
func NewHTTPPutDashboardsParamsWithContext(ctx context.Context) *HTTPPutDashboardsParams {
	var ()
	return &HTTPPutDashboardsParams{

		Context: ctx,
	}
}

// NewHTTPPutDashboardsParamsWithHTTPClient creates a new HTTPPutDashboardsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHTTPPutDashboardsParamsWithHTTPClient(client *http.Client) *HTTPPutDashboardsParams {
	var ()
	return &HTTPPutDashboardsParams{
		HTTPClient: client,
	}
}

/*HTTPPutDashboardsParams contains all the parameters to send to the API endpoint
for the http put dashboards operation typically these are written to a http.Request
*/
type HTTPPutDashboardsParams struct {

	/*Body*/
	Body *models.DashboardConfigBean
	/*DashboardID*/
	DashboardID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the http put dashboards params
func (o *HTTPPutDashboardsParams) WithTimeout(timeout time.Duration) *HTTPPutDashboardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the http put dashboards params
func (o *HTTPPutDashboardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the http put dashboards params
func (o *HTTPPutDashboardsParams) WithContext(ctx context.Context) *HTTPPutDashboardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the http put dashboards params
func (o *HTTPPutDashboardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the http put dashboards params
func (o *HTTPPutDashboardsParams) WithHTTPClient(client *http.Client) *HTTPPutDashboardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the http put dashboards params
func (o *HTTPPutDashboardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the http put dashboards params
func (o *HTTPPutDashboardsParams) WithBody(body *models.DashboardConfigBean) *HTTPPutDashboardsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the http put dashboards params
func (o *HTTPPutDashboardsParams) SetBody(body *models.DashboardConfigBean) {
	o.Body = body
}

// WithDashboardID adds the dashboardID to the http put dashboards params
func (o *HTTPPutDashboardsParams) WithDashboardID(dashboardID int32) *HTTPPutDashboardsParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the http put dashboards params
func (o *HTTPPutDashboardsParams) SetDashboardID(dashboardID int32) {
	o.DashboardID = dashboardID
}

// WriteToRequest writes these params to a swagger request
func (o *HTTPPutDashboardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param dashboardId
	if err := r.SetPathParam("dashboardId", swag.FormatInt32(o.DashboardID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}