// Code generated by go-swagger; DO NOT EDIT.

package urcap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"universal-robots.com/urservice/models"
)

// UpdateUrcapByIDReader is a Reader for the UpdateUrcapByID structure.
type UpdateUrcapByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUrcapByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUrcapByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateUrcapByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateUrcapByIDOK creates a UpdateUrcapByIDOK with default headers values
func NewUpdateUrcapByIDOK() *UpdateUrcapByIDOK {
	return &UpdateUrcapByIDOK{}
}

/*
UpdateUrcapByIDOK describes a response with status code 200, with default header values.

Updated urcap.
*/
type UpdateUrcapByIDOK struct {
	Payload *models.Manifest
}

// IsSuccess returns true when this update urcap by Id o k response has a 2xx status code
func (o *UpdateUrcapByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update urcap by Id o k response has a 3xx status code
func (o *UpdateUrcapByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update urcap by Id o k response has a 4xx status code
func (o *UpdateUrcapByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update urcap by Id o k response has a 5xx status code
func (o *UpdateUrcapByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update urcap by Id o k response a status code equal to that given
func (o *UpdateUrcapByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateUrcapByIDOK) Error() string {
	return fmt.Sprintf("[PUT /urcaps/{vendorID}/{urcapID}][%d] updateUrcapByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateUrcapByIDOK) String() string {
	return fmt.Sprintf("[PUT /urcaps/{vendorID}/{urcapID}][%d] updateUrcapByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateUrcapByIDOK) GetPayload() *models.Manifest {
	return o.Payload
}

func (o *UpdateUrcapByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUrcapByIDDefault creates a UpdateUrcapByIDDefault with default headers values
func NewUpdateUrcapByIDDefault(code int) *UpdateUrcapByIDDefault {
	return &UpdateUrcapByIDDefault{
		_statusCode: code,
	}
}

/*
UpdateUrcapByIDDefault describes a response with status code -1, with default header values.

error
*/
type UpdateUrcapByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update urcap by Id default response
func (o *UpdateUrcapByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update urcap by Id default response has a 2xx status code
func (o *UpdateUrcapByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update urcap by Id default response has a 3xx status code
func (o *UpdateUrcapByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update urcap by Id default response has a 4xx status code
func (o *UpdateUrcapByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update urcap by Id default response has a 5xx status code
func (o *UpdateUrcapByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update urcap by Id default response a status code equal to that given
func (o *UpdateUrcapByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateUrcapByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /urcaps/{vendorID}/{urcapID}][%d] updateUrcapById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateUrcapByIDDefault) String() string {
	return fmt.Sprintf("[PUT /urcaps/{vendorID}/{urcapID}][%d] updateUrcapById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateUrcapByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateUrcapByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
