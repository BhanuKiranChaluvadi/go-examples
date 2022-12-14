// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"urcapCli/models"
)

// GetContainerPortMappingByIDReader is a Reader for the GetContainerPortMappingByID structure.
type GetContainerPortMappingByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContainerPortMappingByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContainerPortMappingByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetContainerPortMappingByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetContainerPortMappingByIDOK creates a GetContainerPortMappingByIDOK with default headers values
func NewGetContainerPortMappingByIDOK() *GetContainerPortMappingByIDOK {
	return &GetContainerPortMappingByIDOK{}
}

/*
GetContainerPortMappingByIDOK describes a response with status code 200, with default header values.

Container Port Mapping.
*/
type GetContainerPortMappingByIDOK struct {
}

// IsSuccess returns true when this get container port mapping by Id o k response has a 2xx status code
func (o *GetContainerPortMappingByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get container port mapping by Id o k response has a 3xx status code
func (o *GetContainerPortMappingByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get container port mapping by Id o k response has a 4xx status code
func (o *GetContainerPortMappingByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get container port mapping by Id o k response has a 5xx status code
func (o *GetContainerPortMappingByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get container port mapping by Id o k response a status code equal to that given
func (o *GetContainerPortMappingByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetContainerPortMappingByIDOK) Error() string {
	return fmt.Sprintf("[GET /containers/{vendorID}/{urcapID}/{artifactName}/portmapping][%d] getContainerPortMappingByIdOK ", 200)
}

func (o *GetContainerPortMappingByIDOK) String() string {
	return fmt.Sprintf("[GET /containers/{vendorID}/{urcapID}/{artifactName}/portmapping][%d] getContainerPortMappingByIdOK ", 200)
}

func (o *GetContainerPortMappingByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetContainerPortMappingByIDDefault creates a GetContainerPortMappingByIDDefault with default headers values
func NewGetContainerPortMappingByIDDefault(code int) *GetContainerPortMappingByIDDefault {
	return &GetContainerPortMappingByIDDefault{
		_statusCode: code,
	}
}

/*
GetContainerPortMappingByIDDefault describes a response with status code -1, with default header values.

api error
*/
type GetContainerPortMappingByIDDefault struct {
	_statusCode int

	Payload *models.APIError
}

// Code gets the status code for the get container port mapping by Id default response
func (o *GetContainerPortMappingByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get container port mapping by Id default response has a 2xx status code
func (o *GetContainerPortMappingByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get container port mapping by Id default response has a 3xx status code
func (o *GetContainerPortMappingByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get container port mapping by Id default response has a 4xx status code
func (o *GetContainerPortMappingByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get container port mapping by Id default response has a 5xx status code
func (o *GetContainerPortMappingByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get container port mapping by Id default response a status code equal to that given
func (o *GetContainerPortMappingByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetContainerPortMappingByIDDefault) Error() string {
	return fmt.Sprintf("[GET /containers/{vendorID}/{urcapID}/{artifactName}/portmapping][%d] getContainerPortMappingById default  %+v", o._statusCode, o.Payload)
}

func (o *GetContainerPortMappingByIDDefault) String() string {
	return fmt.Sprintf("[GET /containers/{vendorID}/{urcapID}/{artifactName}/portmapping][%d] getContainerPortMappingById default  %+v", o._statusCode, o.Payload)
}

func (o *GetContainerPortMappingByIDDefault) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetContainerPortMappingByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
