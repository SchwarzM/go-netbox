// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// NewIPAMAggregatesReadParams creates a new IPAMAggregatesReadParams object
// with the default values initialized.
func NewIPAMAggregatesReadParams() *IPAMAggregatesReadParams {
	var ()
	return &IPAMAggregatesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMAggregatesReadParamsWithTimeout creates a new IPAMAggregatesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMAggregatesReadParamsWithTimeout(timeout time.Duration) *IPAMAggregatesReadParams {
	var ()
	return &IPAMAggregatesReadParams{

		timeout: timeout,
	}
}

// NewIPAMAggregatesReadParamsWithContext creates a new IPAMAggregatesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMAggregatesReadParamsWithContext(ctx context.Context) *IPAMAggregatesReadParams {
	var ()
	return &IPAMAggregatesReadParams{

		Context: ctx,
	}
}

// NewIPAMAggregatesReadParamsWithHTTPClient creates a new IPAMAggregatesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMAggregatesReadParamsWithHTTPClient(client *http.Client) *IPAMAggregatesReadParams {
	var ()
	return &IPAMAggregatesReadParams{
		HTTPClient: client,
	}
}

/*IPAMAggregatesReadParams contains all the parameters to send to the API endpoint
for the ipam aggregates read operation typically these are written to a http.Request
*/
type IPAMAggregatesReadParams struct {

	/*DateAdded*/
	DateAdded *string
	/*Family*/
	Family *string
	/*ID
	  A unique integer value identifying this aggregate.

	*/
	ID int64
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *float64
	/*Q*/
	Q *string
	/*Rir*/
	Rir *string
	/*RirID*/
	RirID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) WithTimeout(timeout time.Duration) *IPAMAggregatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) WithContext(ctx context.Context) *IPAMAggregatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) WithHTTPClient(client *http.Client) *IPAMAggregatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateAdded adds the dateAdded to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) WithDateAdded(dateAdded *string) *IPAMAggregatesReadParams {
	o.SetDateAdded(dateAdded)
	return o
}

// SetDateAdded adds the dateAdded to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) SetDateAdded(dateAdded *string) {
	o.DateAdded = dateAdded
}

// WithFamily adds the family to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) WithFamily(family *string) *IPAMAggregatesReadParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) SetFamily(family *string) {
	o.Family = family
}

// WithID adds the id to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) WithID(id int64) *IPAMAggregatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) SetID(id int64) {
	o.ID = id
}

// WithIDIn adds the iDIn to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) WithIDIn(iDIn *float64) *IPAMAggregatesReadParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) SetIDIn(iDIn *float64) {
	o.IDIn = iDIn
}

// WithQ adds the q to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) WithQ(q *string) *IPAMAggregatesReadParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) SetQ(q *string) {
	o.Q = q
}

// WithRir adds the rir to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) WithRir(rir *string) *IPAMAggregatesReadParams {
	o.SetRir(rir)
	return o
}

// SetRir adds the rir to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) SetRir(rir *string) {
	o.Rir = rir
}

// WithRirID adds the rirID to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) WithRirID(rirID *string) *IPAMAggregatesReadParams {
	o.SetRirID(rirID)
	return o
}

// SetRirID adds the rirId to the ipam aggregates read params
func (o *IPAMAggregatesReadParams) SetRirID(rirID *string) {
	o.RirID = rirID
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMAggregatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DateAdded != nil {

		// query param date_added
		var qrDateAdded string
		if o.DateAdded != nil {
			qrDateAdded = *o.DateAdded
		}
		qDateAdded := qrDateAdded
		if qDateAdded != "" {
			if err := r.SetQueryParam("date_added", qDateAdded); err != nil {
				return err
			}
		}

	}

	if o.Family != nil {

		// query param family
		var qrFamily string
		if o.Family != nil {
			qrFamily = *o.Family
		}
		qFamily := qrFamily
		if qFamily != "" {
			if err := r.SetQueryParam("family", qFamily); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn float64
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := swag.FormatFloat64(qrIDIn)
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Rir != nil {

		// query param rir
		var qrRir string
		if o.Rir != nil {
			qrRir = *o.Rir
		}
		qRir := qrRir
		if qRir != "" {
			if err := r.SetQueryParam("rir", qRir); err != nil {
				return err
			}
		}

	}

	if o.RirID != nil {

		// query param rir_id
		var qrRirID string
		if o.RirID != nil {
			qrRirID = *o.RirID
		}
		qRirID := qrRirID
		if qRirID != "" {
			if err := r.SetQueryParam("rir_id", qRirID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
