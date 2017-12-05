// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewVirtualizationInterfacesReadParams creates a new VirtualizationInterfacesReadParams object
// with the default values initialized.
func NewVirtualizationInterfacesReadParams() *VirtualizationInterfacesReadParams {
	var ()
	return &VirtualizationInterfacesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesReadParamsWithTimeout creates a new VirtualizationInterfacesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationInterfacesReadParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesReadParams {
	var ()
	return &VirtualizationInterfacesReadParams{

		timeout: timeout,
	}
}

// NewVirtualizationInterfacesReadParamsWithContext creates a new VirtualizationInterfacesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationInterfacesReadParamsWithContext(ctx context.Context) *VirtualizationInterfacesReadParams {
	var ()
	return &VirtualizationInterfacesReadParams{

		Context: ctx,
	}
}

// NewVirtualizationInterfacesReadParamsWithHTTPClient creates a new VirtualizationInterfacesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationInterfacesReadParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesReadParams {
	var ()
	return &VirtualizationInterfacesReadParams{
		HTTPClient: client,
	}
}

/*VirtualizationInterfacesReadParams contains all the parameters to send to the API endpoint
for the virtualization interfaces read operation typically these are written to a http.Request
*/
type VirtualizationInterfacesReadParams struct {

	/*Enabled*/
	Enabled *string
	/*ID
	  A unique integer value identifying this interface.

	*/
	ID int64
	/*MacAddress*/
	MacAddress *string
	/*Mtu*/
	Mtu *float64
	/*Name*/
	Name *string
	/*VirtualMachine*/
	VirtualMachine *string
	/*VirtualMachineID*/
	VirtualMachineID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithContext(ctx context.Context) *VirtualizationInterfacesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithEnabled(enabled *string) *VirtualizationInterfacesReadParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithID adds the id to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithID(id int64) *VirtualizationInterfacesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetID(id int64) {
	o.ID = id
}

// WithMacAddress adds the macAddress to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithMacAddress(macAddress *string) *VirtualizationInterfacesReadParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithMtu adds the mtu to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithMtu(mtu *float64) *VirtualizationInterfacesReadParams {
	o.SetMtu(mtu)
	return o
}

// SetMtu adds the mtu to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetMtu(mtu *float64) {
	o.Mtu = mtu
}

// WithName adds the name to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithName(name *string) *VirtualizationInterfacesReadParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetName(name *string) {
	o.Name = name
}

// WithVirtualMachine adds the virtualMachine to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithVirtualMachine(virtualMachine *string) *VirtualizationInterfacesReadParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetVirtualMachine(virtualMachine *string) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachineID adds the virtualMachineID to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithVirtualMachineID(virtualMachineID *string) *VirtualizationInterfacesReadParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetVirtualMachineID(virtualMachineID *string) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled string
		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := qrEnabled
		if qEnabled != "" {
			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MacAddress != nil {

		// query param mac_address
		var qrMacAddress string
		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {
			if err := r.SetQueryParam("mac_address", qMacAddress); err != nil {
				return err
			}
		}

	}

	if o.Mtu != nil {

		// query param mtu
		var qrMtu float64
		if o.Mtu != nil {
			qrMtu = *o.Mtu
		}
		qMtu := swag.FormatFloat64(qrMtu)
		if qMtu != "" {
			if err := r.SetQueryParam("mtu", qMtu); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.VirtualMachine != nil {

		// query param virtual_machine
		var qrVirtualMachine string
		if o.VirtualMachine != nil {
			qrVirtualMachine = *o.VirtualMachine
		}
		qVirtualMachine := qrVirtualMachine
		if qVirtualMachine != "" {
			if err := r.SetQueryParam("virtual_machine", qVirtualMachine); err != nil {
				return err
			}
		}

	}

	if o.VirtualMachineID != nil {

		// query param virtual_machine_id
		var qrVirtualMachineID string
		if o.VirtualMachineID != nil {
			qrVirtualMachineID = *o.VirtualMachineID
		}
		qVirtualMachineID := qrVirtualMachineID
		if qVirtualMachineID != "" {
			if err := r.SetQueryParam("virtual_machine_id", qVirtualMachineID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
