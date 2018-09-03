// Code generated by go-swagger; DO NOT EDIT.

package config_descriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new config descriptions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for config descriptions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetByURI gets a config description by URI
*/
func (a *Client) GetByURI(params *GetByURIParams) (*GetByURIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetByURIParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getByURI",
		Method:             "GET",
		PathPattern:        "/config-descriptions/{uri}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetByURIReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetByURIOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}