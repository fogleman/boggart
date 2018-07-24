// Code generated by go-swagger; DO NOT EDIT.

package habpanel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new habpanel API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for habpanel API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetGalleryWidgetList gets the list of widget gallery items
*/
func (a *Client) GetGalleryWidgetList(params *GetGalleryWidgetListParams) (*GetGalleryWidgetListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGalleryWidgetListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGalleryWidgetList",
		Method:             "GET",
		PathPattern:        "/habpanel/gallery/{galleryName}/widgets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGalleryWidgetListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGalleryWidgetListOK), nil

}

/*
GetGalleryWidgetsItem gets the details about a widget gallery item
*/
func (a *Client) GetGalleryWidgetsItem(params *GetGalleryWidgetsItemParams) (*GetGalleryWidgetsItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGalleryWidgetsItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGalleryWidgetsItem",
		Method:             "GET",
		PathPattern:        "/habpanel/gallery/{galleryName}/widgets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGalleryWidgetsItemReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGalleryWidgetsItemOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
