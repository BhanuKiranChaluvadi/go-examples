// Code generated by go-swagger; DO NOT EDIT.

package urcap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"swagger-example/models"
)

// AddUrcapReader is a Reader for the AddUrcap structure.
type AddUrcapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUrcapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddUrcapCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddUrcapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddUrcapCreated creates a AddUrcapCreated with default headers values
func NewAddUrcapCreated() *AddUrcapCreated {
	return &AddUrcapCreated{}
}

/*
AddUrcapCreated describes a response with status code 201, with default header values.

metadata
*/
type AddUrcapCreated struct {
	Payload *models.Manifest
}

// IsSuccess returns true when this add urcap created response has a 2xx status code
func (o *AddUrcapCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add urcap created response has a 3xx status code
func (o *AddUrcapCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add urcap created response has a 4xx status code
func (o *AddUrcapCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add urcap created response has a 5xx status code
func (o *AddUrcapCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add urcap created response a status code equal to that given
func (o *AddUrcapCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddUrcapCreated) Error() string {
	return fmt.Sprintf("[POST /urcaps][%d] addUrcapCreated  %+v", 201, o.Payload)
}

func (o *AddUrcapCreated) String() string {
	return fmt.Sprintf("[POST /urcaps][%d] addUrcapCreated  %+v", 201, o.Payload)
}

func (o *AddUrcapCreated) GetPayload() *models.Manifest {
	return o.Payload
}

func (o *AddUrcapCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUrcapDefault creates a AddUrcapDefault with default headers values
func NewAddUrcapDefault(code int) *AddUrcapDefault {
	return &AddUrcapDefault{
		_statusCode: code,
	}
}

/*
AddUrcapDefault describes a response with status code -1, with default header values.

error response
*/
type AddUrcapDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the add urcap default response
func (o *AddUrcapDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add urcap default response has a 2xx status code
func (o *AddUrcapDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add urcap default response has a 3xx status code
func (o *AddUrcapDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add urcap default response has a 4xx status code
func (o *AddUrcapDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add urcap default response has a 5xx status code
func (o *AddUrcapDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add urcap default response a status code equal to that given
func (o *AddUrcapDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddUrcapDefault) Error() string {
	return fmt.Sprintf("[POST /urcaps][%d] addUrcap default  %+v", o._statusCode, o.Payload)
}

func (o *AddUrcapDefault) String() string {
	return fmt.Sprintf("[POST /urcaps][%d] addUrcap default  %+v", o._statusCode, o.Payload)
}

func (o *AddUrcapDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *AddUrcapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
