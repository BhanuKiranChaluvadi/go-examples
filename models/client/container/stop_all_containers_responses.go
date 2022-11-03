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

// StopAllContainersReader is a Reader for the StopAllContainers structure.
type StopAllContainersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopAllContainersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopAllContainersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStopAllContainersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStopAllContainersOK creates a StopAllContainersOK with default headers values
func NewStopAllContainersOK() *StopAllContainersOK {
	return &StopAllContainersOK{}
}

/*
StopAllContainersOK describes a response with status code 200, with default header values.

Container Port Mapping.
*/
type StopAllContainersOK struct {
}

// IsSuccess returns true when this stop all containers o k response has a 2xx status code
func (o *StopAllContainersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop all containers o k response has a 3xx status code
func (o *StopAllContainersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all containers o k response has a 4xx status code
func (o *StopAllContainersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop all containers o k response has a 5xx status code
func (o *StopAllContainersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all containers o k response a status code equal to that given
func (o *StopAllContainersOK) IsCode(code int) bool {
	return code == 200
}

func (o *StopAllContainersOK) Error() string {
	return fmt.Sprintf("[PUT /containers/stop][%d] stopAllContainersOK ", 200)
}

func (o *StopAllContainersOK) String() string {
	return fmt.Sprintf("[PUT /containers/stop][%d] stopAllContainersOK ", 200)
}

func (o *StopAllContainersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopAllContainersDefault creates a StopAllContainersDefault with default headers values
func NewStopAllContainersDefault(code int) *StopAllContainersDefault {
	return &StopAllContainersDefault{
		_statusCode: code,
	}
}

/*
StopAllContainersDefault describes a response with status code -1, with default header values.

error
*/
type StopAllContainersDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the stop all containers default response
func (o *StopAllContainersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this stop all containers default response has a 2xx status code
func (o *StopAllContainersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this stop all containers default response has a 3xx status code
func (o *StopAllContainersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this stop all containers default response has a 4xx status code
func (o *StopAllContainersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this stop all containers default response has a 5xx status code
func (o *StopAllContainersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this stop all containers default response a status code equal to that given
func (o *StopAllContainersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *StopAllContainersDefault) Error() string {
	return fmt.Sprintf("[PUT /containers/stop][%d] stopAllContainers default  %+v", o._statusCode, o.Payload)
}

func (o *StopAllContainersDefault) String() string {
	return fmt.Sprintf("[PUT /containers/stop][%d] stopAllContainers default  %+v", o._statusCode, o.Payload)
}

func (o *StopAllContainersDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StopAllContainersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}