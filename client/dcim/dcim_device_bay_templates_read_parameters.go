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

// NewDcimDeviceBayTemplatesReadParams creates a new DcimDeviceBayTemplatesReadParams object
// with the default values initialized.
func NewDcimDeviceBayTemplatesReadParams() *DcimDeviceBayTemplatesReadParams {
	var ()
	return &DcimDeviceBayTemplatesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesReadParamsWithTimeout creates a new DcimDeviceBayTemplatesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceBayTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesReadParams {
	var ()
	return &DcimDeviceBayTemplatesReadParams{

		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesReadParamsWithContext creates a new DcimDeviceBayTemplatesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceBayTemplatesReadParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesReadParams {
	var ()
	return &DcimDeviceBayTemplatesReadParams{

		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesReadParamsWithHTTPClient creates a new DcimDeviceBayTemplatesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceBayTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesReadParams {
	var ()
	return &DcimDeviceBayTemplatesReadParams{
		HTTPClient: client,
	}
}

/*DcimDeviceBayTemplatesReadParams contains all the parameters to send to the API endpoint
for the dcim device bay templates read operation typically these are written to a http.Request
*/
type DcimDeviceBayTemplatesReadParams struct {

	/*DevicetypeID*/
	DevicetypeID *string
	/*ID
	  A unique integer value identifying this device bay template.

	*/
	ID int64
	/*Name*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) WithDevicetypeID(devicetypeID *string) *DcimDeviceBayTemplatesReadParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithID adds the id to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) WithID(id int64) *DcimDeviceBayTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) WithName(name *string) *DcimDeviceBayTemplatesReadParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim device bay templates read params
func (o *DcimDeviceBayTemplatesReadParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID string
		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := qrDevicetypeID
		if qDevicetypeID != "" {
			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
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
