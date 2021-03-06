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

// NewDcimRackRolesReadParams creates a new DcimRackRolesReadParams object
// with the default values initialized.
func NewDcimRackRolesReadParams() *DcimRackRolesReadParams {
	var ()
	return &DcimRackRolesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackRolesReadParamsWithTimeout creates a new DcimRackRolesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackRolesReadParamsWithTimeout(timeout time.Duration) *DcimRackRolesReadParams {
	var ()
	return &DcimRackRolesReadParams{

		timeout: timeout,
	}
}

// NewDcimRackRolesReadParamsWithContext creates a new DcimRackRolesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackRolesReadParamsWithContext(ctx context.Context) *DcimRackRolesReadParams {
	var ()
	return &DcimRackRolesReadParams{

		Context: ctx,
	}
}

// NewDcimRackRolesReadParamsWithHTTPClient creates a new DcimRackRolesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackRolesReadParamsWithHTTPClient(client *http.Client) *DcimRackRolesReadParams {
	var ()
	return &DcimRackRolesReadParams{
		HTTPClient: client,
	}
}

/*DcimRackRolesReadParams contains all the parameters to send to the API endpoint
for the dcim rack roles read operation typically these are written to a http.Request
*/
type DcimRackRolesReadParams struct {

	/*Color*/
	Color *string
	/*ID
	  A unique integer value identifying this rack role.

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

// WithTimeout adds the timeout to the dcim rack roles read params
func (o *DcimRackRolesReadParams) WithTimeout(timeout time.Duration) *DcimRackRolesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack roles read params
func (o *DcimRackRolesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack roles read params
func (o *DcimRackRolesReadParams) WithContext(ctx context.Context) *DcimRackRolesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack roles read params
func (o *DcimRackRolesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack roles read params
func (o *DcimRackRolesReadParams) WithHTTPClient(client *http.Client) *DcimRackRolesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack roles read params
func (o *DcimRackRolesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColor adds the color to the dcim rack roles read params
func (o *DcimRackRolesReadParams) WithColor(color *string) *DcimRackRolesReadParams {
	o.SetColor(color)
	return o
}

// SetColor adds the color to the dcim rack roles read params
func (o *DcimRackRolesReadParams) SetColor(color *string) {
	o.Color = color
}

// WithID adds the id to the dcim rack roles read params
func (o *DcimRackRolesReadParams) WithID(id int64) *DcimRackRolesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack roles read params
func (o *DcimRackRolesReadParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the dcim rack roles read params
func (o *DcimRackRolesReadParams) WithName(name *string) *DcimRackRolesReadParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim rack roles read params
func (o *DcimRackRolesReadParams) SetName(name *string) {
	o.Name = name
}

// WithSlug adds the slug to the dcim rack roles read params
func (o *DcimRackRolesReadParams) WithSlug(slug *string) *DcimRackRolesReadParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim rack roles read params
func (o *DcimRackRolesReadParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackRolesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Color != nil {

		// query param color
		var qrColor string
		if o.Color != nil {
			qrColor = *o.Color
		}
		qColor := qrColor
		if qColor != "" {
			if err := r.SetQueryParam("color", qColor); err != nil {
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
