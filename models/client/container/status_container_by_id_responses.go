// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"universal-robots.com/urservice/models"
)

// StatusContainerByIDReader is a Reader for the StatusContainerByID structure.
type StatusContainerByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusContainerByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusContainerByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatusContainerByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatusContainerByIDOK creates a StatusContainerByIDOK with default headers values
func NewStatusContainerByIDOK() *StatusContainerByIDOK {
	return &StatusContainerByIDOK{}
}

/*
StatusContainerByIDOK describes a response with status code 200, with default header values.

OK
*/
type StatusContainerByIDOK struct {
}

// IsSuccess returns true when this status container by Id o k response has a 2xx status code
func (o *StatusContainerByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status container by Id o k response has a 3xx status code
func (o *StatusContainerByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status container by Id o k response has a 4xx status code
func (o *StatusContainerByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status container by Id o k response has a 5xx status code
func (o *StatusContainerByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status container by Id o k response a status code equal to that given
func (o *StatusContainerByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *StatusContainerByIDOK) Error() string {
	return fmt.Sprintf("[PUT /containers/{vendorID}/{urcapID}/{artifactName}/status][%d] statusContainerByIdOK ", 200)
}

func (o *StatusContainerByIDOK) String() string {
	return fmt.Sprintf("[PUT /containers/{vendorID}/{urcapID}/{artifactName}/status][%d] statusContainerByIdOK ", 200)
}

func (o *StatusContainerByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusContainerByIDDefault creates a StatusContainerByIDDefault with default headers values
func NewStatusContainerByIDDefault(code int) *StatusContainerByIDDefault {
	return &StatusContainerByIDDefault{
		_statusCode: code,
	}
}

/*
StatusContainerByIDDefault describes a response with status code -1, with default header values.

error
*/
type StatusContainerByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the status container by Id default response
func (o *StatusContainerByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this status container by Id default response has a 2xx status code
func (o *StatusContainerByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this status container by Id default response has a 3xx status code
func (o *StatusContainerByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this status container by Id default response has a 4xx status code
func (o *StatusContainerByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this status container by Id default response has a 5xx status code
func (o *StatusContainerByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this status container by Id default response a status code equal to that given
func (o *StatusContainerByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *StatusContainerByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /containers/{vendorID}/{urcapID}/{artifactName}/status][%d] statusContainerById default  %+v", o._statusCode, o.Payload)
}

func (o *StatusContainerByIDDefault) String() string {
	return fmt.Sprintf("[PUT /containers/{vendorID}/{urcapID}/{artifactName}/status][%d] statusContainerById default  %+v", o._statusCode, o.Payload)
}

func (o *StatusContainerByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StatusContainerByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
