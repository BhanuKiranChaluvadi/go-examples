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

// GetUrcapsReader is a Reader for the GetUrcaps structure.
type GetUrcapsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUrcapsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUrcapsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUrcapsOK creates a GetUrcapsOK with default headers values
func NewGetUrcapsOK() *GetUrcapsOK {
	return &GetUrcapsOK{}
}

/*
GetUrcapsOK describes a response with status code 200, with default header values.

OK
*/
type GetUrcapsOK struct {
	Payload []*models.Manifest
}

// IsSuccess returns true when this get urcaps o k response has a 2xx status code
func (o *GetUrcapsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get urcaps o k response has a 3xx status code
func (o *GetUrcapsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get urcaps o k response has a 4xx status code
func (o *GetUrcapsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get urcaps o k response has a 5xx status code
func (o *GetUrcapsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get urcaps o k response a status code equal to that given
func (o *GetUrcapsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUrcapsOK) Error() string {
	return fmt.Sprintf("[GET /urcaps][%d] getUrcapsOK  %+v", 200, o.Payload)
}

func (o *GetUrcapsOK) String() string {
	return fmt.Sprintf("[GET /urcaps][%d] getUrcapsOK  %+v", 200, o.Payload)
}

func (o *GetUrcapsOK) GetPayload() []*models.Manifest {
	return o.Payload
}

func (o *GetUrcapsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
