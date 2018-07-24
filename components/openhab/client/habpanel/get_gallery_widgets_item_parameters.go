// Code generated by go-swagger; DO NOT EDIT.

package habpanel

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

// NewGetGalleryWidgetsItemParams creates a new GetGalleryWidgetsItemParams object
// with the default values initialized.
func NewGetGalleryWidgetsItemParams() *GetGalleryWidgetsItemParams {
	var ()
	return &GetGalleryWidgetsItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGalleryWidgetsItemParamsWithTimeout creates a new GetGalleryWidgetsItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGalleryWidgetsItemParamsWithTimeout(timeout time.Duration) *GetGalleryWidgetsItemParams {
	var ()
	return &GetGalleryWidgetsItemParams{

		timeout: timeout,
	}
}

// NewGetGalleryWidgetsItemParamsWithContext creates a new GetGalleryWidgetsItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGalleryWidgetsItemParamsWithContext(ctx context.Context) *GetGalleryWidgetsItemParams {
	var ()
	return &GetGalleryWidgetsItemParams{

		Context: ctx,
	}
}

// NewGetGalleryWidgetsItemParamsWithHTTPClient creates a new GetGalleryWidgetsItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGalleryWidgetsItemParamsWithHTTPClient(client *http.Client) *GetGalleryWidgetsItemParams {
	var ()
	return &GetGalleryWidgetsItemParams{
		HTTPClient: client,
	}
}

/*GetGalleryWidgetsItemParams contains all the parameters to send to the API endpoint
for the get gallery widgets item operation typically these are written to a http.Request
*/
type GetGalleryWidgetsItemParams struct {

	/*GalleryName
	  gallery name e.g. 'community'

	*/
	GalleryName string
	/*ID
	  id within the gallery

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gallery widgets item params
func (o *GetGalleryWidgetsItemParams) WithTimeout(timeout time.Duration) *GetGalleryWidgetsItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gallery widgets item params
func (o *GetGalleryWidgetsItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gallery widgets item params
func (o *GetGalleryWidgetsItemParams) WithContext(ctx context.Context) *GetGalleryWidgetsItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gallery widgets item params
func (o *GetGalleryWidgetsItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gallery widgets item params
func (o *GetGalleryWidgetsItemParams) WithHTTPClient(client *http.Client) *GetGalleryWidgetsItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gallery widgets item params
func (o *GetGalleryWidgetsItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGalleryName adds the galleryName to the get gallery widgets item params
func (o *GetGalleryWidgetsItemParams) WithGalleryName(galleryName string) *GetGalleryWidgetsItemParams {
	o.SetGalleryName(galleryName)
	return o
}

// SetGalleryName adds the galleryName to the get gallery widgets item params
func (o *GetGalleryWidgetsItemParams) SetGalleryName(galleryName string) {
	o.GalleryName = galleryName
}

// WithID adds the id to the get gallery widgets item params
func (o *GetGalleryWidgetsItemParams) WithID(id string) *GetGalleryWidgetsItemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get gallery widgets item params
func (o *GetGalleryWidgetsItemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetGalleryWidgetsItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param galleryName
	if err := r.SetPathParam("galleryName", o.GalleryName); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
