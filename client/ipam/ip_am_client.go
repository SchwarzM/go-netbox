// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ipam API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ipam API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
IPAMChoicesList ipam choices list API
*/
func (a *Client) IPAMChoicesList(params *IPAMChoicesListParams) (*IPAMChoicesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMChoicesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam__choices_list",
		Method:             "GET",
		PathPattern:        "/api/ipam/_choices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMChoicesListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMChoicesListOK), nil

}

/*
IPAMChoicesRead ipam choices read API
*/
func (a *Client) IPAMChoicesRead(params *IPAMChoicesReadParams) (*IPAMChoicesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMChoicesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam__choices_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/_choices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMChoicesReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMChoicesReadOK), nil

}

/*
IPAMAggregatesList ipam aggregates list API
*/
func (a *Client) IPAMAggregatesList(params *IPAMAggregatesListParams) (*IPAMAggregatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMAggregatesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_aggregates_list",
		Method:             "GET",
		PathPattern:        "/api/ipam/aggregates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMAggregatesListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMAggregatesListOK), nil

}

/*
IPAMAggregatesRead ipam aggregates read API
*/
func (a *Client) IPAMAggregatesRead(params *IPAMAggregatesReadParams) (*IPAMAggregatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMAggregatesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_aggregates_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/aggregates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMAggregatesReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMAggregatesReadOK), nil

}

/*
IPAMIPAddressesList ipam ip addresses list API
*/
func (a *Client) IPAMIPAddressesList(params *IPAMIPAddressesListParams) (*IPAMIPAddressesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMIPAddressesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_list",
		Method:             "GET",
		PathPattern:        "/api/ipam/ip-addresses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMIPAddressesListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMIPAddressesListOK), nil

}

/*
IPAMIPAddressesRead ipam ip addresses read API
*/
func (a *Client) IPAMIPAddressesRead(params *IPAMIPAddressesReadParams) (*IPAMIPAddressesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMIPAddressesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/ip-addresses/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMIPAddressesReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMIPAddressesReadOK), nil

}

/*
IPAMPrefixesAvailableIpsRead as convenience method for returning available IP addresses within a prefix by default the number of ips

A convenience method for returning available IP addresses within a prefix. By default, the number of IPs
returned will be equivalent to PAGINATE_COUNT. An arbitrary limit (up to MAX_PAGE_SIZE, if set) may be passed,
however results will not be paginated.
*/
func (a *Client) IPAMPrefixesAvailableIpsRead(params *IPAMPrefixesAvailableIpsReadParams) (*IPAMPrefixesAvailableIpsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMPrefixesAvailableIpsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_prefixes_available-ips_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/prefixes/{id}/available-ips/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMPrefixesAvailableIpsReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMPrefixesAvailableIpsReadOK), nil

}

/*
IPAMPrefixesList ipam prefixes list API
*/
func (a *Client) IPAMPrefixesList(params *IPAMPrefixesListParams) (*IPAMPrefixesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMPrefixesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_prefixes_list",
		Method:             "GET",
		PathPattern:        "/api/ipam/prefixes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMPrefixesListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMPrefixesListOK), nil

}

/*
IPAMPrefixesRead ipam prefixes read API
*/
func (a *Client) IPAMPrefixesRead(params *IPAMPrefixesReadParams) (*IPAMPrefixesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMPrefixesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_prefixes_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/prefixes/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMPrefixesReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMPrefixesReadOK), nil

}

/*
IPAMRirsList ipam rirs list API
*/
func (a *Client) IPAMRirsList(params *IPAMRirsListParams) (*IPAMRirsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMRirsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_rirs_list",
		Method:             "GET",
		PathPattern:        "/api/ipam/rirs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMRirsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMRirsListOK), nil

}

/*
IPAMRirsRead ipam rirs read API
*/
func (a *Client) IPAMRirsRead(params *IPAMRirsReadParams) (*IPAMRirsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMRirsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_rirs_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/rirs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMRirsReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMRirsReadOK), nil

}

/*
IPAMRolesList ipam roles list API
*/
func (a *Client) IPAMRolesList(params *IPAMRolesListParams) (*IPAMRolesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMRolesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_roles_list",
		Method:             "GET",
		PathPattern:        "/api/ipam/roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMRolesListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMRolesListOK), nil

}

/*
IPAMRolesRead ipam roles read API
*/
func (a *Client) IPAMRolesRead(params *IPAMRolesReadParams) (*IPAMRolesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMRolesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_roles_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMRolesReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMRolesReadOK), nil

}

/*
IPAMServicesList ipam services list API
*/
func (a *Client) IPAMServicesList(params *IPAMServicesListParams) (*IPAMServicesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMServicesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_services_list",
		Method:             "GET",
		PathPattern:        "/api/ipam/services/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMServicesListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMServicesListOK), nil

}

/*
IPAMServicesRead ipam services read API
*/
func (a *Client) IPAMServicesRead(params *IPAMServicesReadParams) (*IPAMServicesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMServicesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_services_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/services/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMServicesReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMServicesReadOK), nil

}

/*
IPAMVlanGroupsList ipam vlan groups list API
*/
func (a *Client) IPAMVlanGroupsList(params *IPAMVlanGroupsListParams) (*IPAMVlanGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMVlanGroupsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_list",
		Method:             "GET",
		PathPattern:        "/api/ipam/vlan-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMVlanGroupsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMVlanGroupsListOK), nil

}

/*
IPAMVlanGroupsRead ipam vlan groups read API
*/
func (a *Client) IPAMVlanGroupsRead(params *IPAMVlanGroupsReadParams) (*IPAMVlanGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMVlanGroupsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/vlan-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMVlanGroupsReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMVlanGroupsReadOK), nil

}

/*
IPAMVlansList ipam vlans list API
*/
func (a *Client) IPAMVlansList(params *IPAMVlansListParams) (*IPAMVlansListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMVlansListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_vlans_list",
		Method:             "GET",
		PathPattern:        "/api/ipam/vlans/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMVlansListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMVlansListOK), nil

}

/*
IPAMVlansRead ipam vlans read API
*/
func (a *Client) IPAMVlansRead(params *IPAMVlansReadParams) (*IPAMVlansReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMVlansReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_vlans_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/vlans/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMVlansReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMVlansReadOK), nil

}

/*
IPAMVrfsList ipam vrfs list API
*/
func (a *Client) IPAMVrfsList(params *IPAMVrfsListParams) (*IPAMVrfsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMVrfsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_vrfs_list",
		Method:             "GET",
		PathPattern:        "/api/ipam/vrfs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMVrfsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMVrfsListOK), nil

}

/*
IPAMVrfsRead ipam vrfs read API
*/
func (a *Client) IPAMVrfsRead(params *IPAMVrfsReadParams) (*IPAMVrfsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPAMVrfsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipam_vrfs_read",
		Method:             "GET",
		PathPattern:        "/api/ipam/vrfs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPAMVrfsReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPAMVrfsReadOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
