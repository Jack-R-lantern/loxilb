// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigBfdAllOKCode is the HTTP code returned for type GetConfigBfdAllOK
const GetConfigBfdAllOKCode int = 200

/*
GetConfigBfdAllOK OK

swagger:response getConfigBfdAllOK
*/
type GetConfigBfdAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigBfdAllOKBody `json:"body,omitempty"`
}

// NewGetConfigBfdAllOK creates GetConfigBfdAllOK with default headers values
func NewGetConfigBfdAllOK() *GetConfigBfdAllOK {

	return &GetConfigBfdAllOK{}
}

// WithPayload adds the payload to the get config bfd all o k response
func (o *GetConfigBfdAllOK) WithPayload(payload *GetConfigBfdAllOKBody) *GetConfigBfdAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bfd all o k response
func (o *GetConfigBfdAllOK) SetPayload(payload *GetConfigBfdAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBfdAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBfdAllUnauthorizedCode is the HTTP code returned for type GetConfigBfdAllUnauthorized
const GetConfigBfdAllUnauthorizedCode int = 401

/*
GetConfigBfdAllUnauthorized Invalid authentication credentials

swagger:response getConfigBfdAllUnauthorized
*/
type GetConfigBfdAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBfdAllUnauthorized creates GetConfigBfdAllUnauthorized with default headers values
func NewGetConfigBfdAllUnauthorized() *GetConfigBfdAllUnauthorized {

	return &GetConfigBfdAllUnauthorized{}
}

// WithPayload adds the payload to the get config bfd all unauthorized response
func (o *GetConfigBfdAllUnauthorized) WithPayload(payload *models.Error) *GetConfigBfdAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bfd all unauthorized response
func (o *GetConfigBfdAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBfdAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBfdAllInternalServerErrorCode is the HTTP code returned for type GetConfigBfdAllInternalServerError
const GetConfigBfdAllInternalServerErrorCode int = 500

/*
GetConfigBfdAllInternalServerError Internal service error

swagger:response getConfigBfdAllInternalServerError
*/
type GetConfigBfdAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBfdAllInternalServerError creates GetConfigBfdAllInternalServerError with default headers values
func NewGetConfigBfdAllInternalServerError() *GetConfigBfdAllInternalServerError {

	return &GetConfigBfdAllInternalServerError{}
}

// WithPayload adds the payload to the get config bfd all internal server error response
func (o *GetConfigBfdAllInternalServerError) WithPayload(payload *models.Error) *GetConfigBfdAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bfd all internal server error response
func (o *GetConfigBfdAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBfdAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBfdAllServiceUnavailableCode is the HTTP code returned for type GetConfigBfdAllServiceUnavailable
const GetConfigBfdAllServiceUnavailableCode int = 503

/*
GetConfigBfdAllServiceUnavailable Maintenance mode

swagger:response getConfigBfdAllServiceUnavailable
*/
type GetConfigBfdAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBfdAllServiceUnavailable creates GetConfigBfdAllServiceUnavailable with default headers values
func NewGetConfigBfdAllServiceUnavailable() *GetConfigBfdAllServiceUnavailable {

	return &GetConfigBfdAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config bfd all service unavailable response
func (o *GetConfigBfdAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigBfdAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bfd all service unavailable response
func (o *GetConfigBfdAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBfdAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
