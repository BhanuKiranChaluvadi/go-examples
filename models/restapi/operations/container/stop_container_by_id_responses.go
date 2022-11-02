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

// StopContainerByIDOKCode is the HTTP code returned for type StopContainerByIDOK
const StopContainerByIDOKCode int = 200

/*
StopContainerByIDOK OK

swagger:response stopContainerByIdOK
*/
type StopContainerByIDOK struct {
}

// NewStopContainerByIDOK creates StopContainerByIDOK with default headers values
func NewStopContainerByIDOK() *StopContainerByIDOK {

	return &StopContainerByIDOK{}
}

// WriteResponse to the client
func (o *StopContainerByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

func (o *StopContainerByIDOK) StopContainerByIDResponder() {}

/*
StopContainerByIDDefault error

swagger:response stopContainerByIdDefault
*/
type StopContainerByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewStopContainerByIDDefault creates StopContainerByIDDefault with default headers values
func NewStopContainerByIDDefault(code int) *StopContainerByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &StopContainerByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the stop container by Id default response
func (o *StopContainerByIDDefault) WithStatusCode(code int) *StopContainerByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the stop container by Id default response
func (o *StopContainerByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the stop container by Id default response
func (o *StopContainerByIDDefault) WithPayload(payload *models.ErrorResponse) *StopContainerByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop container by Id default response
func (o *StopContainerByIDDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopContainerByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *StopContainerByIDDefault) StopContainerByIDResponder() {}

type StopContainerByIDNotImplementedResponder struct {
	middleware.Responder
}

func (*StopContainerByIDNotImplementedResponder) StopContainerByIDResponder() {}

func StopContainerByIDNotImplemented() StopContainerByIDResponder {
	return &StopContainerByIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.StopContainerByID has not yet been implemented",
		),
	}
}

type StopContainerByIDResponder interface {
	middleware.Responder
	StopContainerByIDResponder()
}
