// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new container API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for container API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetContainerByID(params *GetContainerByIDParams, opts ...ClientOption) (*GetContainerByIDOK, error)

	GetContainerPortMappingByID(params *GetContainerPortMappingByIDParams, opts ...ClientOption) (*GetContainerPortMappingByIDOK, error)

	StartContainerByID(params *StartContainerByIDParams, opts ...ClientOption) (*StartContainerByIDOK, error)

	StatusContainerByID(params *StatusContainerByIDParams, opts ...ClientOption) (*StatusContainerByIDOK, error)

	StopAllContainers(params *StopAllContainersParams, opts ...ClientOption) (*StopAllContainersOK, error)

	StopContainerByID(params *StopContainerByIDParams, opts ...ClientOption) (*StopContainerByIDOK, error)

	ValidateContainerByID(params *ValidateContainerByIDParams, opts ...ClientOption) (*ValidateContainerByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetContainerByID gets a container info by urcap ID s and artifact name

TODO
*/
func (a *Client) GetContainerByID(params *GetContainerByIDParams, opts ...ClientOption) (*GetContainerByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContainerByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getContainerById",
		Method:             "GET",
		PathPattern:        "/containers/{vendorID}/{urcapID}/{artifactName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContainerByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetContainerByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetContainerByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetContainerPortMappingByID ports mapping of a container

Gets host port mapped to the container port
*/
func (a *Client) GetContainerPortMappingByID(params *GetContainerPortMappingByIDParams, opts ...ClientOption) (*GetContainerPortMappingByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContainerPortMappingByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getContainerPortMappingById",
		Method:             "GET",
		PathPattern:        "/containers/{vendorID}/{urcapID}/{artifactName}/portmapping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContainerPortMappingByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetContainerPortMappingByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetContainerPortMappingByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StartContainerByID starts a container

TODO
*/
func (a *Client) StartContainerByID(params *StartContainerByIDParams, opts ...ClientOption) (*StartContainerByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartContainerByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startContainerById",
		Method:             "PUT",
		PathPattern:        "/containers/{vendorID}/{urcapID}/{artifactName}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StartContainerByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartContainerByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartContainerByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StatusContainerByID statuses of a container

TODO
*/
func (a *Client) StatusContainerByID(params *StatusContainerByIDParams, opts ...ClientOption) (*StatusContainerByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusContainerByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusContainerById",
		Method:             "PUT",
		PathPattern:        "/containers/{vendorID}/{urcapID}/{artifactName}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StatusContainerByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatusContainerByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StatusContainerByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StopAllContainers stops all containers

TODO
*/
func (a *Client) StopAllContainers(params *StopAllContainersParams, opts ...ClientOption) (*StopAllContainersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopAllContainersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopAllContainers",
		Method:             "PUT",
		PathPattern:        "/containers/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StopAllContainersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopAllContainersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StopAllContainersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StopContainerByID stops a container

TODO
*/
func (a *Client) StopContainerByID(params *StopContainerByIDParams, opts ...ClientOption) (*StopContainerByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopContainerByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopContainerById",
		Method:             "PUT",
		PathPattern:        "/containers/{vendorID}/{urcapID}/{artifactName}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StopContainerByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopContainerByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StopContainerByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ValidateContainerByID validates container artifact presence

TODO
*/
func (a *Client) ValidateContainerByID(params *ValidateContainerByIDParams, opts ...ClientOption) (*ValidateContainerByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateContainerByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateContainerById",
		Method:             "GET",
		PathPattern:        "/containers/{vendorID}/{urcapID}/{artifactName}/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateContainerByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateContainerByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ValidateContainerByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}