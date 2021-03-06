// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// DownloadClusterManifestOKCode is the HTTP code returned for type DownloadClusterManifestOK
const DownloadClusterManifestOKCode int = 200

/*DownloadClusterManifestOK Success.

swagger:response downloadClusterManifestOK
*/
type DownloadClusterManifestOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadClusterManifestOK creates DownloadClusterManifestOK with default headers values
func NewDownloadClusterManifestOK() *DownloadClusterManifestOK {

	return &DownloadClusterManifestOK{}
}

// WithPayload adds the payload to the download cluster manifest o k response
func (o *DownloadClusterManifestOK) WithPayload(payload io.ReadCloser) *DownloadClusterManifestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster manifest o k response
func (o *DownloadClusterManifestOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterManifestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadClusterManifestUnauthorizedCode is the HTTP code returned for type DownloadClusterManifestUnauthorized
const DownloadClusterManifestUnauthorizedCode int = 401

/*DownloadClusterManifestUnauthorized Unauthorized.

swagger:response downloadClusterManifestUnauthorized
*/
type DownloadClusterManifestUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadClusterManifestUnauthorized creates DownloadClusterManifestUnauthorized with default headers values
func NewDownloadClusterManifestUnauthorized() *DownloadClusterManifestUnauthorized {

	return &DownloadClusterManifestUnauthorized{}
}

// WithPayload adds the payload to the download cluster manifest unauthorized response
func (o *DownloadClusterManifestUnauthorized) WithPayload(payload *models.InfraError) *DownloadClusterManifestUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster manifest unauthorized response
func (o *DownloadClusterManifestUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterManifestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterManifestForbiddenCode is the HTTP code returned for type DownloadClusterManifestForbidden
const DownloadClusterManifestForbiddenCode int = 403

/*DownloadClusterManifestForbidden Forbidden.

swagger:response downloadClusterManifestForbidden
*/
type DownloadClusterManifestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadClusterManifestForbidden creates DownloadClusterManifestForbidden with default headers values
func NewDownloadClusterManifestForbidden() *DownloadClusterManifestForbidden {

	return &DownloadClusterManifestForbidden{}
}

// WithPayload adds the payload to the download cluster manifest forbidden response
func (o *DownloadClusterManifestForbidden) WithPayload(payload *models.InfraError) *DownloadClusterManifestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster manifest forbidden response
func (o *DownloadClusterManifestForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterManifestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterManifestNotFoundCode is the HTTP code returned for type DownloadClusterManifestNotFound
const DownloadClusterManifestNotFoundCode int = 404

/*DownloadClusterManifestNotFound Error.

swagger:response downloadClusterManifestNotFound
*/
type DownloadClusterManifestNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterManifestNotFound creates DownloadClusterManifestNotFound with default headers values
func NewDownloadClusterManifestNotFound() *DownloadClusterManifestNotFound {

	return &DownloadClusterManifestNotFound{}
}

// WithPayload adds the payload to the download cluster manifest not found response
func (o *DownloadClusterManifestNotFound) WithPayload(payload *models.Error) *DownloadClusterManifestNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster manifest not found response
func (o *DownloadClusterManifestNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterManifestNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterManifestMethodNotAllowedCode is the HTTP code returned for type DownloadClusterManifestMethodNotAllowed
const DownloadClusterManifestMethodNotAllowedCode int = 405

/*DownloadClusterManifestMethodNotAllowed Method Not Allowed.

swagger:response downloadClusterManifestMethodNotAllowed
*/
type DownloadClusterManifestMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterManifestMethodNotAllowed creates DownloadClusterManifestMethodNotAllowed with default headers values
func NewDownloadClusterManifestMethodNotAllowed() *DownloadClusterManifestMethodNotAllowed {

	return &DownloadClusterManifestMethodNotAllowed{}
}

// WithPayload adds the payload to the download cluster manifest method not allowed response
func (o *DownloadClusterManifestMethodNotAllowed) WithPayload(payload *models.Error) *DownloadClusterManifestMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster manifest method not allowed response
func (o *DownloadClusterManifestMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterManifestMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterManifestConflictCode is the HTTP code returned for type DownloadClusterManifestConflict
const DownloadClusterManifestConflictCode int = 409

/*DownloadClusterManifestConflict Error.

swagger:response downloadClusterManifestConflict
*/
type DownloadClusterManifestConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterManifestConflict creates DownloadClusterManifestConflict with default headers values
func NewDownloadClusterManifestConflict() *DownloadClusterManifestConflict {

	return &DownloadClusterManifestConflict{}
}

// WithPayload adds the payload to the download cluster manifest conflict response
func (o *DownloadClusterManifestConflict) WithPayload(payload *models.Error) *DownloadClusterManifestConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster manifest conflict response
func (o *DownloadClusterManifestConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterManifestConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterManifestInternalServerErrorCode is the HTTP code returned for type DownloadClusterManifestInternalServerError
const DownloadClusterManifestInternalServerErrorCode int = 500

/*DownloadClusterManifestInternalServerError Error.

swagger:response downloadClusterManifestInternalServerError
*/
type DownloadClusterManifestInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterManifestInternalServerError creates DownloadClusterManifestInternalServerError with default headers values
func NewDownloadClusterManifestInternalServerError() *DownloadClusterManifestInternalServerError {

	return &DownloadClusterManifestInternalServerError{}
}

// WithPayload adds the payload to the download cluster manifest internal server error response
func (o *DownloadClusterManifestInternalServerError) WithPayload(payload *models.Error) *DownloadClusterManifestInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster manifest internal server error response
func (o *DownloadClusterManifestInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterManifestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
