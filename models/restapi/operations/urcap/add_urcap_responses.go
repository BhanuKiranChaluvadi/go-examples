// Code generated by go-swagger; DO NOT EDIT.

package urcap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"universal-robots.com/urservice/models"
)

// AddUrcapCreatedCode is the HTTP code returned for type AddUrcapCreated
const AddUrcapCreatedCode int = 201

/*
AddUrcapCreated metadata

swagger:response addUrcapCreated
*/
type AddUrcapCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Manifest `json:"body,omitempty"`
}

// NewAddUrcapCreated creates AddUrcapCreated with default headers values
func NewAddUrcapCreated() *AddUrcapCreated {

	return &AddUrcapCreated{}
}

// WithPayload adds the payload to the add urcap created response
func (o *AddUrcapCreated) WithPayload(payload *models.Manifest) *AddUrcapCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add urcap created response
func (o *AddUrcapCreated) SetPayload(payload *models.Manifest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUrcapCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *AddUrcapCreated) AddUrcapResponder() {}

/*
AddUrcapDefault error response

swagger:response addUrcapDefault
*/
type AddUrcapDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddUrcapDefault creates AddUrcapDefault with default headers values
func NewAddUrcapDefault(code int) *AddUrcapDefault {
	if code <= 0 {
		code = 500
	}

	return &AddUrcapDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add urcap default response
func (o *AddUrcapDefault) WithStatusCode(code int) *AddUrcapDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add urcap default response
func (o *AddUrcapDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add urcap default response
func (o *AddUrcapDefault) WithPayload(payload *models.ErrorResponse) *AddUrcapDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add urcap default response
func (o *AddUrcapDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUrcapDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *AddUrcapDefault) AddUrcapResponder() {}

type AddUrcapNotImplementedResponder struct {
	middleware.Responder
}

func (*AddUrcapNotImplementedResponder) AddUrcapResponder() {}

func AddUrcapNotImplemented() AddUrcapResponder {
	return &AddUrcapNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.AddUrcap has not yet been implemented",
		),
	}
}

type AddUrcapResponder interface {
	middleware.Responder
	AddUrcapResponder()
}