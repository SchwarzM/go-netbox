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

// NewDcimRackReservationsReadParams creates a new DcimRackReservationsReadParams object
// with the default values initialized.
func NewDcimRackReservationsReadParams() *DcimRackReservationsReadParams {
	var ()
	return &DcimRackReservationsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackReservationsReadParamsWithTimeout creates a new DcimRackReservationsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackReservationsReadParamsWithTimeout(timeout time.Duration) *DcimRackReservationsReadParams {
	var ()
	return &DcimRackReservationsReadParams{

		timeout: timeout,
	}
}

// NewDcimRackReservationsReadParamsWithContext creates a new DcimRackReservationsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackReservationsReadParamsWithContext(ctx context.Context) *DcimRackReservationsReadParams {
	var ()
	return &DcimRackReservationsReadParams{

		Context: ctx,
	}
}

// NewDcimRackReservationsReadParamsWithHTTPClient creates a new DcimRackReservationsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackReservationsReadParamsWithHTTPClient(client *http.Client) *DcimRackReservationsReadParams {
	var ()
	return &DcimRackReservationsReadParams{
		HTTPClient: client,
	}
}

/*DcimRackReservationsReadParams contains all the parameters to send to the API endpoint
for the dcim rack reservations read operation typically these are written to a http.Request
*/
type DcimRackReservationsReadParams struct {

	/*Created*/
	Created *string
	/*Group*/
	Group *string
	/*GroupID*/
	GroupID *string
	/*ID
	  A unique integer value identifying this rack reservation.

	*/
	ID int64
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *float64
	/*Q*/
	Q *string
	/*RackID*/
	RackID *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*User*/
	User *string
	/*UserID*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithTimeout(timeout time.Duration) *DcimRackReservationsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithContext(ctx context.Context) *DcimRackReservationsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithHTTPClient(client *http.Client) *DcimRackReservationsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithCreated(created *string) *DcimRackReservationsReadParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetCreated(created *string) {
	o.Created = created
}

// WithGroup adds the group to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithGroup(group *string) *DcimRackReservationsReadParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupID adds the groupID to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithGroupID(groupID *string) *DcimRackReservationsReadParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithID adds the id to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithID(id int64) *DcimRackReservationsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetID(id int64) {
	o.ID = id
}

// WithIDIn adds the iDIn to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithIDIn(iDIn *float64) *DcimRackReservationsReadParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetIDIn(iDIn *float64) {
	o.IDIn = iDIn
}

// WithQ adds the q to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithQ(q *string) *DcimRackReservationsReadParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetQ(q *string) {
	o.Q = q
}

// WithRackID adds the rackID to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithRackID(rackID *string) *DcimRackReservationsReadParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetRackID(rackID *string) {
	o.RackID = rackID
}

// WithSite adds the site to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithSite(site *string) *DcimRackReservationsReadParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithSiteID(siteID *string) *DcimRackReservationsReadParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithUser adds the user to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithUser(user *string) *DcimRackReservationsReadParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetUser(user *string) {
	o.User = user
}

// WithUserID adds the userID to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithUserID(userID *string) *DcimRackReservationsReadParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackReservationsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string
		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {
			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}

	}

	if o.Group != nil {

		// query param group
		var qrGroup string
		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {
			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}

	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID string
		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := qrGroupID
		if qGroupID != "" {
			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
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

	if o.RackID != nil {

		// query param rack_id
		var qrRackID string
		if o.RackID != nil {
			qrRackID = *o.RackID
		}
		qRackID := qrRackID
		if qRackID != "" {
			if err := r.SetQueryParam("rack_id", qRackID); err != nil {
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

	if o.User != nil {

		// query param user
		var qrUser string
		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {
			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}

	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
