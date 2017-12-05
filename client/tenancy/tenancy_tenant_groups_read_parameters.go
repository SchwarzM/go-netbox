// Code generated by go-swagger; DO NOT EDIT.

package tenancy

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

// NewTenancyTenantGroupsReadParams creates a new TenancyTenantGroupsReadParams object
// with the default values initialized.
func NewTenancyTenantGroupsReadParams() *TenancyTenantGroupsReadParams {
	var ()
	return &TenancyTenantGroupsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantGroupsReadParamsWithTimeout creates a new TenancyTenantGroupsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTenancyTenantGroupsReadParamsWithTimeout(timeout time.Duration) *TenancyTenantGroupsReadParams {
	var ()
	return &TenancyTenantGroupsReadParams{

		timeout: timeout,
	}
}

// NewTenancyTenantGroupsReadParamsWithContext creates a new TenancyTenantGroupsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewTenancyTenantGroupsReadParamsWithContext(ctx context.Context) *TenancyTenantGroupsReadParams {
	var ()
	return &TenancyTenantGroupsReadParams{

		Context: ctx,
	}
}

// NewTenancyTenantGroupsReadParamsWithHTTPClient creates a new TenancyTenantGroupsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTenancyTenantGroupsReadParamsWithHTTPClient(client *http.Client) *TenancyTenantGroupsReadParams {
	var ()
	return &TenancyTenantGroupsReadParams{
		HTTPClient: client,
	}
}

/*TenancyTenantGroupsReadParams contains all the parameters to send to the API endpoint
for the tenancy tenant groups read operation typically these are written to a http.Request
*/
type TenancyTenantGroupsReadParams struct {

	/*ID
	  A unique integer value identifying this tenant group.

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

// WithTimeout adds the timeout to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) WithTimeout(timeout time.Duration) *TenancyTenantGroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) WithContext(ctx context.Context) *TenancyTenantGroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) WithHTTPClient(client *http.Client) *TenancyTenantGroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) WithID(id int64) *TenancyTenantGroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) WithName(name *string) *TenancyTenantGroupsReadParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) SetName(name *string) {
	o.Name = name
}

// WithSlug adds the slug to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) WithSlug(slug *string) *TenancyTenantGroupsReadParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantGroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
