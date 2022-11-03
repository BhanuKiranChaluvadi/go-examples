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

// GetUrcapByIDReader is a Reader for the GetUrcapByID structure.
type GetUrcapByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUrcapByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUrcapByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUrcapByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUrcapByIDOK creates a GetUrcapByIDOK with default headers values
func NewGetUrcapByIDOK() *GetUrcapByIDOK {
	return &GetUrcapByIDOK{}
}

/*
GetUrcapByIDOK describes a response with status code 200, with default header values.

A single urcap.
*/
type GetUrcapByIDOK struct {
	Payload *models.Manifest
}

// IsSuccess returns true when this get urcap by Id o k response has a 2xx status code
func (o *GetUrcapByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get urcap by Id o k response has a 3xx status code
func (o *GetUrcapByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get urcap by Id o k response has a 4xx status code
func (o *GetUrcapByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get urcap by Id o k response has a 5xx status code
func (o *GetUrcapByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get urcap by Id o k response a status code equal to that given
func (o *GetUrcapByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUrcapByIDOK) Error() string {
	return fmt.Sprintf("[GET /urcaps/{vendorID}/{urcapID}][%d] getUrcapByIdOK  %+v", 200, o.Payload)
}

func (o *GetUrcapByIDOK) String() string {
	return fmt.Sprintf("[GET /urcaps/{vendorID}/{urcapID}][%d] getUrcapByIdOK  %+v", 200, o.Payload)
}

func (o *GetUrcapByIDOK) GetPayload() *models.Manifest {
	return o.Payload
}

func (o *GetUrcapByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUrcapByIDDefault creates a GetUrcapByIDDefault with default headers values
func NewGetUrcapByIDDefault(code int) *GetUrcapByIDDefault {
	return &GetUrcapByIDDefault{
		_statusCode: code,
	}
}

/*
GetUrcapByIDDefault describes a response with status code -1, with default header values.

api error
*/
type GetUrcapByIDDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the get urcap by Id default response
func (o *GetUrcapByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get urcap by Id default response has a 2xx status code
func (o *GetUrcapByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get urcap by Id default response has a 3xx status code
func (o *GetUrcapByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get urcap by Id default response has a 4xx status code
func (o *GetUrcapByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get urcap by Id default response has a 5xx status code
func (o *GetUrcapByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get urcap by Id default response a status code equal to that given
func (o *GetUrcapByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetUrcapByIDDefault) Error() string {
	return fmt.Sprintf("[GET /urcaps/{vendorID}/{urcapID}][%d] getUrcapById default  %+v", o._statusCode, o.Payload)
}

func (o *GetUrcapByIDDefault) String() string {
	return fmt.Sprintf("[GET /urcaps/{vendorID}/{urcapID}][%d] getUrcapById default  %+v", o._statusCode, o.Payload)
}

func (o *GetUrcapByIDDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUrcapByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
