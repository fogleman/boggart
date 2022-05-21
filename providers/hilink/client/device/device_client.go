// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new device API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for device API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeviceControl(params *DeviceControlParams, opts ...ClientOption) (*DeviceControlOK, error)

	GetCompressLogFile(params *GetCompressLogFileParams, opts ...ClientOption) (*GetCompressLogFileOK, error)

	GetDeviceAutoRunVersion(params *GetDeviceAutoRunVersionParams, opts ...ClientOption) (*GetDeviceAutoRunVersionOK, error)

	GetDeviceBasicInformation(params *GetDeviceBasicInformationParams, opts ...ClientOption) (*GetDeviceBasicInformationOK, error)

	GetDeviceInformation(params *GetDeviceInformationParams, opts ...ClientOption) (*GetDeviceInformationOK, error)

	GetDeviceSignal(params *GetDeviceSignalParams, opts ...ClientOption) (*GetDeviceSignalOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeviceControl device control API
*/
func (a *Client) DeviceControl(params *DeviceControlParams, opts ...ClientOption) (*DeviceControlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeviceControlParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeviceControl",
		Method:             "POST",
		PathPattern:        "/api/device/control",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeviceControlReader{formats: a.formats},
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
	success, ok := result.(*DeviceControlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeviceControlDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCompressLogFile get compress log file API
*/
func (a *Client) GetCompressLogFile(params *GetCompressLogFileParams, opts ...ClientOption) (*GetCompressLogFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCompressLogFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCompressLogFile",
		Method:             "GET",
		PathPattern:        "/api/device/compresslogfile",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCompressLogFileReader{formats: a.formats},
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
	success, ok := result.(*GetCompressLogFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCompressLogFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetDeviceAutoRunVersion get device auto run version API
*/
func (a *Client) GetDeviceAutoRunVersion(params *GetDeviceAutoRunVersionParams, opts ...ClientOption) (*GetDeviceAutoRunVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceAutoRunVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceAutoRunVersion",
		Method:             "GET",
		PathPattern:        "/api/device/autorun-version",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeviceAutoRunVersionReader{formats: a.formats},
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
	success, ok := result.(*GetDeviceAutoRunVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDeviceAutoRunVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetDeviceBasicInformation get device basic information API
*/
func (a *Client) GetDeviceBasicInformation(params *GetDeviceBasicInformationParams, opts ...ClientOption) (*GetDeviceBasicInformationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceBasicInformationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceBasicInformation",
		Method:             "GET",
		PathPattern:        "/api/device/basic_information",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeviceBasicInformationReader{formats: a.formats},
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
	success, ok := result.(*GetDeviceBasicInformationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDeviceBasicInformationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetDeviceInformation get device information API
*/
func (a *Client) GetDeviceInformation(params *GetDeviceInformationParams, opts ...ClientOption) (*GetDeviceInformationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceInformationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceInformation",
		Method:             "GET",
		PathPattern:        "/api/device/information",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeviceInformationReader{formats: a.formats},
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
	success, ok := result.(*GetDeviceInformationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDeviceInformationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetDeviceSignal get device signal API
*/
func (a *Client) GetDeviceSignal(params *GetDeviceSignalParams, opts ...ClientOption) (*GetDeviceSignalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceSignalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceSignal",
		Method:             "GET",
		PathPattern:        "/api/device/signal",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeviceSignalReader{formats: a.formats},
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
	success, ok := result.(*GetDeviceSignalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDeviceSignalDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
