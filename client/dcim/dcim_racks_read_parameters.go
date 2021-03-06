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

// NewDcimRacksReadParams creates a new DcimRacksReadParams object
// with the default values initialized.
func NewDcimRacksReadParams() *DcimRacksReadParams {
	var ()
	return &DcimRacksReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksReadParamsWithTimeout creates a new DcimRacksReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRacksReadParamsWithTimeout(timeout time.Duration) *DcimRacksReadParams {
	var ()
	return &DcimRacksReadParams{

		timeout: timeout,
	}
}

// NewDcimRacksReadParamsWithContext creates a new DcimRacksReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRacksReadParamsWithContext(ctx context.Context) *DcimRacksReadParams {
	var ()
	return &DcimRacksReadParams{

		Context: ctx,
	}
}

// NewDcimRacksReadParamsWithHTTPClient creates a new DcimRacksReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRacksReadParamsWithHTTPClient(client *http.Client) *DcimRacksReadParams {
	var ()
	return &DcimRacksReadParams{
		HTTPClient: client,
	}
}

/*DcimRacksReadParams contains all the parameters to send to the API endpoint
for the dcim racks read operation typically these are written to a http.Request
*/
type DcimRacksReadParams struct {

	/*DescUnits*/
	DescUnits *string
	/*FacilityID*/
	FacilityID *string
	/*Group*/
	Group *string
	/*GroupID*/
	GroupID *string
	/*ID
	  A unique integer value identifying this rack.

	*/
	ID int64
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *float64
	/*Q*/
	Q *string
	/*Role*/
	Role *string
	/*RoleID*/
	RoleID *string
	/*Serial*/
	Serial *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *string
	/*Type*/
	Type *string
	/*UHeight*/
	UHeight *float64
	/*Width*/
	Width *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim racks read params
func (o *DcimRacksReadParams) WithTimeout(timeout time.Duration) *DcimRacksReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks read params
func (o *DcimRacksReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks read params
func (o *DcimRacksReadParams) WithContext(ctx context.Context) *DcimRacksReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks read params
func (o *DcimRacksReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks read params
func (o *DcimRacksReadParams) WithHTTPClient(client *http.Client) *DcimRacksReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks read params
func (o *DcimRacksReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescUnits adds the descUnits to the dcim racks read params
func (o *DcimRacksReadParams) WithDescUnits(descUnits *string) *DcimRacksReadParams {
	o.SetDescUnits(descUnits)
	return o
}

// SetDescUnits adds the descUnits to the dcim racks read params
func (o *DcimRacksReadParams) SetDescUnits(descUnits *string) {
	o.DescUnits = descUnits
}

// WithFacilityID adds the facilityID to the dcim racks read params
func (o *DcimRacksReadParams) WithFacilityID(facilityID *string) *DcimRacksReadParams {
	o.SetFacilityID(facilityID)
	return o
}

// SetFacilityID adds the facilityId to the dcim racks read params
func (o *DcimRacksReadParams) SetFacilityID(facilityID *string) {
	o.FacilityID = facilityID
}

// WithGroup adds the group to the dcim racks read params
func (o *DcimRacksReadParams) WithGroup(group *string) *DcimRacksReadParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the dcim racks read params
func (o *DcimRacksReadParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupID adds the groupID to the dcim racks read params
func (o *DcimRacksReadParams) WithGroupID(groupID *string) *DcimRacksReadParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the dcim racks read params
func (o *DcimRacksReadParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithID adds the id to the dcim racks read params
func (o *DcimRacksReadParams) WithID(id int64) *DcimRacksReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks read params
func (o *DcimRacksReadParams) SetID(id int64) {
	o.ID = id
}

// WithIDIn adds the iDIn to the dcim racks read params
func (o *DcimRacksReadParams) WithIDIn(iDIn *float64) *DcimRacksReadParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the dcim racks read params
func (o *DcimRacksReadParams) SetIDIn(iDIn *float64) {
	o.IDIn = iDIn
}

// WithQ adds the q to the dcim racks read params
func (o *DcimRacksReadParams) WithQ(q *string) *DcimRacksReadParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim racks read params
func (o *DcimRacksReadParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the dcim racks read params
func (o *DcimRacksReadParams) WithRole(role *string) *DcimRacksReadParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the dcim racks read params
func (o *DcimRacksReadParams) SetRole(role *string) {
	o.Role = role
}

// WithRoleID adds the roleID to the dcim racks read params
func (o *DcimRacksReadParams) WithRoleID(roleID *string) *DcimRacksReadParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the dcim racks read params
func (o *DcimRacksReadParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithSerial adds the serial to the dcim racks read params
func (o *DcimRacksReadParams) WithSerial(serial *string) *DcimRacksReadParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the dcim racks read params
func (o *DcimRacksReadParams) SetSerial(serial *string) {
	o.Serial = serial
}

// WithSite adds the site to the dcim racks read params
func (o *DcimRacksReadParams) WithSite(site *string) *DcimRacksReadParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim racks read params
func (o *DcimRacksReadParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the dcim racks read params
func (o *DcimRacksReadParams) WithSiteID(siteID *string) *DcimRacksReadParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim racks read params
func (o *DcimRacksReadParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithTenant adds the tenant to the dcim racks read params
func (o *DcimRacksReadParams) WithTenant(tenant *string) *DcimRacksReadParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim racks read params
func (o *DcimRacksReadParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the dcim racks read params
func (o *DcimRacksReadParams) WithTenantID(tenantID *string) *DcimRacksReadParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim racks read params
func (o *DcimRacksReadParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithType adds the typeVar to the dcim racks read params
func (o *DcimRacksReadParams) WithType(typeVar *string) *DcimRacksReadParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim racks read params
func (o *DcimRacksReadParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUHeight adds the uHeight to the dcim racks read params
func (o *DcimRacksReadParams) WithUHeight(uHeight *float64) *DcimRacksReadParams {
	o.SetUHeight(uHeight)
	return o
}

// SetUHeight adds the uHeight to the dcim racks read params
func (o *DcimRacksReadParams) SetUHeight(uHeight *float64) {
	o.UHeight = uHeight
}

// WithWidth adds the width to the dcim racks read params
func (o *DcimRacksReadParams) WithWidth(width *string) *DcimRacksReadParams {
	o.SetWidth(width)
	return o
}

// SetWidth adds the width to the dcim racks read params
func (o *DcimRacksReadParams) SetWidth(width *string) {
	o.Width = width
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DescUnits != nil {

		// query param desc_units
		var qrDescUnits string
		if o.DescUnits != nil {
			qrDescUnits = *o.DescUnits
		}
		qDescUnits := qrDescUnits
		if qDescUnits != "" {
			if err := r.SetQueryParam("desc_units", qDescUnits); err != nil {
				return err
			}
		}

	}

	if o.FacilityID != nil {

		// query param facility_id
		var qrFacilityID string
		if o.FacilityID != nil {
			qrFacilityID = *o.FacilityID
		}
		qFacilityID := qrFacilityID
		if qFacilityID != "" {
			if err := r.SetQueryParam("facility_id", qFacilityID); err != nil {
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

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID string
		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {
			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
				return err
			}
		}

	}

	if o.Serial != nil {

		// query param serial
		var qrSerial string
		if o.Serial != nil {
			qrSerial = *o.Serial
		}
		qSerial := qrSerial
		if qSerial != "" {
			if err := r.SetQueryParam("serial", qSerial); err != nil {
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

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string
		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {
			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}

	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if o.UHeight != nil {

		// query param u_height
		var qrUHeight float64
		if o.UHeight != nil {
			qrUHeight = *o.UHeight
		}
		qUHeight := swag.FormatFloat64(qrUHeight)
		if qUHeight != "" {
			if err := r.SetQueryParam("u_height", qUHeight); err != nil {
				return err
			}
		}

	}

	if o.Width != nil {

		// query param width
		var qrWidth string
		if o.Width != nil {
			qrWidth = *o.Width
		}
		qWidth := qrWidth
		if qWidth != "" {
			if err := r.SetQueryParam("width", qWidth); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
