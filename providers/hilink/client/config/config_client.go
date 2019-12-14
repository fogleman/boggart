// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetGlobalConfig get global config API
*/
func (a *Client) GetGlobalConfig(params *GetGlobalConfigParams) (*GetGlobalConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGlobalConfig",
		Method:             "GET",
		PathPattern:        "/config/global/config.xml",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGlobalConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGlobalConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetGlobalConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}