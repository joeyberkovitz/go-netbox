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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimInventoryItemRolesCreateReader is a Reader for the DcimInventoryItemRolesCreate structure.
type DcimInventoryItemRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimInventoryItemRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemRolesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemRolesCreateCreated creates a DcimInventoryItemRolesCreateCreated with default headers values
func NewDcimInventoryItemRolesCreateCreated() *DcimInventoryItemRolesCreateCreated {
	return &DcimInventoryItemRolesCreateCreated{}
}

/* DcimInventoryItemRolesCreateCreated describes a response with status code 201, with default header values.

DcimInventoryItemRolesCreateCreated dcim inventory item roles create created
*/
type DcimInventoryItemRolesCreateCreated struct {
	Payload *models.InventoryItemRole
}

func (o *DcimInventoryItemRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/inventory-item-roles/][%d] dcimInventoryItemRolesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimInventoryItemRolesCreateCreated) GetPayload() *models.InventoryItemRole {
	return o.Payload
}

func (o *DcimInventoryItemRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemRolesCreateDefault creates a DcimInventoryItemRolesCreateDefault with default headers values
func NewDcimInventoryItemRolesCreateDefault(code int) *DcimInventoryItemRolesCreateDefault {
	return &DcimInventoryItemRolesCreateDefault{
		_statusCode: code,
	}
}

/* DcimInventoryItemRolesCreateDefault describes a response with status code -1, with default header values.

DcimInventoryItemRolesCreateDefault dcim inventory item roles create default
*/
type DcimInventoryItemRolesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item roles create default response
func (o *DcimInventoryItemRolesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemRolesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/inventory-item-roles/][%d] dcim_inventory-item-roles_create default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInventoryItemRolesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemRolesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
