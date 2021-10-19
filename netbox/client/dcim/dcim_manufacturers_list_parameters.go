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
	"github.com/go-openapi/swag"
)

// NewDcimManufacturersListParams creates a new DcimManufacturersListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimManufacturersListParams() *DcimManufacturersListParams {
	return &DcimManufacturersListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersListParamsWithTimeout creates a new DcimManufacturersListParams object
// with the ability to set a timeout on a request.
func NewDcimManufacturersListParamsWithTimeout(timeout time.Duration) *DcimManufacturersListParams {
	return &DcimManufacturersListParams{
		timeout: timeout,
	}
}

// NewDcimManufacturersListParamsWithContext creates a new DcimManufacturersListParams object
// with the ability to set a context for a request.
func NewDcimManufacturersListParamsWithContext(ctx context.Context) *DcimManufacturersListParams {
	return &DcimManufacturersListParams{
		Context: ctx,
	}
}

// NewDcimManufacturersListParamsWithHTTPClient creates a new DcimManufacturersListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimManufacturersListParamsWithHTTPClient(client *http.Client) *DcimManufacturersListParams {
	return &DcimManufacturersListParams{
		HTTPClient: client,
	}
}

/* DcimManufacturersListParams contains all the parameters to send to the API endpoint
   for the dcim manufacturers list operation.

   Typically these are written to a http.Request.
*/
type DcimManufacturersListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Description.
	Description *string

	// DescriptionEmpty.
	DescriptionEmpty *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameEmpty.
	NameEmpty *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// Slug.
	Slug *string

	// SlugEmpty.
	SlugEmpty *string

	// SlugIc.
	SlugIc *string

	// SlugIe.
	SlugIe *string

	// SlugIew.
	SlugIew *string

	// SlugIsw.
	SlugIsw *string

	// Slugn.
	Slugn *string

	// SlugNic.
	SlugNic *string

	// SlugNie.
	SlugNie *string

	// SlugNiew.
	SlugNiew *string

	// SlugNisw.
	SlugNisw *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim manufacturers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersListParams) WithDefaults() *DcimManufacturersListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim manufacturers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithTimeout(timeout time.Duration) *DcimManufacturersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithContext(ctx context.Context) *DcimManufacturersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithHTTPClient(client *http.Client) *DcimManufacturersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithCreated(created *string) *DcimManufacturersListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithCreatedGte(createdGte *string) *DcimManufacturersListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithCreatedLte(createdLte *string) *DcimManufacturersListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDescription adds the description to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescription(description *string) *DcimManufacturersListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionEmpty adds the descriptionEmpty to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescriptionEmpty(descriptionEmpty *string) *DcimManufacturersListParams {
	o.SetDescriptionEmpty(descriptionEmpty)
	return o
}

// SetDescriptionEmpty adds the descriptionEmpty to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescriptionEmpty(descriptionEmpty *string) {
	o.DescriptionEmpty = descriptionEmpty
}

// WithDescriptionIc adds the descriptionIc to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescriptionIc(descriptionIc *string) *DcimManufacturersListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescriptionIe(descriptionIe *string) *DcimManufacturersListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescriptionIew(descriptionIew *string) *DcimManufacturersListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescriptionIsw(descriptionIsw *string) *DcimManufacturersListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescriptionn(descriptionn *string) *DcimManufacturersListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescriptionNic(descriptionNic *string) *DcimManufacturersListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescriptionNie(descriptionNie *string) *DcimManufacturersListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescriptionNiew(descriptionNiew *string) *DcimManufacturersListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithDescriptionNisw(descriptionNisw *string) *DcimManufacturersListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithID adds the id to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithID(id *string) *DcimManufacturersListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithIDGt(iDGt *string) *DcimManufacturersListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithIDGte(iDGte *string) *DcimManufacturersListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithIDLt(iDLt *string) *DcimManufacturersListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithIDLte(iDLte *string) *DcimManufacturersListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithIDn(iDn *string) *DcimManufacturersListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithLastUpdated(lastUpdated *string) *DcimManufacturersListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimManufacturersListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimManufacturersListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithLimit(limit *int64) *DcimManufacturersListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithName(name *string) *DcimManufacturersListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithNameEmpty(nameEmpty *string) *DcimManufacturersListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithNameIc(nameIc *string) *DcimManufacturersListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithNameIe(nameIe *string) *DcimManufacturersListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithNameIew(nameIew *string) *DcimManufacturersListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithNameIsw(nameIsw *string) *DcimManufacturersListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithNamen(namen *string) *DcimManufacturersListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithNameNic(nameNic *string) *DcimManufacturersListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithNameNie(nameNie *string) *DcimManufacturersListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithNameNiew(nameNiew *string) *DcimManufacturersListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithNameNisw(nameNisw *string) *DcimManufacturersListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithOffset(offset *int64) *DcimManufacturersListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithQ(q *string) *DcimManufacturersListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlug(slug *string) *DcimManufacturersListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugEmpty adds the slugEmpty to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlugEmpty(slugEmpty *string) *DcimManufacturersListParams {
	o.SetSlugEmpty(slugEmpty)
	return o
}

// SetSlugEmpty adds the slugEmpty to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlugEmpty(slugEmpty *string) {
	o.SlugEmpty = slugEmpty
}

// WithSlugIc adds the slugIc to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlugIc(slugIc *string) *DcimManufacturersListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlugIe(slugIe *string) *DcimManufacturersListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlugIew(slugIew *string) *DcimManufacturersListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlugIsw(slugIsw *string) *DcimManufacturersListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlugn(slugn *string) *DcimManufacturersListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlugNic(slugNic *string) *DcimManufacturersListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlugNie(slugNie *string) *DcimManufacturersListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlugNiew(slugNiew *string) *DcimManufacturersListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) WithSlugNisw(slugNisw *string) *DcimManufacturersListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the dcim manufacturers list params
func (o *DcimManufacturersListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionEmpty != nil {

		// query param description__empty
		var qrDescriptionEmpty string

		if o.DescriptionEmpty != nil {
			qrDescriptionEmpty = *o.DescriptionEmpty
		}
		qDescriptionEmpty := qrDescriptionEmpty
		if qDescriptionEmpty != "" {

			if err := r.SetQueryParam("description__empty", qDescriptionEmpty); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
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

	if o.NameEmpty != nil {

		// query param name__empty
		var qrNameEmpty string

		if o.NameEmpty != nil {
			qrNameEmpty = *o.NameEmpty
		}
		qNameEmpty := qrNameEmpty
		if qNameEmpty != "" {

			if err := r.SetQueryParam("name__empty", qNameEmpty); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
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

	if o.SlugEmpty != nil {

		// query param slug__empty
		var qrSlugEmpty string

		if o.SlugEmpty != nil {
			qrSlugEmpty = *o.SlugEmpty
		}
		qSlugEmpty := qrSlugEmpty
		if qSlugEmpty != "" {

			if err := r.SetQueryParam("slug__empty", qSlugEmpty); err != nil {
				return err
			}
		}
	}

	if o.SlugIc != nil {

		// query param slug__ic
		var qrSlugIc string

		if o.SlugIc != nil {
			qrSlugIc = *o.SlugIc
		}
		qSlugIc := qrSlugIc
		if qSlugIc != "" {

			if err := r.SetQueryParam("slug__ic", qSlugIc); err != nil {
				return err
			}
		}
	}

	if o.SlugIe != nil {

		// query param slug__ie
		var qrSlugIe string

		if o.SlugIe != nil {
			qrSlugIe = *o.SlugIe
		}
		qSlugIe := qrSlugIe
		if qSlugIe != "" {

			if err := r.SetQueryParam("slug__ie", qSlugIe); err != nil {
				return err
			}
		}
	}

	if o.SlugIew != nil {

		// query param slug__iew
		var qrSlugIew string

		if o.SlugIew != nil {
			qrSlugIew = *o.SlugIew
		}
		qSlugIew := qrSlugIew
		if qSlugIew != "" {

			if err := r.SetQueryParam("slug__iew", qSlugIew); err != nil {
				return err
			}
		}
	}

	if o.SlugIsw != nil {

		// query param slug__isw
		var qrSlugIsw string

		if o.SlugIsw != nil {
			qrSlugIsw = *o.SlugIsw
		}
		qSlugIsw := qrSlugIsw
		if qSlugIsw != "" {

			if err := r.SetQueryParam("slug__isw", qSlugIsw); err != nil {
				return err
			}
		}
	}

	if o.Slugn != nil {

		// query param slug__n
		var qrSlugn string

		if o.Slugn != nil {
			qrSlugn = *o.Slugn
		}
		qSlugn := qrSlugn
		if qSlugn != "" {

			if err := r.SetQueryParam("slug__n", qSlugn); err != nil {
				return err
			}
		}
	}

	if o.SlugNic != nil {

		// query param slug__nic
		var qrSlugNic string

		if o.SlugNic != nil {
			qrSlugNic = *o.SlugNic
		}
		qSlugNic := qrSlugNic
		if qSlugNic != "" {

			if err := r.SetQueryParam("slug__nic", qSlugNic); err != nil {
				return err
			}
		}
	}

	if o.SlugNie != nil {

		// query param slug__nie
		var qrSlugNie string

		if o.SlugNie != nil {
			qrSlugNie = *o.SlugNie
		}
		qSlugNie := qrSlugNie
		if qSlugNie != "" {

			if err := r.SetQueryParam("slug__nie", qSlugNie); err != nil {
				return err
			}
		}
	}

	if o.SlugNiew != nil {

		// query param slug__niew
		var qrSlugNiew string

		if o.SlugNiew != nil {
			qrSlugNiew = *o.SlugNiew
		}
		qSlugNiew := qrSlugNiew
		if qSlugNiew != "" {

			if err := r.SetQueryParam("slug__niew", qSlugNiew); err != nil {
				return err
			}
		}
	}

	if o.SlugNisw != nil {

		// query param slug__nisw
		var qrSlugNisw string

		if o.SlugNisw != nil {
			qrSlugNisw = *o.SlugNisw
		}
		qSlugNisw := qrSlugNisw
		if qSlugNisw != "" {

			if err := r.SetQueryParam("slug__nisw", qSlugNisw); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
