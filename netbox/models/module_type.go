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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModuleType module type
//
// swagger:model ModuleType
type ModuleType struct {

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// manufacturer
	// Required: true
	Manufacturer *NestedManufacturer `json:"manufacturer"`

	// Model
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Model *string `json:"model"`

	// Part number
	//
	// Discrete part number (optional)
	// Max Length: 50
	PartNumber string `json:"part_number,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Weight
	Weight *float64 `json:"weight,omitempty"`

	// weight unit
	WeightUnit *ModuleTypeWeightUnit `json:"weight_unit,omitempty"`
}

// Validate validates this module type
func (m *ModuleType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModuleType) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) validateManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer", "body", m.Manufacturer); err != nil {
		return err
	}

	if m.Manufacturer != nil {
		if err := m.Manufacturer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *ModuleType) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if err := validate.MinLength("model", "body", *m.Model, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("model", "body", *m.Model, 100); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) validatePartNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.PartNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("part_number", "body", m.PartNumber, 50); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModuleType) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) validateWeightUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.WeightUnit) { // not required
		return nil
	}

	if m.WeightUnit != nil {
		if err := m.WeightUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight_unit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this module type based on the context it is used
func (m *ModuleType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManufacturer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeightUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModuleType) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) contextValidateManufacturer(ctx context.Context, formats strfmt.Registry) error {

	if m.Manufacturer != nil {
		if err := m.Manufacturer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *ModuleType) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModuleType) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *ModuleType) contextValidateWeightUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.WeightUnit != nil {
		if err := m.WeightUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight_unit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModuleType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModuleType) UnmarshalBinary(b []byte) error {
	var res ModuleType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ModuleTypeWeightUnit Weight unit
//
// swagger:model ModuleTypeWeightUnit
type ModuleTypeWeightUnit struct {

	// label
	// Required: true
	// Enum: [Kilograms Grams Pounds Ounces]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [kg g lb oz]
	Value *string `json:"value"`
}

// Validate validates this module type weight unit
func (m *ModuleTypeWeightUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var moduleTypeWeightUnitTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Kilograms","Grams","Pounds","Ounces"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moduleTypeWeightUnitTypeLabelPropEnum = append(moduleTypeWeightUnitTypeLabelPropEnum, v)
	}
}

const (

	// ModuleTypeWeightUnitLabelKilograms captures enum value "Kilograms"
	ModuleTypeWeightUnitLabelKilograms string = "Kilograms"

	// ModuleTypeWeightUnitLabelGrams captures enum value "Grams"
	ModuleTypeWeightUnitLabelGrams string = "Grams"

	// ModuleTypeWeightUnitLabelPounds captures enum value "Pounds"
	ModuleTypeWeightUnitLabelPounds string = "Pounds"

	// ModuleTypeWeightUnitLabelOunces captures enum value "Ounces"
	ModuleTypeWeightUnitLabelOunces string = "Ounces"
)

// prop value enum
func (m *ModuleTypeWeightUnit) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, moduleTypeWeightUnitTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ModuleTypeWeightUnit) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("weight_unit"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("weight_unit"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var moduleTypeWeightUnitTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kg","g","lb","oz"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moduleTypeWeightUnitTypeValuePropEnum = append(moduleTypeWeightUnitTypeValuePropEnum, v)
	}
}

const (

	// ModuleTypeWeightUnitValueKg captures enum value "kg"
	ModuleTypeWeightUnitValueKg string = "kg"

	// ModuleTypeWeightUnitValueG captures enum value "g"
	ModuleTypeWeightUnitValueG string = "g"

	// ModuleTypeWeightUnitValueLb captures enum value "lb"
	ModuleTypeWeightUnitValueLb string = "lb"

	// ModuleTypeWeightUnitValueOz captures enum value "oz"
	ModuleTypeWeightUnitValueOz string = "oz"
)

// prop value enum
func (m *ModuleTypeWeightUnit) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, moduleTypeWeightUnitTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ModuleTypeWeightUnit) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("weight_unit"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("weight_unit"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this module type weight unit based on context it is used
func (m *ModuleTypeWeightUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModuleTypeWeightUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModuleTypeWeightUnit) UnmarshalBinary(b []byte) error {
	var res ModuleTypeWeightUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
