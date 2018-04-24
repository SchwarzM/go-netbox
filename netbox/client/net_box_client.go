// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/client/circuits"
	"github.com/digitalocean/go-netbox/netbox/client/dcim"
	"github.com/digitalocean/go-netbox/netbox/client/extras"
	"github.com/digitalocean/go-netbox/netbox/client/ipam"
	"github.com/digitalocean/go-netbox/netbox/client/secrets"
	"github.com/digitalocean/go-netbox/netbox/client/tenancy"
	"github.com/digitalocean/go-netbox/netbox/client/virtualization"
)

// Default net box HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:8000"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new net box HTTP client.
func NewHTTPClient(formats strfmt.Registry) *NetBox {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new net box HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *NetBox {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new net box client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *NetBox {
	cli := new(NetBox)
	cli.Transport = transport

	cli.Circuits = circuits.New(transport, formats)

	cli.Dcim = dcim.New(transport, formats)

	cli.Extras = extras.New(transport, formats)

	cli.IPAM = ipam.New(transport, formats)

	cli.Secrets = secrets.New(transport, formats)

	cli.Tenancy = tenancy.New(transport, formats)

	cli.Virtualization = virtualization.New(transport, formats)

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

// NetBox is a client for net box
type NetBox struct {
	Circuits *circuits.Client

	Dcim *dcim.Client

	Extras *extras.Client

	IPAM *ipam.Client

	Secrets *secrets.Client

	Tenancy *tenancy.Client

	Virtualization *virtualization.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *NetBox) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Circuits.SetTransport(transport)

	c.Dcim.SetTransport(transport)

	c.Extras.SetTransport(transport)

	c.IPAM.SetTransport(transport)

	c.Secrets.SetTransport(transport)

	c.Tenancy.SetTransport(transport)

	c.Virtualization.SetTransport(transport)

}
