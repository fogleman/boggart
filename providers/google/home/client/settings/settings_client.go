// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SetEurekaInfo this can set custom values to some options

Only fields to be modified need to be sent, not all. Sending non-existant fields will still return a 200 OK, but they are not saved.
*/
func (a *Client) SetEurekaInfo(params *SetEurekaInfoParams) (*SetEurekaInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetEurekaInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setEurekaInfo",
		Method:             "POST",
		PathPattern:        "/setup/set_eureka_info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetEurekaInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetEurekaInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setEurekaInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
