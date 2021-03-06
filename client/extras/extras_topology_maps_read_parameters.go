// Code generated by go-swagger; DO NOT EDIT.

package extras

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

// NewExtrasTopologyMapsReadParams creates a new ExtrasTopologyMapsReadParams object
// with the default values initialized.
func NewExtrasTopologyMapsReadParams() *ExtrasTopologyMapsReadParams {
	var ()
	return &ExtrasTopologyMapsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTopologyMapsReadParamsWithTimeout creates a new ExtrasTopologyMapsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTopologyMapsReadParamsWithTimeout(timeout time.Duration) *ExtrasTopologyMapsReadParams {
	var ()
	return &ExtrasTopologyMapsReadParams{

		timeout: timeout,
	}
}

// NewExtrasTopologyMapsReadParamsWithContext creates a new ExtrasTopologyMapsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTopologyMapsReadParamsWithContext(ctx context.Context) *ExtrasTopologyMapsReadParams {
	var ()
	return &ExtrasTopologyMapsReadParams{

		Context: ctx,
	}
}

// NewExtrasTopologyMapsReadParamsWithHTTPClient creates a new ExtrasTopologyMapsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTopologyMapsReadParamsWithHTTPClient(client *http.Client) *ExtrasTopologyMapsReadParams {
	var ()
	return &ExtrasTopologyMapsReadParams{
		HTTPClient: client,
	}
}

/*ExtrasTopologyMapsReadParams contains all the parameters to send to the API endpoint
for the extras topology maps read operation typically these are written to a http.Request
*/
type ExtrasTopologyMapsReadParams struct {

	/*ID
	  A unique integer value identifying this topology map.

	*/
	ID int64
	/*Name*/
	Name *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) WithTimeout(timeout time.Duration) *ExtrasTopologyMapsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) WithContext(ctx context.Context) *ExtrasTopologyMapsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) WithHTTPClient(client *http.Client) *ExtrasTopologyMapsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) WithID(id int64) *ExtrasTopologyMapsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) WithName(name *string) *ExtrasTopologyMapsReadParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) SetName(name *string) {
	o.Name = name
}

// WithSite adds the site to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) WithSite(site *string) *ExtrasTopologyMapsReadParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) WithSiteID(siteID *string) *ExtrasTopologyMapsReadParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSlug adds the slug to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) WithSlug(slug *string) *ExtrasTopologyMapsReadParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the extras topology maps read params
func (o *ExtrasTopologyMapsReadParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTopologyMapsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
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
