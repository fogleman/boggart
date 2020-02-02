// Code generated by go-swagger; DO NOT EDIT.

package items

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

// NewAddMemberParams creates a new AddMemberParams object
// with the default values initialized.
func NewAddMemberParams() *AddMemberParams {
	var ()
	return &AddMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddMemberParamsWithTimeout creates a new AddMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddMemberParamsWithTimeout(timeout time.Duration) *AddMemberParams {
	var ()
	return &AddMemberParams{

		timeout: timeout,
	}
}

// NewAddMemberParamsWithContext creates a new AddMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddMemberParamsWithContext(ctx context.Context) *AddMemberParams {
	var ()
	return &AddMemberParams{

		Context: ctx,
	}
}

// NewAddMemberParamsWithHTTPClient creates a new AddMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddMemberParamsWithHTTPClient(client *http.Client) *AddMemberParams {
	var ()
	return &AddMemberParams{
		HTTPClient: client,
	}
}

/*AddMemberParams contains all the parameters to send to the API endpoint
for the add member operation typically these are written to a http.Request
*/
type AddMemberParams struct {

	/*ItemName
	  item name

	*/
	ItemName string
	/*MemberItemName
	  member item name

	*/
	MemberItemName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add member params
func (o *AddMemberParams) WithTimeout(timeout time.Duration) *AddMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add member params
func (o *AddMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add member params
func (o *AddMemberParams) WithContext(ctx context.Context) *AddMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add member params
func (o *AddMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add member params
func (o *AddMemberParams) WithHTTPClient(client *http.Client) *AddMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add member params
func (o *AddMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemName adds the itemName to the add member params
func (o *AddMemberParams) WithItemName(itemName string) *AddMemberParams {
	o.SetItemName(itemName)
	return o
}

// SetItemName adds the itemName to the add member params
func (o *AddMemberParams) SetItemName(itemName string) {
	o.ItemName = itemName
}

// WithMemberItemName adds the memberItemName to the add member params
func (o *AddMemberParams) WithMemberItemName(memberItemName string) *AddMemberParams {
	o.SetMemberItemName(memberItemName)
	return o
}

// SetMemberItemName adds the memberItemName to the add member params
func (o *AddMemberParams) SetMemberItemName(memberItemName string) {
	o.MemberItemName = memberItemName
}

// WriteToRequest writes these params to a swagger request
func (o *AddMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param itemName
	if err := r.SetPathParam("itemName", o.ItemName); err != nil {
		return err
	}

	// path param memberItemName
	if err := r.SetPathParam("memberItemName", o.MemberItemName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}