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

// NewDcimPowerOutletTemplatesListParams creates a new DcimPowerOutletTemplatesListParams object
// with the default values initialized.
func NewDcimPowerOutletTemplatesListParams() *DcimPowerOutletTemplatesListParams {
	var ()
	return &DcimPowerOutletTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletTemplatesListParamsWithTimeout creates a new DcimPowerOutletTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerOutletTemplatesListParamsWithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesListParams {
	var ()
	return &DcimPowerOutletTemplatesListParams{

		timeout: timeout,
	}
}

// NewDcimPowerOutletTemplatesListParamsWithContext creates a new DcimPowerOutletTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerOutletTemplatesListParamsWithContext(ctx context.Context) *DcimPowerOutletTemplatesListParams {
	var ()
	return &DcimPowerOutletTemplatesListParams{

		Context: ctx,
	}
}

// NewDcimPowerOutletTemplatesListParamsWithHTTPClient creates a new DcimPowerOutletTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerOutletTemplatesListParamsWithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesListParams {
	var ()
	return &DcimPowerOutletTemplatesListParams{
		HTTPClient: client,
	}
}

/*DcimPowerOutletTemplatesListParams contains all the parameters to send to the API endpoint
for the dcim power outlet templates list operation typically these are written to a http.Request
*/
type DcimPowerOutletTemplatesListParams struct {

	/*DevicetypeID*/
	DevicetypeID *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) WithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) WithContext(ctx context.Context) *DcimPowerOutletTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) WithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimPowerOutletTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithLimit adds the limit to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) WithLimit(limit *int64) *DcimPowerOutletTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) WithName(name *string) *DcimPowerOutletTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) WithOffset(offset *int64) *DcimPowerOutletTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power outlet templates list params
func (o *DcimPowerOutletTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
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

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
