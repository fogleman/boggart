// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/octoprint/client/authorization"
	"github.com/kihamo/boggart/providers/octoprint/client/connection"
	"github.com/kihamo/boggart/providers/octoprint/client/languages"
	"github.com/kihamo/boggart/providers/octoprint/client/printer"
	"github.com/kihamo/boggart/providers/octoprint/client/system"
	"github.com/kihamo/boggart/providers/octoprint/client/version"
)

// Default octo print HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new octo print HTTP client.
func NewHTTPClient(formats strfmt.Registry) *OctoPrint {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new octo print HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *OctoPrint {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new octo print client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *OctoPrint {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(OctoPrint)
	cli.Transport = transport

	cli.Authorization = authorization.New(transport, formats)

	cli.Connection = connection.New(transport, formats)

	cli.Languages = languages.New(transport, formats)

	cli.Printer = printer.New(transport, formats)

	cli.System = system.New(transport, formats)

	cli.Version = version.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// OctoPrint is a client for octo print
type OctoPrint struct {
	Authorization *authorization.Client

	Connection *connection.Client

	Languages *languages.Client

	Printer *printer.Client

	System *system.Client

	Version *version.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *OctoPrint) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Authorization.SetTransport(transport)

	c.Connection.SetTransport(transport)

	c.Languages.SetTransport(transport)

	c.Printer.SetTransport(transport)

	c.System.SetTransport(transport)

	c.Version.SetTransport(transport)

}
