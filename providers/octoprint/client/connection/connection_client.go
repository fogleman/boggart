// Code generated by go-swagger; DO NOT EDIT.

package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new connection API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for connection API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetConnection gets connection settings
*/
func (a *Client) GetConnection(params *GetConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConnectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConnection",
		Method:             "GET",
		PathPattern:        "/connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SendConnectionCommand issues a connection command
*/
func (a *Client) SendConnectionCommand(params *SendConnectionCommandParams, authInfo runtime.ClientAuthInfoWriter) (*SendConnectionCommandNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendConnectionCommandParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendConnectionCommand",
		Method:             "POST",
		PathPattern:        "/connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SendConnectionCommandReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SendConnectionCommandNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendConnectionCommand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
