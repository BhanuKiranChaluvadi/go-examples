// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"swagger-example/models"
)

// StopContainerByIDReader is a Reader for the StopContainerByID structure.
type StopContainerByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopContainerByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopContainerByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStopContainerByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStopContainerByIDOK creates a StopContainerByIDOK with default headers values
func NewStopContainerByIDOK() *StopContainerByIDOK {
	return &StopContainerByIDOK{}
}

/*
StopContainerByIDOK describes a response with status code 200, with default header values.

OK
*/
type StopContainerByIDOK struct {
}

// IsSuccess returns true when this stop container by Id o k response has a 2xx status code
func (o *StopContainerByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop container by Id o k response has a 3xx status code
func (o *StopContainerByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop container by Id o k response has a 4xx status code
func (o *StopContainerByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop container by Id o k response has a 5xx status code
func (o *StopContainerByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop container by Id o k response a status code equal to that given
func (o *StopContainerByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *StopContainerByIDOK) Error() string {
	return fmt.Sprintf("[PUT /containers/{vendorID}/{urcapID}/{artifactName}/stop][%d] stopContainerByIdOK ", 200)
}

func (o *StopContainerByIDOK) String() string {
	return fmt.Sprintf("[PUT /containers/{vendorID}/{urcapID}/{artifactName}/stop][%d] stopContainerByIdOK ", 200)
}

func (o *StopContainerByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopContainerByIDDefault creates a StopContainerByIDDefault with default headers values
func NewStopContainerByIDDefault(code int) *StopContainerByIDDefault {
	return &StopContainerByIDDefault{
		_statusCode: code,
	}
}

/*
StopContainerByIDDefault describes a response with status code -1, with default header values.

api error
*/
type StopContainerByIDDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the stop container by Id default response
func (o *StopContainerByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this stop container by Id default response has a 2xx status code
func (o *StopContainerByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this stop container by Id default response has a 3xx status code
func (o *StopContainerByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this stop container by Id default response has a 4xx status code
func (o *StopContainerByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this stop container by Id default response has a 5xx status code
func (o *StopContainerByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this stop container by Id default response a status code equal to that given
func (o *StopContainerByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *StopContainerByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /containers/{vendorID}/{urcapID}/{artifactName}/stop][%d] stopContainerById default  %+v", o._statusCode, o.Payload)
}

func (o *StopContainerByIDDefault) String() string {
	return fmt.Sprintf("[PUT /containers/{vendorID}/{urcapID}/{artifactName}/stop][%d] stopContainerById default  %+v", o._statusCode, o.Payload)
}

func (o *StopContainerByIDDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *StopContainerByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
