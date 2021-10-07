// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewShowZoneDetailsParams creates a new ShowZoneDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowZoneDetailsParams() *ShowZoneDetailsParams {
	return &ShowZoneDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowZoneDetailsParamsWithTimeout creates a new ShowZoneDetailsParams object
// with the ability to set a timeout on a request.
func NewShowZoneDetailsParamsWithTimeout(timeout time.Duration) *ShowZoneDetailsParams {
	return &ShowZoneDetailsParams{
		timeout: timeout,
	}
}

// NewShowZoneDetailsParamsWithContext creates a new ShowZoneDetailsParams object
// with the ability to set a context for a request.
func NewShowZoneDetailsParamsWithContext(ctx context.Context) *ShowZoneDetailsParams {
	return &ShowZoneDetailsParams{
		Context: ctx,
	}
}

// NewShowZoneDetailsParamsWithHTTPClient creates a new ShowZoneDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowZoneDetailsParamsWithHTTPClient(client *http.Client) *ShowZoneDetailsParams {
	return &ShowZoneDetailsParams{
		HTTPClient: client,
	}
}

/* ShowZoneDetailsParams contains all the parameters to send to the API endpoint
   for the show zone details operation.

   Typically these are written to a http.Request.
*/
type ShowZoneDetailsParams struct {

	/* HomeID.

	   The ID of a home.

	   Format: int64
	*/
	HomeID int64

	/* ZoneID.

	   The ID of a zone.

	   Format: int64
	*/
	ZoneID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the show zone details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowZoneDetailsParams) WithDefaults() *ShowZoneDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the show zone details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowZoneDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the show zone details params
func (o *ShowZoneDetailsParams) WithTimeout(timeout time.Duration) *ShowZoneDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show zone details params
func (o *ShowZoneDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show zone details params
func (o *ShowZoneDetailsParams) WithContext(ctx context.Context) *ShowZoneDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show zone details params
func (o *ShowZoneDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show zone details params
func (o *ShowZoneDetailsParams) WithHTTPClient(client *http.Client) *ShowZoneDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show zone details params
func (o *ShowZoneDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHomeID adds the homeID to the show zone details params
func (o *ShowZoneDetailsParams) WithHomeID(homeID int64) *ShowZoneDetailsParams {
	o.SetHomeID(homeID)
	return o
}

// SetHomeID adds the homeId to the show zone details params
func (o *ShowZoneDetailsParams) SetHomeID(homeID int64) {
	o.HomeID = homeID
}

// WithZoneID adds the zoneID to the show zone details params
func (o *ShowZoneDetailsParams) WithZoneID(zoneID int64) *ShowZoneDetailsParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the show zone details params
func (o *ShowZoneDetailsParams) SetZoneID(zoneID int64) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *ShowZoneDetailsParams) WriteToRequest(r runtime.ClientRequest, _ strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param home_id
	if err := r.SetPathParam("home_id", swag.FormatInt64(o.HomeID)); err != nil {
		return err
	}

	// path param zone_id
	if err := r.SetPathParam("zone_id", swag.FormatInt64(o.ZoneID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
