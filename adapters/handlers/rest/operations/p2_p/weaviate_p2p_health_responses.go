//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package p2_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateP2pHealthOKCode is the HTTP code returned for type WeaviateP2pHealthOK
const WeaviateP2pHealthOKCode int = 200

/*WeaviateP2pHealthOK Alive and kicking!

swagger:response weaviateP2pHealthOK
*/
type WeaviateP2pHealthOK struct {
}

// NewWeaviateP2pHealthOK creates WeaviateP2pHealthOK with default headers values
func NewWeaviateP2pHealthOK() *WeaviateP2pHealthOK {

	return &WeaviateP2pHealthOK{}
}

// WriteResponse to the client
func (o *WeaviateP2pHealthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateP2pHealthInternalServerErrorCode is the HTTP code returned for type WeaviateP2pHealthInternalServerError
const WeaviateP2pHealthInternalServerErrorCode int = 500

/*WeaviateP2pHealthInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateP2pHealthInternalServerError
*/
type WeaviateP2pHealthInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateP2pHealthInternalServerError creates WeaviateP2pHealthInternalServerError with default headers values
func NewWeaviateP2pHealthInternalServerError() *WeaviateP2pHealthInternalServerError {

	return &WeaviateP2pHealthInternalServerError{}
}

// WithPayload adds the payload to the weaviate p2p health internal server error response
func (o *WeaviateP2pHealthInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateP2pHealthInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate p2p health internal server error response
func (o *WeaviateP2pHealthInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateP2pHealthInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
