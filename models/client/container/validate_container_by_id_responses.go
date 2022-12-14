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

// ValidateContainerByIDReader is a Reader for the ValidateContainerByID structure.
type ValidateContainerByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateContainerByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateContainerByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewValidateContainerByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateContainerByIDOK creates a ValidateContainerByIDOK with default headers values
func NewValidateContainerByIDOK() *ValidateContainerByIDOK {
	return &ValidateContainerByIDOK{}
}

/*
ValidateContainerByIDOK describes a response with status code 200, with default header values.

Container Port Mapping.
*/
type ValidateContainerByIDOK struct {
}

// IsSuccess returns true when this validate container by Id o k response has a 2xx status code
func (o *ValidateContainerByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate container by Id o k response has a 3xx status code
func (o *ValidateContainerByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate container by Id o k response has a 4xx status code
func (o *ValidateContainerByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate container by Id o k response has a 5xx status code
func (o *ValidateContainerByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate container by Id o k response a status code equal to that given
func (o *ValidateContainerByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateContainerByIDOK) Error() string {
	return fmt.Sprintf("[GET /containers/{vendorID}/{urcapID}/{artifactName}/validate][%d] validateContainerByIdOK ", 200)
}

func (o *ValidateContainerByIDOK) String() string {
	return fmt.Sprintf("[GET /containers/{vendorID}/{urcapID}/{artifactName}/validate][%d] validateContainerByIdOK ", 200)
}

func (o *ValidateContainerByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateContainerByIDDefault creates a ValidateContainerByIDDefault with default headers values
func NewValidateContainerByIDDefault(code int) *ValidateContainerByIDDefault {
	return &ValidateContainerByIDDefault{
		_statusCode: code,
	}
}

/*
ValidateContainerByIDDefault describes a response with status code -1, with default header values.

error
*/
type ValidateContainerByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the validate container by Id default response
func (o *ValidateContainerByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this validate container by Id default response has a 2xx status code
func (o *ValidateContainerByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this validate container by Id default response has a 3xx status code
func (o *ValidateContainerByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this validate container by Id default response has a 4xx status code
func (o *ValidateContainerByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this validate container by Id default response has a 5xx status code
func (o *ValidateContainerByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this validate container by Id default response a status code equal to that given
func (o *ValidateContainerByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ValidateContainerByIDDefault) Error() string {
	return fmt.Sprintf("[GET /containers/{vendorID}/{urcapID}/{artifactName}/validate][%d] validateContainerById default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateContainerByIDDefault) String() string {
	return fmt.Sprintf("[GET /containers/{vendorID}/{urcapID}/{artifactName}/validate][%d] validateContainerById default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateContainerByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ValidateContainerByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
