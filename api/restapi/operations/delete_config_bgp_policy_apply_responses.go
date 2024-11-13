// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// DeleteConfigBgpPolicyApplyNoContentCode is the HTTP code returned for type DeleteConfigBgpPolicyApplyNoContent
const DeleteConfigBgpPolicyApplyNoContentCode int = 204

/*
DeleteConfigBgpPolicyApplyNoContent OK

swagger:response deleteConfigBgpPolicyApplyNoContent
*/
type DeleteConfigBgpPolicyApplyNoContent struct {
}

// NewDeleteConfigBgpPolicyApplyNoContent creates DeleteConfigBgpPolicyApplyNoContent with default headers values
func NewDeleteConfigBgpPolicyApplyNoContent() *DeleteConfigBgpPolicyApplyNoContent {

	return &DeleteConfigBgpPolicyApplyNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyApplyNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigBgpPolicyApplyBadRequestCode is the HTTP code returned for type DeleteConfigBgpPolicyApplyBadRequest
const DeleteConfigBgpPolicyApplyBadRequestCode int = 400

/*
DeleteConfigBgpPolicyApplyBadRequest Malformed arguments for API call

swagger:response deleteConfigBgpPolicyApplyBadRequest
*/
type DeleteConfigBgpPolicyApplyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyApplyBadRequest creates DeleteConfigBgpPolicyApplyBadRequest with default headers values
func NewDeleteConfigBgpPolicyApplyBadRequest() *DeleteConfigBgpPolicyApplyBadRequest {

	return &DeleteConfigBgpPolicyApplyBadRequest{}
}

// WithPayload adds the payload to the delete config bgp policy apply bad request response
func (o *DeleteConfigBgpPolicyApplyBadRequest) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyApplyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy apply bad request response
func (o *DeleteConfigBgpPolicyApplyBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyApplyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyApplyUnauthorizedCode is the HTTP code returned for type DeleteConfigBgpPolicyApplyUnauthorized
const DeleteConfigBgpPolicyApplyUnauthorizedCode int = 401

/*
DeleteConfigBgpPolicyApplyUnauthorized Invalid authentication credentials

swagger:response deleteConfigBgpPolicyApplyUnauthorized
*/
type DeleteConfigBgpPolicyApplyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyApplyUnauthorized creates DeleteConfigBgpPolicyApplyUnauthorized with default headers values
func NewDeleteConfigBgpPolicyApplyUnauthorized() *DeleteConfigBgpPolicyApplyUnauthorized {

	return &DeleteConfigBgpPolicyApplyUnauthorized{}
}

// WithPayload adds the payload to the delete config bgp policy apply unauthorized response
func (o *DeleteConfigBgpPolicyApplyUnauthorized) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyApplyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy apply unauthorized response
func (o *DeleteConfigBgpPolicyApplyUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyApplyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyApplyForbiddenCode is the HTTP code returned for type DeleteConfigBgpPolicyApplyForbidden
const DeleteConfigBgpPolicyApplyForbiddenCode int = 403

/*
DeleteConfigBgpPolicyApplyForbidden Capacity insufficient

swagger:response deleteConfigBgpPolicyApplyForbidden
*/
type DeleteConfigBgpPolicyApplyForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyApplyForbidden creates DeleteConfigBgpPolicyApplyForbidden with default headers values
func NewDeleteConfigBgpPolicyApplyForbidden() *DeleteConfigBgpPolicyApplyForbidden {

	return &DeleteConfigBgpPolicyApplyForbidden{}
}

// WithPayload adds the payload to the delete config bgp policy apply forbidden response
func (o *DeleteConfigBgpPolicyApplyForbidden) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyApplyForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy apply forbidden response
func (o *DeleteConfigBgpPolicyApplyForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyApplyForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyApplyNotFoundCode is the HTTP code returned for type DeleteConfigBgpPolicyApplyNotFound
const DeleteConfigBgpPolicyApplyNotFoundCode int = 404

/*
DeleteConfigBgpPolicyApplyNotFound Resource not found

swagger:response deleteConfigBgpPolicyApplyNotFound
*/
type DeleteConfigBgpPolicyApplyNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyApplyNotFound creates DeleteConfigBgpPolicyApplyNotFound with default headers values
func NewDeleteConfigBgpPolicyApplyNotFound() *DeleteConfigBgpPolicyApplyNotFound {

	return &DeleteConfigBgpPolicyApplyNotFound{}
}

// WithPayload adds the payload to the delete config bgp policy apply not found response
func (o *DeleteConfigBgpPolicyApplyNotFound) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyApplyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy apply not found response
func (o *DeleteConfigBgpPolicyApplyNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyApplyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyApplyConflictCode is the HTTP code returned for type DeleteConfigBgpPolicyApplyConflict
const DeleteConfigBgpPolicyApplyConflictCode int = 409

/*
DeleteConfigBgpPolicyApplyConflict Resource Conflict. Neigh already exists

swagger:response deleteConfigBgpPolicyApplyConflict
*/
type DeleteConfigBgpPolicyApplyConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyApplyConflict creates DeleteConfigBgpPolicyApplyConflict with default headers values
func NewDeleteConfigBgpPolicyApplyConflict() *DeleteConfigBgpPolicyApplyConflict {

	return &DeleteConfigBgpPolicyApplyConflict{}
}

// WithPayload adds the payload to the delete config bgp policy apply conflict response
func (o *DeleteConfigBgpPolicyApplyConflict) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyApplyConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy apply conflict response
func (o *DeleteConfigBgpPolicyApplyConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyApplyConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyApplyInternalServerErrorCode is the HTTP code returned for type DeleteConfigBgpPolicyApplyInternalServerError
const DeleteConfigBgpPolicyApplyInternalServerErrorCode int = 500

/*
DeleteConfigBgpPolicyApplyInternalServerError Internal service error

swagger:response deleteConfigBgpPolicyApplyInternalServerError
*/
type DeleteConfigBgpPolicyApplyInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyApplyInternalServerError creates DeleteConfigBgpPolicyApplyInternalServerError with default headers values
func NewDeleteConfigBgpPolicyApplyInternalServerError() *DeleteConfigBgpPolicyApplyInternalServerError {

	return &DeleteConfigBgpPolicyApplyInternalServerError{}
}

// WithPayload adds the payload to the delete config bgp policy apply internal server error response
func (o *DeleteConfigBgpPolicyApplyInternalServerError) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyApplyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy apply internal server error response
func (o *DeleteConfigBgpPolicyApplyInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyApplyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyApplyServiceUnavailableCode is the HTTP code returned for type DeleteConfigBgpPolicyApplyServiceUnavailable
const DeleteConfigBgpPolicyApplyServiceUnavailableCode int = 503

/*
DeleteConfigBgpPolicyApplyServiceUnavailable Maintenance mode

swagger:response deleteConfigBgpPolicyApplyServiceUnavailable
*/
type DeleteConfigBgpPolicyApplyServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyApplyServiceUnavailable creates DeleteConfigBgpPolicyApplyServiceUnavailable with default headers values
func NewDeleteConfigBgpPolicyApplyServiceUnavailable() *DeleteConfigBgpPolicyApplyServiceUnavailable {

	return &DeleteConfigBgpPolicyApplyServiceUnavailable{}
}

// WithPayload adds the payload to the delete config bgp policy apply service unavailable response
func (o *DeleteConfigBgpPolicyApplyServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyApplyServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy apply service unavailable response
func (o *DeleteConfigBgpPolicyApplyServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyApplyServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
