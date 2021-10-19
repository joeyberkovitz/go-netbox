// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewDcimLocationsBulkPartialUpdateParams creates a new DcimLocationsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimLocationsBulkPartialUpdateParams() *DcimLocationsBulkPartialUpdateParams {
	return &DcimLocationsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimLocationsBulkPartialUpdateParamsWithTimeout creates a new DcimLocationsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimLocationsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimLocationsBulkPartialUpdateParams {
	return &DcimLocationsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimLocationsBulkPartialUpdateParamsWithContext creates a new DcimLocationsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimLocationsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimLocationsBulkPartialUpdateParams {
	return &DcimLocationsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimLocationsBulkPartialUpdateParamsWithHTTPClient creates a new DcimLocationsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimLocationsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimLocationsBulkPartialUpdateParams {
	return &DcimLocationsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimLocationsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim locations bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimLocationsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableLocation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim locations bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimLocationsBulkPartialUpdateParams) WithDefaults() *DcimLocationsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim locations bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimLocationsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim locations bulk partial update params
func (o *DcimLocationsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimLocationsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim locations bulk partial update params
func (o *DcimLocationsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim locations bulk partial update params
func (o *DcimLocationsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimLocationsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim locations bulk partial update params
func (o *DcimLocationsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim locations bulk partial update params
func (o *DcimLocationsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimLocationsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim locations bulk partial update params
func (o *DcimLocationsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim locations bulk partial update params
func (o *DcimLocationsBulkPartialUpdateParams) WithData(data *models.WritableLocation) *DcimLocationsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim locations bulk partial update params
func (o *DcimLocationsBulkPartialUpdateParams) SetData(data *models.WritableLocation) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimLocationsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
