// Code generated by go-swagger; DO NOT EDIT.

package virtualization

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

// NewVirtualizationVirtualMachinesReadParams creates a new VirtualizationVirtualMachinesReadParams object
// with the default values initialized.
func NewVirtualizationVirtualMachinesReadParams() *VirtualizationVirtualMachinesReadParams {
	var ()
	return &VirtualizationVirtualMachinesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesReadParamsWithTimeout creates a new VirtualizationVirtualMachinesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationVirtualMachinesReadParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesReadParams {
	var ()
	return &VirtualizationVirtualMachinesReadParams{

		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesReadParamsWithContext creates a new VirtualizationVirtualMachinesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationVirtualMachinesReadParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesReadParams {
	var ()
	return &VirtualizationVirtualMachinesReadParams{

		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesReadParamsWithHTTPClient creates a new VirtualizationVirtualMachinesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationVirtualMachinesReadParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesReadParams {
	var ()
	return &VirtualizationVirtualMachinesReadParams{
		HTTPClient: client,
	}
}

/*VirtualizationVirtualMachinesReadParams contains all the parameters to send to the API endpoint
for the virtualization virtual machines read operation typically these are written to a http.Request
*/
type VirtualizationVirtualMachinesReadParams struct {

	/*Cluster*/
	Cluster *string
	/*ClusterGroup*/
	ClusterGroup *string
	/*ClusterGroupID*/
	ClusterGroupID *string
	/*ClusterID*/
	ClusterID *string
	/*ID
	  A unique integer value identifying this virtual machine.

	*/
	ID int64
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *float64
	/*Name*/
	Name *string
	/*Platform*/
	Platform *string
	/*PlatformID*/
	PlatformID *string
	/*Q*/
	Q *string
	/*Role*/
	Role *string
	/*RoleID*/
	RoleID *string
	/*Status*/
	Status *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithCluster(cluster *string) *VirtualizationVirtualMachinesReadParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetCluster(cluster *string) {
	o.Cluster = cluster
}

// WithClusterGroup adds the clusterGroup to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithClusterGroup(clusterGroup *string) *VirtualizationVirtualMachinesReadParams {
	o.SetClusterGroup(clusterGroup)
	return o
}

// SetClusterGroup adds the clusterGroup to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetClusterGroup(clusterGroup *string) {
	o.ClusterGroup = clusterGroup
}

// WithClusterGroupID adds the clusterGroupID to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithClusterGroupID(clusterGroupID *string) *VirtualizationVirtualMachinesReadParams {
	o.SetClusterGroupID(clusterGroupID)
	return o
}

// SetClusterGroupID adds the clusterGroupId to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetClusterGroupID(clusterGroupID *string) {
	o.ClusterGroupID = clusterGroupID
}

// WithClusterID adds the clusterID to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithClusterID(clusterID *string) *VirtualizationVirtualMachinesReadParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithID adds the id to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithID(id int64) *VirtualizationVirtualMachinesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetID(id int64) {
	o.ID = id
}

// WithIDIn adds the iDIn to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithIDIn(iDIn *float64) *VirtualizationVirtualMachinesReadParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetIDIn(iDIn *float64) {
	o.IDIn = iDIn
}

// WithName adds the name to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithName(name *string) *VirtualizationVirtualMachinesReadParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetName(name *string) {
	o.Name = name
}

// WithPlatform adds the platform to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithPlatform(platform *string) *VirtualizationVirtualMachinesReadParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPlatformID adds the platformID to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithPlatformID(platformID *string) *VirtualizationVirtualMachinesReadParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetPlatformID(platformID *string) {
	o.PlatformID = platformID
}

// WithQ adds the q to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithQ(q *string) *VirtualizationVirtualMachinesReadParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithRole(role *string) *VirtualizationVirtualMachinesReadParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetRole(role *string) {
	o.Role = role
}

// WithRoleID adds the roleID to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithRoleID(roleID *string) *VirtualizationVirtualMachinesReadParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithStatus adds the status to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithStatus(status *string) *VirtualizationVirtualMachinesReadParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetStatus(status *string) {
	o.Status = status
}

// WithTenant adds the tenant to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithTenant(tenant *string) *VirtualizationVirtualMachinesReadParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithTenantID(tenantID *string) *VirtualizationVirtualMachinesReadParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cluster != nil {

		// query param cluster
		var qrCluster string
		if o.Cluster != nil {
			qrCluster = *o.Cluster
		}
		qCluster := qrCluster
		if qCluster != "" {
			if err := r.SetQueryParam("cluster", qCluster); err != nil {
				return err
			}
		}

	}

	if o.ClusterGroup != nil {

		// query param cluster_group
		var qrClusterGroup string
		if o.ClusterGroup != nil {
			qrClusterGroup = *o.ClusterGroup
		}
		qClusterGroup := qrClusterGroup
		if qClusterGroup != "" {
			if err := r.SetQueryParam("cluster_group", qClusterGroup); err != nil {
				return err
			}
		}

	}

	if o.ClusterGroupID != nil {

		// query param cluster_group_id
		var qrClusterGroupID string
		if o.ClusterGroupID != nil {
			qrClusterGroupID = *o.ClusterGroupID
		}
		qClusterGroupID := qrClusterGroupID
		if qClusterGroupID != "" {
			if err := r.SetQueryParam("cluster_group_id", qClusterGroupID); err != nil {
				return err
			}
		}

	}

	if o.ClusterID != nil {

		// query param cluster_id
		var qrClusterID string
		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {
			if err := r.SetQueryParam("cluster_id", qClusterID); err != nil {
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

	if o.Platform != nil {

		// query param platform
		var qrPlatform string
		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {
			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}

	}

	if o.PlatformID != nil {

		// query param platform_id
		var qrPlatformID string
		if o.PlatformID != nil {
			qrPlatformID = *o.PlatformID
		}
		qPlatformID := qrPlatformID
		if qPlatformID != "" {
			if err := r.SetQueryParam("platform_id", qPlatformID); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
