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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// WirelessWirelessLansPartialUpdateReader is a Reader for the WirelessWirelessLansPartialUpdate structure.
type WirelessWirelessLansPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLansPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLansPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLansPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLansPartialUpdateOK creates a WirelessWirelessLansPartialUpdateOK with default headers values
func NewWirelessWirelessLansPartialUpdateOK() *WirelessWirelessLansPartialUpdateOK {
	return &WirelessWirelessLansPartialUpdateOK{}
}

/* WirelessWirelessLansPartialUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLansPartialUpdateOK wireless wireless lans partial update o k
*/
type WirelessWirelessLansPartialUpdateOK struct {
	Payload *models.WirelessLAN
}

func (o *WirelessWirelessLansPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /wireless/wireless-lans/{id}/][%d] wirelessWirelessLansPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *WirelessWirelessLansPartialUpdateOK) GetPayload() *models.WirelessLAN {
	return o.Payload
}

func (o *WirelessWirelessLansPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLansPartialUpdateDefault creates a WirelessWirelessLansPartialUpdateDefault with default headers values
func NewWirelessWirelessLansPartialUpdateDefault(code int) *WirelessWirelessLansPartialUpdateDefault {
	return &WirelessWirelessLansPartialUpdateDefault{
		_statusCode: code,
	}
}

/* WirelessWirelessLansPartialUpdateDefault describes a response with status code -1, with default header values.

WirelessWirelessLansPartialUpdateDefault wireless wireless lans partial update default
*/
type WirelessWirelessLansPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the wireless wireless lans partial update default response
func (o *WirelessWirelessLansPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLansPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /wireless/wireless-lans/{id}/][%d] wireless_wireless-lans_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *WirelessWirelessLansPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLansPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
