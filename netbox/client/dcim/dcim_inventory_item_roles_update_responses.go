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

// DcimInventoryItemRolesUpdateReader is a Reader for the DcimInventoryItemRolesUpdate structure.
type DcimInventoryItemRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemRolesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemRolesUpdateOK creates a DcimInventoryItemRolesUpdateOK with default headers values
func NewDcimInventoryItemRolesUpdateOK() *DcimInventoryItemRolesUpdateOK {
	return &DcimInventoryItemRolesUpdateOK{}
}

/* DcimInventoryItemRolesUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemRolesUpdateOK dcim inventory item roles update o k
*/
type DcimInventoryItemRolesUpdateOK struct {
	Payload *models.InventoryItemRole
}

func (o *DcimInventoryItemRolesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-roles/{id}/][%d] dcimInventoryItemRolesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemRolesUpdateOK) GetPayload() *models.InventoryItemRole {
	return o.Payload
}

func (o *DcimInventoryItemRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemRolesUpdateDefault creates a DcimInventoryItemRolesUpdateDefault with default headers values
func NewDcimInventoryItemRolesUpdateDefault(code int) *DcimInventoryItemRolesUpdateDefault {
	return &DcimInventoryItemRolesUpdateDefault{
		_statusCode: code,
	}
}

/* DcimInventoryItemRolesUpdateDefault describes a response with status code -1, with default header values.

DcimInventoryItemRolesUpdateDefault dcim inventory item roles update default
*/
type DcimInventoryItemRolesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item roles update default response
func (o *DcimInventoryItemRolesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemRolesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-roles/{id}/][%d] dcim_inventory-item-roles_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInventoryItemRolesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemRolesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}