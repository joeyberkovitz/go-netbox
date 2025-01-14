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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersPermissionsDeleteReader is a Reader for the UsersPermissionsDelete structure.
type UsersPermissionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersPermissionsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersPermissionsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersPermissionsDeleteNoContent creates a UsersPermissionsDeleteNoContent with default headers values
func NewUsersPermissionsDeleteNoContent() *UsersPermissionsDeleteNoContent {
	return &UsersPermissionsDeleteNoContent{}
}

/*
UsersPermissionsDeleteNoContent describes a response with status code 204, with default header values.

UsersPermissionsDeleteNoContent users permissions delete no content
*/
type UsersPermissionsDeleteNoContent struct {
}

// IsSuccess returns true when this users permissions delete no content response has a 2xx status code
func (o *UsersPermissionsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users permissions delete no content response has a 3xx status code
func (o *UsersPermissionsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users permissions delete no content response has a 4xx status code
func (o *UsersPermissionsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this users permissions delete no content response has a 5xx status code
func (o *UsersPermissionsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this users permissions delete no content response a status code equal to that given
func (o *UsersPermissionsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UsersPermissionsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/permissions/{id}/][%d] usersPermissionsDeleteNoContent ", 204)
}

func (o *UsersPermissionsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/permissions/{id}/][%d] usersPermissionsDeleteNoContent ", 204)
}

func (o *UsersPermissionsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersPermissionsDeleteDefault creates a UsersPermissionsDeleteDefault with default headers values
func NewUsersPermissionsDeleteDefault(code int) *UsersPermissionsDeleteDefault {
	return &UsersPermissionsDeleteDefault{
		_statusCode: code,
	}
}

/*
UsersPermissionsDeleteDefault describes a response with status code -1, with default header values.

UsersPermissionsDeleteDefault users permissions delete default
*/
type UsersPermissionsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users permissions delete default response
func (o *UsersPermissionsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users permissions delete default response has a 2xx status code
func (o *UsersPermissionsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users permissions delete default response has a 3xx status code
func (o *UsersPermissionsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users permissions delete default response has a 4xx status code
func (o *UsersPermissionsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users permissions delete default response has a 5xx status code
func (o *UsersPermissionsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users permissions delete default response a status code equal to that given
func (o *UsersPermissionsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersPermissionsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/permissions/{id}/][%d] users_permissions_delete default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /users/permissions/{id}/][%d] users_permissions_delete default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersPermissionsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
