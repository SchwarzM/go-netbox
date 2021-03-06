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

// NewExtrasExportTemplatesReadParams creates a new ExtrasExportTemplatesReadParams object
// with the default values initialized.
func NewExtrasExportTemplatesReadParams() *ExtrasExportTemplatesReadParams {
	var ()
	return &ExtrasExportTemplatesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasExportTemplatesReadParamsWithTimeout creates a new ExtrasExportTemplatesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasExportTemplatesReadParamsWithTimeout(timeout time.Duration) *ExtrasExportTemplatesReadParams {
	var ()
	return &ExtrasExportTemplatesReadParams{

		timeout: timeout,
	}
}

// NewExtrasExportTemplatesReadParamsWithContext creates a new ExtrasExportTemplatesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasExportTemplatesReadParamsWithContext(ctx context.Context) *ExtrasExportTemplatesReadParams {
	var ()
	return &ExtrasExportTemplatesReadParams{

		Context: ctx,
	}
}

// NewExtrasExportTemplatesReadParamsWithHTTPClient creates a new ExtrasExportTemplatesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasExportTemplatesReadParamsWithHTTPClient(client *http.Client) *ExtrasExportTemplatesReadParams {
	var ()
	return &ExtrasExportTemplatesReadParams{
		HTTPClient: client,
	}
}

/*ExtrasExportTemplatesReadParams contains all the parameters to send to the API endpoint
for the extras export templates read operation typically these are written to a http.Request
*/
type ExtrasExportTemplatesReadParams struct {

	/*ContentType*/
	ContentType *string
	/*ID
	  A unique integer value identifying this export template.

	*/
	ID int64
	/*Name*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) WithTimeout(timeout time.Duration) *ExtrasExportTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) WithContext(ctx context.Context) *ExtrasExportTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) WithHTTPClient(client *http.Client) *ExtrasExportTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) WithContentType(contentType *string) *ExtrasExportTemplatesReadParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithID adds the id to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) WithID(id int64) *ExtrasExportTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) WithName(name *string) *ExtrasExportTemplatesReadParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasExportTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// query param content_type
		var qrContentType string
		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {
			if err := r.SetQueryParam("content_type", qContentType); err != nil {
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
