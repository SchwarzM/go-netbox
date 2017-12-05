// Code generated by go-swagger; DO NOT EDIT.

package dcim

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

// NewDcimPowerPortsReadParams creates a new DcimPowerPortsReadParams object
// with the default values initialized.
func NewDcimPowerPortsReadParams() *DcimPowerPortsReadParams {
	var ()
	return &DcimPowerPortsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsReadParamsWithTimeout creates a new DcimPowerPortsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerPortsReadParamsWithTimeout(timeout time.Duration) *DcimPowerPortsReadParams {
	var ()
	return &DcimPowerPortsReadParams{

		timeout: timeout,
	}
}

// NewDcimPowerPortsReadParamsWithContext creates a new DcimPowerPortsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerPortsReadParamsWithContext(ctx context.Context) *DcimPowerPortsReadParams {
	var ()
	return &DcimPowerPortsReadParams{

		Context: ctx,
	}
}

// NewDcimPowerPortsReadParamsWithHTTPClient creates a new DcimPowerPortsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerPortsReadParamsWithHTTPClient(client *http.Client) *DcimPowerPortsReadParams {
	var ()
	return &DcimPowerPortsReadParams{
		HTTPClient: client,
	}
}

/*DcimPowerPortsReadParams contains all the parameters to send to the API endpoint
for the dcim power ports read operation typically these are written to a http.Request
*/
type DcimPowerPortsReadParams struct {

	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *string
	/*ID
	  A unique integer value identifying this power port.

	*/
	ID int64
	/*Name*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power ports read params
func (o *DcimPowerPortsReadParams) WithTimeout(timeout time.Duration) *DcimPowerPortsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports read params
func (o *DcimPowerPortsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports read params
func (o *DcimPowerPortsReadParams) WithContext(ctx context.Context) *DcimPowerPortsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports read params
func (o *DcimPowerPortsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports read params
func (o *DcimPowerPortsReadParams) WithHTTPClient(client *http.Client) *DcimPowerPortsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports read params
func (o *DcimPowerPortsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the dcim power ports read params
func (o *DcimPowerPortsReadParams) WithDevice(device *string) *DcimPowerPortsReadParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim power ports read params
func (o *DcimPowerPortsReadParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim power ports read params
func (o *DcimPowerPortsReadParams) WithDeviceID(deviceID *string) *DcimPowerPortsReadParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim power ports read params
func (o *DcimPowerPortsReadParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithID adds the id to the dcim power ports read params
func (o *DcimPowerPortsReadParams) WithID(id int64) *DcimPowerPortsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power ports read params
func (o *DcimPowerPortsReadParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the dcim power ports read params
func (o *DcimPowerPortsReadParams) WithName(name *string) *DcimPowerPortsReadParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power ports read params
func (o *DcimPowerPortsReadParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
