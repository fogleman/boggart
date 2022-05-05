// Code generated by go-swagger; DO NOT EDIT.

package streaming

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new streaming API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for streaming API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetStreamingPicture(params *GetStreamingPictureParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*GetStreamingPictureOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetStreamingPicture get streaming picture API
*/
func (a *Client) GetStreamingPicture(params *GetStreamingPictureParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*GetStreamingPictureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStreamingPictureParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStreamingPicture",
		Method:             "GET",
		PathPattern:        "/Streaming/channels/{channel}/picture",
		ProducesMediaTypes: []string{"image/jpeg; charset=\"UTF-8\""},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStreamingPictureReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStreamingPictureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStreamingPicture: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
