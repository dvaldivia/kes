// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO KES
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package k_m_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/kes/models"
)

// KMSListPoliciesOKCode is the HTTP code returned for type KMSListPoliciesOK
const KMSListPoliciesOKCode int = 200

/*
KMSListPoliciesOK A successful response.

swagger:response kMSListPoliciesOK
*/
type KMSListPoliciesOK struct {

	/*
	  In: Body
	*/
	Payload *models.KmsListPoliciesResponse `json:"body,omitempty"`
}

// NewKMSListPoliciesOK creates KMSListPoliciesOK with default headers values
func NewKMSListPoliciesOK() *KMSListPoliciesOK {

	return &KMSListPoliciesOK{}
}

// WithPayload adds the payload to the k m s list policies o k response
func (o *KMSListPoliciesOK) WithPayload(payload *models.KmsListPoliciesResponse) *KMSListPoliciesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s list policies o k response
func (o *KMSListPoliciesOK) SetPayload(payload *models.KmsListPoliciesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSListPoliciesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
KMSListPoliciesDefault Generic error response.

swagger:response kMSListPoliciesDefault
*/
type KMSListPoliciesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewKMSListPoliciesDefault creates KMSListPoliciesDefault with default headers values
func NewKMSListPoliciesDefault(code int) *KMSListPoliciesDefault {
	if code <= 0 {
		code = 500
	}

	return &KMSListPoliciesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the k m s list policies default response
func (o *KMSListPoliciesDefault) WithStatusCode(code int) *KMSListPoliciesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the k m s list policies default response
func (o *KMSListPoliciesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the k m s list policies default response
func (o *KMSListPoliciesDefault) WithPayload(payload *models.Error) *KMSListPoliciesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s list policies default response
func (o *KMSListPoliciesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSListPoliciesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
