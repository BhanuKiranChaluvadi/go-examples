// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"universal-robots.com/urservice/models"
)

// StatusContainerByIDOKCode is the HTTP code returned for type StatusContainerByIDOK
const StatusContainerByIDOKCode int = 200

/*
StatusContainerByIDOK OK

swagger:response statusContainerByIdOK
*/
type StatusContainerByIDOK struct {
}

// NewStatusContainerByIDOK creates StatusContainerByIDOK with default headers values
func NewStatusContainerByIDOK() *StatusContainerByIDOK {

	return &StatusContainerByIDOK{}
}

// WriteResponse to the client
func (o *StatusContainerByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

func (o *StatusContainerByIDOK) StatusContainerByIDResponder() {}

/*
StatusContainerByIDDefault error

swagger:response statusContainerByIdDefault
*/
type StatusContainerByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewStatusContainerByIDDefault creates StatusContainerByIDDefault with default headers values
func NewStatusContainerByIDDefault(code int) *StatusContainerByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &StatusContainerByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the status container by Id default response
func (o *StatusContainerByIDDefault) WithStatusCode(code int) *StatusContainerByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the status container by Id default response
func (o *StatusContainerByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the status container by Id default response
func (o *StatusContainerByIDDefault) WithPayload(payload *models.ErrorResponse) *StatusContainerByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the status container by Id default response
func (o *StatusContainerByIDDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StatusContainerByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *StatusContainerByIDDefault) StatusContainerByIDResponder() {}

type StatusContainerByIDNotImplementedResponder struct {
	middleware.Responder
}

func (*StatusContainerByIDNotImplementedResponder) StatusContainerByIDResponder() {}

func StatusContainerByIDNotImplemented() StatusContainerByIDResponder {
	return &StatusContainerByIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.StatusContainerByID has not yet been implemented",
		),
	}
}

type StatusContainerByIDResponder interface {
	middleware.Responder
	StatusContainerByIDResponder()
}
