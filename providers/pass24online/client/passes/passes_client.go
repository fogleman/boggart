// Code generated by go-swagger; DO NOT EDIT.

package passes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new passes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for passes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetPassByID(params *GetPassByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPassByIDOK, error)

	GetPasses(params *GetPassesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPassesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetPassByID get pass by ID API
*/
func (a *Client) GetPassByID(params *GetPassByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPassByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPassByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPassByID",
		Method:             "GET",
		PathPattern:        "/passes/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPassByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPassByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPassByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPasses get passes API
*/
func (a *Client) GetPasses(params *GetPassesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPassesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPassesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPasses",
		Method:             "GET",
		PathPattern:        "/passes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPassesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPassesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPassesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}