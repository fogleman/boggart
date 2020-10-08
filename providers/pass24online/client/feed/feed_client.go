// Code generated by go-swagger; DO NOT EDIT.

package feed

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new feed API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for feed API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetFeed(params *GetFeedParams, authInfo runtime.ClientAuthInfoWriter) (*GetFeedOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetFeed get feed API
*/
func (a *Client) GetFeed(params *GetFeedParams, authInfo runtime.ClientAuthInfoWriter) (*GetFeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFeed",
		Method:             "GET",
		PathPattern:        "/feed",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFeedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFeedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFeedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
