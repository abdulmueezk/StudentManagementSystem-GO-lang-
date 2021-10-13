// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"crud/gen/models"
)

// GetstudentByNameOKCode is the HTTP code returned for type GetstudentByNameOK
const GetstudentByNameOKCode int = 200

/*GetstudentByNameOK successful operation

swagger:response getstudentByNameOK
*/
type GetstudentByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Student `json:"body,omitempty"`
}

// NewGetstudentByNameOK creates GetstudentByNameOK with default headers values
func NewGetstudentByNameOK() *GetstudentByNameOK {

	return &GetstudentByNameOK{}
}

// WithPayload adds the payload to the getstudent by name o k response
func (o *GetstudentByNameOK) WithPayload(payload *models.Student) *GetstudentByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the getstudent by name o k response
func (o *GetstudentByNameOK) SetPayload(payload *models.Student) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetstudentByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetstudentByNameBadRequestCode is the HTTP code returned for type GetstudentByNameBadRequest
const GetstudentByNameBadRequestCode int = 400

/*GetstudentByNameBadRequest Invalid email supplied

swagger:response getstudentByNameBadRequest
*/
type GetstudentByNameBadRequest struct {
}

// NewGetstudentByNameBadRequest creates GetstudentByNameBadRequest with default headers values
func NewGetstudentByNameBadRequest() *GetstudentByNameBadRequest {

	return &GetstudentByNameBadRequest{}
}

// WriteResponse to the client
func (o *GetstudentByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetstudentByNameNotFoundCode is the HTTP code returned for type GetstudentByNameNotFound
const GetstudentByNameNotFoundCode int = 404

/*GetstudentByNameNotFound email not found

swagger:response getstudentByNameNotFound
*/
type GetstudentByNameNotFound struct {
}

// NewGetstudentByNameNotFound creates GetstudentByNameNotFound with default headers values
func NewGetstudentByNameNotFound() *GetstudentByNameNotFound {

	return &GetstudentByNameNotFound{}
}

// WriteResponse to the client
func (o *GetstudentByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}