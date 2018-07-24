// Code generated by go-swagger; DO NOT EDIT.

package persistence

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

// NewHTTPPutPersistenceItemDataParams creates a new HTTPPutPersistenceItemDataParams object
// with the default values initialized.
func NewHTTPPutPersistenceItemDataParams() *HTTPPutPersistenceItemDataParams {
	var ()
	return &HTTPPutPersistenceItemDataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHTTPPutPersistenceItemDataParamsWithTimeout creates a new HTTPPutPersistenceItemDataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHTTPPutPersistenceItemDataParamsWithTimeout(timeout time.Duration) *HTTPPutPersistenceItemDataParams {
	var ()
	return &HTTPPutPersistenceItemDataParams{

		timeout: timeout,
	}
}

// NewHTTPPutPersistenceItemDataParamsWithContext creates a new HTTPPutPersistenceItemDataParams object
// with the default values initialized, and the ability to set a context for a request
func NewHTTPPutPersistenceItemDataParamsWithContext(ctx context.Context) *HTTPPutPersistenceItemDataParams {
	var ()
	return &HTTPPutPersistenceItemDataParams{

		Context: ctx,
	}
}

// NewHTTPPutPersistenceItemDataParamsWithHTTPClient creates a new HTTPPutPersistenceItemDataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHTTPPutPersistenceItemDataParamsWithHTTPClient(client *http.Client) *HTTPPutPersistenceItemDataParams {
	var ()
	return &HTTPPutPersistenceItemDataParams{
		HTTPClient: client,
	}
}

/*HTTPPutPersistenceItemDataParams contains all the parameters to send to the API endpoint
for the http put persistence item data operation typically these are written to a http.Request
*/
type HTTPPutPersistenceItemDataParams struct {

	/*Itemname
	  The item name.

	*/
	Itemname string
	/*ServiceID
	  Id of the persistence service. If not provided the default service will be used

	*/
	ServiceID *string
	/*State
	  The state to store.

	*/
	State string
	/*Time
	  Time of the data to be stored. Will default to current time. [yyyy-MM-dd'T'HH:mm:ss.SSSZ]

	*/
	Time string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) WithTimeout(timeout time.Duration) *HTTPPutPersistenceItemDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) WithContext(ctx context.Context) *HTTPPutPersistenceItemDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) WithHTTPClient(client *http.Client) *HTTPPutPersistenceItemDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemname adds the itemname to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) WithItemname(itemname string) *HTTPPutPersistenceItemDataParams {
	o.SetItemname(itemname)
	return o
}

// SetItemname adds the itemname to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) SetItemname(itemname string) {
	o.Itemname = itemname
}

// WithServiceID adds the serviceID to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) WithServiceID(serviceID *string) *HTTPPutPersistenceItemDataParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) SetServiceID(serviceID *string) {
	o.ServiceID = serviceID
}

// WithState adds the state to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) WithState(state string) *HTTPPutPersistenceItemDataParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) SetState(state string) {
	o.State = state
}

// WithTime adds the time to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) WithTime(time string) *HTTPPutPersistenceItemDataParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the http put persistence item data params
func (o *HTTPPutPersistenceItemDataParams) SetTime(time string) {
	o.Time = time
}

// WriteToRequest writes these params to a swagger request
func (o *HTTPPutPersistenceItemDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param itemname
	if err := r.SetPathParam("itemname", o.Itemname); err != nil {
		return err
	}

	if o.ServiceID != nil {

		// query param serviceId
		var qrServiceID string
		if o.ServiceID != nil {
			qrServiceID = *o.ServiceID
		}
		qServiceID := qrServiceID
		if qServiceID != "" {
			if err := r.SetQueryParam("serviceId", qServiceID); err != nil {
				return err
			}
		}

	}

	// query param state
	qrState := o.State
	qState := qrState
	if qState != "" {
		if err := r.SetQueryParam("state", qState); err != nil {
			return err
		}
	}

	// query param time
	qrTime := o.Time
	qTime := qrTime
	if qTime != "" {
		if err := r.SetQueryParam("time", qTime); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
