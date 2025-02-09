// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetIpamStatusReader is a Reader for the GetIpamStatus structure.
type GetIpamStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIpamStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIpamStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetIpamStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIpamStatusOK creates a GetIpamStatusOK with default headers values
func NewGetIpamStatusOK() *GetIpamStatusOK {
	return &GetIpamStatusOK{}
}

/*
GetIpamStatusOK describes a response with status code 200, with default header values.

Success
*/
type GetIpamStatusOK struct {
}

func (o *GetIpamStatusOK) Error() string {
	return fmt.Sprintf("[GET /ipam/status][%d] getIpamStatusOK ", 200)
}

func (o *GetIpamStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIpamStatusInternalServerError creates a GetIpamStatusInternalServerError with default headers values
func NewGetIpamStatusInternalServerError() *GetIpamStatusInternalServerError {
	return &GetIpamStatusInternalServerError{}
}

/*
GetIpamStatusInternalServerError describes a response with status code 500, with default header values.

Get ipam status failure
*/
type GetIpamStatusInternalServerError struct {
}

func (o *GetIpamStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ipam/status][%d] getIpamStatusInternalServerError ", 500)
}

func (o *GetIpamStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
