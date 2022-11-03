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

// GetContainerByIDOKCode is the HTTP code returned for type GetContainerByIDOK
const GetContainerByIDOKCode int = 200

/*
GetContainerByIDOK A single container.

swagger:response getContainerByIdOK
*/
type GetContainerByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Container `json:"body,omitempty"`
}

// NewGetContainerByIDOK creates GetContainerByIDOK with default headers values
func NewGetContainerByIDOK() *GetContainerByIDOK {

	return &GetContainerByIDOK{}
}

// WithPayload adds the payload to the get container by Id o k response
func (o *GetContainerByIDOK) WithPayload(payload *models.Container) *GetContainerByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get container by Id o k response
func (o *GetContainerByIDOK) SetPayload(payload *models.Container) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetContainerByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetContainerByIDOK) GetContainerByIDResponder() {}

/*
GetContainerByIDDefault error

swagger:response getContainerByIdDefault
*/
type GetContainerByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetContainerByIDDefault creates GetContainerByIDDefault with default headers values
func NewGetContainerByIDDefault(code int) *GetContainerByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetContainerByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get container by Id default response
func (o *GetContainerByIDDefault) WithStatusCode(code int) *GetContainerByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get container by Id default response
func (o *GetContainerByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get container by Id default response
func (o *GetContainerByIDDefault) WithPayload(payload *models.ErrorResponse) *GetContainerByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get container by Id default response
func (o *GetContainerByIDDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetContainerByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetContainerByIDDefault) GetContainerByIDResponder() {}

type GetContainerByIDNotImplementedResponder struct {
	middleware.Responder
}

func (*GetContainerByIDNotImplementedResponder) GetContainerByIDResponder() {}

func GetContainerByIDNotImplemented() GetContainerByIDResponder {
	return &GetContainerByIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetContainerByID has not yet been implemented",
		),
	}
}

type GetContainerByIDResponder interface {
	middleware.Responder
	GetContainerByIDResponder()
}