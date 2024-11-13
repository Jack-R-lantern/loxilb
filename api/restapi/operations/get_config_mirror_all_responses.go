// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigMirrorAllOKCode is the HTTP code returned for type GetConfigMirrorAllOK
const GetConfigMirrorAllOKCode int = 200

/*
GetConfigMirrorAllOK OK

swagger:response getConfigMirrorAllOK
*/
type GetConfigMirrorAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigMirrorAllOKBody `json:"body,omitempty"`
}

// NewGetConfigMirrorAllOK creates GetConfigMirrorAllOK with default headers values
func NewGetConfigMirrorAllOK() *GetConfigMirrorAllOK {

	return &GetConfigMirrorAllOK{}
}

// WithPayload adds the payload to the get config mirror all o k response
func (o *GetConfigMirrorAllOK) WithPayload(payload *GetConfigMirrorAllOKBody) *GetConfigMirrorAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config mirror all o k response
func (o *GetConfigMirrorAllOK) SetPayload(payload *GetConfigMirrorAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigMirrorAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigMirrorAllUnauthorizedCode is the HTTP code returned for type GetConfigMirrorAllUnauthorized
const GetConfigMirrorAllUnauthorizedCode int = 401

/*
GetConfigMirrorAllUnauthorized Invalid authentication credentials

swagger:response getConfigMirrorAllUnauthorized
*/
type GetConfigMirrorAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigMirrorAllUnauthorized creates GetConfigMirrorAllUnauthorized with default headers values
func NewGetConfigMirrorAllUnauthorized() *GetConfigMirrorAllUnauthorized {

	return &GetConfigMirrorAllUnauthorized{}
}

// WithPayload adds the payload to the get config mirror all unauthorized response
func (o *GetConfigMirrorAllUnauthorized) WithPayload(payload *models.Error) *GetConfigMirrorAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config mirror all unauthorized response
func (o *GetConfigMirrorAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigMirrorAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigMirrorAllInternalServerErrorCode is the HTTP code returned for type GetConfigMirrorAllInternalServerError
const GetConfigMirrorAllInternalServerErrorCode int = 500

/*
GetConfigMirrorAllInternalServerError Internal service error

swagger:response getConfigMirrorAllInternalServerError
*/
type GetConfigMirrorAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigMirrorAllInternalServerError creates GetConfigMirrorAllInternalServerError with default headers values
func NewGetConfigMirrorAllInternalServerError() *GetConfigMirrorAllInternalServerError {

	return &GetConfigMirrorAllInternalServerError{}
}

// WithPayload adds the payload to the get config mirror all internal server error response
func (o *GetConfigMirrorAllInternalServerError) WithPayload(payload *models.Error) *GetConfigMirrorAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config mirror all internal server error response
func (o *GetConfigMirrorAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigMirrorAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigMirrorAllServiceUnavailableCode is the HTTP code returned for type GetConfigMirrorAllServiceUnavailable
const GetConfigMirrorAllServiceUnavailableCode int = 503

/*
GetConfigMirrorAllServiceUnavailable Maintenance mode

swagger:response getConfigMirrorAllServiceUnavailable
*/
type GetConfigMirrorAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigMirrorAllServiceUnavailable creates GetConfigMirrorAllServiceUnavailable with default headers values
func NewGetConfigMirrorAllServiceUnavailable() *GetConfigMirrorAllServiceUnavailable {

	return &GetConfigMirrorAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config mirror all service unavailable response
func (o *GetConfigMirrorAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigMirrorAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config mirror all service unavailable response
func (o *GetConfigMirrorAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigMirrorAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
