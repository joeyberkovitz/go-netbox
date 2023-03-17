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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasSavedFiltersCreateReader is a Reader for the ExtrasSavedFiltersCreate structure.
type ExtrasSavedFiltersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasSavedFiltersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasSavedFiltersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasSavedFiltersCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasSavedFiltersCreateCreated creates a ExtrasSavedFiltersCreateCreated with default headers values
func NewExtrasSavedFiltersCreateCreated() *ExtrasSavedFiltersCreateCreated {
	return &ExtrasSavedFiltersCreateCreated{}
}

/*
ExtrasSavedFiltersCreateCreated describes a response with status code 201, with default header values.

ExtrasSavedFiltersCreateCreated extras saved filters create created
*/
type ExtrasSavedFiltersCreateCreated struct {
	Payload *models.SavedFilter
}

// IsSuccess returns true when this extras saved filters create created response has a 2xx status code
func (o *ExtrasSavedFiltersCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras saved filters create created response has a 3xx status code
func (o *ExtrasSavedFiltersCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras saved filters create created response has a 4xx status code
func (o *ExtrasSavedFiltersCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras saved filters create created response has a 5xx status code
func (o *ExtrasSavedFiltersCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this extras saved filters create created response a status code equal to that given
func (o *ExtrasSavedFiltersCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ExtrasSavedFiltersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/saved-filters/][%d] extrasSavedFiltersCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasSavedFiltersCreateCreated) String() string {
	return fmt.Sprintf("[POST /extras/saved-filters/][%d] extrasSavedFiltersCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasSavedFiltersCreateCreated) GetPayload() *models.SavedFilter {
	return o.Payload
}

func (o *ExtrasSavedFiltersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SavedFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasSavedFiltersCreateDefault creates a ExtrasSavedFiltersCreateDefault with default headers values
func NewExtrasSavedFiltersCreateDefault(code int) *ExtrasSavedFiltersCreateDefault {
	return &ExtrasSavedFiltersCreateDefault{
		_statusCode: code,
	}
}

/*
ExtrasSavedFiltersCreateDefault describes a response with status code -1, with default header values.

ExtrasSavedFiltersCreateDefault extras saved filters create default
*/
type ExtrasSavedFiltersCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras saved filters create default response
func (o *ExtrasSavedFiltersCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras saved filters create default response has a 2xx status code
func (o *ExtrasSavedFiltersCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras saved filters create default response has a 3xx status code
func (o *ExtrasSavedFiltersCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras saved filters create default response has a 4xx status code
func (o *ExtrasSavedFiltersCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras saved filters create default response has a 5xx status code
func (o *ExtrasSavedFiltersCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras saved filters create default response a status code equal to that given
func (o *ExtrasSavedFiltersCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasSavedFiltersCreateDefault) Error() string {
	return fmt.Sprintf("[POST /extras/saved-filters/][%d] extras_saved-filters_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasSavedFiltersCreateDefault) String() string {
	return fmt.Sprintf("[POST /extras/saved-filters/][%d] extras_saved-filters_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasSavedFiltersCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasSavedFiltersCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
