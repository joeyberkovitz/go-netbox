// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableRackReservation writable rack reservation
// swagger:model WritableRackReservation
type WritableRackReservation struct {

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Description
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Description *string `json:"description"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Rack
	// Required: true
	Rack *int64 `json:"rack"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// units
	// Required: true
	Units []*int64 `json:"units"`

	// User
	// Required: true
	User *int64 `json:"user"`
}

// Validate validates this writable rack reservation
func (m *WritableRackReservation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRackReservation) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackReservation) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackReservation) validateRack(formats strfmt.Registry) error {

	if err := validate.Required("rack", "body", m.Rack); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackReservation) validateUnits(formats strfmt.Registry) error {

	if err := validate.Required("units", "body", m.Units); err != nil {
		return err
	}

	for i := 0; i < len(m.Units); i++ {
		if swag.IsZero(m.Units[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("units"+"."+strconv.Itoa(i), "body", int64(*m.Units[i]), 0, false); err != nil {
			return err
		}

		if err := validate.MaximumInt("units"+"."+strconv.Itoa(i), "body", int64(*m.Units[i]), 32767, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *WritableRackReservation) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableRackReservation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableRackReservation) UnmarshalBinary(b []byte) error {
	var res WritableRackReservation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}