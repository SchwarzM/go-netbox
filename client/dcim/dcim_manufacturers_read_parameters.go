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

// NewDcimManufacturersReadParams creates a new DcimManufacturersReadParams object
// with the default values initialized.
func NewDcimManufacturersReadParams() *DcimManufacturersReadParams {
	var ()
	return &DcimManufacturersReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersReadParamsWithTimeout creates a new DcimManufacturersReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimManufacturersReadParamsWithTimeout(timeout time.Duration) *DcimManufacturersReadParams {
	var ()
	return &DcimManufacturersReadParams{

		timeout: timeout,
	}
}

// NewDcimManufacturersReadParamsWithContext creates a new DcimManufacturersReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimManufacturersReadParamsWithContext(ctx context.Context) *DcimManufacturersReadParams {
	var ()
	return &DcimManufacturersReadParams{

		Context: ctx,
	}
}

// NewDcimManufacturersReadParamsWithHTTPClient creates a new DcimManufacturersReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimManufacturersReadParamsWithHTTPClient(client *http.Client) *DcimManufacturersReadParams {
	var ()
	return &DcimManufacturersReadParams{
		HTTPClient: client,
	}
}

/*DcimManufacturersReadParams contains all the parameters to send to the API endpoint
for the dcim manufacturers read operation typically these are written to a http.Request
*/
type DcimManufacturersReadParams struct {

	/*ID
	  A unique integer value identifying this manufacturer.

	*/
	ID int64
	/*Name*/
	Name *string
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) WithTimeout(timeout time.Duration) *DcimManufacturersReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) WithContext(ctx context.Context) *DcimManufacturersReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) WithHTTPClient(client *http.Client) *DcimManufacturersReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) WithID(id int64) *DcimManufacturersReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) WithName(name *string) *DcimManufacturersReadParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) SetName(name *string) {
	o.Name = name
}

// WithSlug adds the slug to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) WithSlug(slug *string) *DcimManufacturersReadParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
