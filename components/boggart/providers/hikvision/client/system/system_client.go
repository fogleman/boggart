// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetStatus get status API
*/
func (a *Client) GetStatus(params *GetStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStatus",
		Method:             "GET",
		PathPattern:        "/System/status",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStatusOK), nil

}

/*
GetSystemDeviceInfo get system device info API
*/
func (a *Client) GetSystemDeviceInfo(params *GetSystemDeviceInfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetSystemDeviceInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemDeviceInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemDeviceInfo",
		Method:             "GET",
		PathPattern:        "/System/deviceInfo",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSystemDeviceInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSystemDeviceInfoOK), nil

}

/*
GetSystemUpgradeStatus get system upgrade status API
*/
func (a *Client) GetSystemUpgradeStatus(params *GetSystemUpgradeStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetSystemUpgradeStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemUpgradeStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemUpgradeStatus",
		Method:             "GET",
		PathPattern:        "/System/upgradeStatus",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSystemUpgradeStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSystemUpgradeStatusOK), nil

}

/*
Reboot reboot API
*/
func (a *Client) Reboot(params *RebootParams, authInfo runtime.ClientAuthInfoWriter) (*RebootOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebootParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reboot",
		Method:             "PUT",
		PathPattern:        "/System/reboot",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RebootReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RebootOK), nil

}

/*
UpdateSystemFirmware updates the firmware of the device
*/
func (a *Client) UpdateSystemFirmware(params *UpdateSystemFirmwareParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSystemFirmwareOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSystemFirmwareParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSystemFirmware",
		Method:             "PUT",
		PathPattern:        "/System/updateFirmware",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSystemFirmwareReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSystemFirmwareOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}