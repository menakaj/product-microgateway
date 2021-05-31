// Code generated by go-swagger; DO NOT EDIT.

//
// This product is licensed by WSO2 Inc. under Apache License 2.0. The license
// can be downloaded from the following locations:
// 	http://www.apache.org/licenses/LICENSE-2.0.html
// 	http://www.apache.org/licenses/LICENSE-2.0.txt
//
// This product also contains software under different licenses. This table below
// all the contained libraries (jar files) and the license under which they are
// provided to you.
//
// At the bottom of this file is a table that shows what each license indicated
// below is and where the actual text of the license can be found.
//
//
// Dependency								License
// github.com/PuerkitoBio/purell						BSD 3-Clause "New" or "Revised" License
// github.com/PuerkitoBio/urlesc						BSD 3-Clause "New" or "Revised" License
// github.com/asaskevich/govalidator					MIT License
// github.com/census-instrumentation/opencensus-proto			Apache License 2.0
// github.com/cncf/udpa/go							Apache License 2.0
// github.com/decred/dcrd/dcrec/secp256k1					ISC License
// github.com/docker/go-units						Apache License 2.0
// github.com/envoyproxy/go-control-plane					Apache License 2.0
// github.com/envoyproxy/protoc-gen-validate				Apache License 2.0
// github.com/fsnotify/fsnotify						BSD 3-Clause "New" or "Revised" License
// github.com/getkin/kin-openapi						MIT License
// github.com/ghodss/yaml							MIT License
// github.com/go-openapi/analysis						Apache License 2.0
// github.com/go-openapi/errors						Apache License 2.0
// github.com/go-openapi/jsonpointer					Apache License 2.0
// github.com/go-openapi/jsonreference					Apache License 2.0
// github.com/go-openapi/loads						Apache License 2.0
// github.com/go-openapi/runtime						Apache License 2.0
// github.com/go-openapi/spec						Apache License 2.0
// github.com/go-openapi/strfmt						Apache License 2.0
// github.com/go-openapi/swag						Apache License 2.0
// github.com/go-openapi/validate						Apache License 2.0
// github.com/go-stack/stack						MIT License
// github.com/golang/protobuf						BSD 3-Clause "New" or "Revised" License
// github.com/google/uuid							BSD 3-Clause "New" or "Revised" License
// github.com/jessevdk/go-flags						BSD 3-Clause "New" or "Revised" License
// github.com/lestrrat-go/backoff						MIT License
// github.com/lestrrat-go/httpcc						MIT License
// github.com/lestrrat-go/iter						MIT License
// github.com/lestrrat-go/jwx						MIT License
// github.com/lestrrat-go/option						MIT License
// github.com/mailru/easyjson						MIT License
// github.com/mitchellh/mapstructure					MIT License
// github.com/pelletier/go-toml						MIT License
// github.com/pkg/errors							BSD 2-Clause "Simplified" License
// github.com/sirupsen/logrus						MIT License
// github.com/streadway/amqp						BSD 2-Clause "Simplified" License
// go.mongodb.org/mongo-driver						Apache License 2.0
// golang.org/x/crypto							BSD 3-Clause "New" or "Revised" License
// golang.org/x/net							BSD 3-Clause "New" or "Revised" License
// golang.org/x/sys							BSD 3-Clause "New" or "Revised" License
// golang.org/x/text							BSD 3-Clause "New" or "Revised" License
// google.golang.org/genproto						Apache License 2.0
// google.golang.org/grpc							Apache License 2.0
// google.golang.org/protobuf						BSD 3-Clause "New" or "Revised" License
// gopkg.in/natefinch/lumberjack.v2					MIT License
// gopkg.in/yaml.v2							Apache License 2.0
//
//
//
//
// The license types used by the above libraries and their information is given below:
//
// apache2        Apache License Version 2.0
//                http://www.apache.org/licenses/LICENSE-2.0.html
// mit            MIT License
//                http://www.opensource.org/licenses/mit-license.php
// bsd2           Berkeley License - 2
//                https://opensource.org/licenses/BSD-2-Clause
// bsd3           Berkeley License - 3
//                http://opensource.org/licenses/BSD-3-Clause
// isc	       Internet Systems Consortium
// 	       https://opensource.org/licenses/ISC
//
//

package api_individual

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/wso2/adapter/internal/api/models"
)

// DeleteApisOKCode is the HTTP code returned for type DeleteApisOK
const DeleteApisOKCode int = 200

/*DeleteApisOK OK.
API successfully undeployed from the Microgateway.


swagger:response deleteApisOK
*/
type DeleteApisOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeployResponse `json:"body,omitempty"`
}

// NewDeleteApisOK creates DeleteApisOK with default headers values
func NewDeleteApisOK() *DeleteApisOK {

	return &DeleteApisOK{}
}

// WithPayload adds the payload to the delete apis o k response
func (o *DeleteApisOK) WithPayload(payload *models.DeployResponse) *DeleteApisOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apis o k response
func (o *DeleteApisOK) SetPayload(payload *models.DeployResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApisOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteApisBadRequestCode is the HTTP code returned for type DeleteApisBadRequest
const DeleteApisBadRequestCode int = 400

/*DeleteApisBadRequest Bad Request.
Invalid request or validation error


swagger:response deleteApisBadRequest
*/
type DeleteApisBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteApisBadRequest creates DeleteApisBadRequest with default headers values
func NewDeleteApisBadRequest() *DeleteApisBadRequest {

	return &DeleteApisBadRequest{}
}

// WithPayload adds the payload to the delete apis bad request response
func (o *DeleteApisBadRequest) WithPayload(payload *models.Error) *DeleteApisBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apis bad request response
func (o *DeleteApisBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApisBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteApisUnauthorizedCode is the HTTP code returned for type DeleteApisUnauthorized
const DeleteApisUnauthorizedCode int = 401

/*DeleteApisUnauthorized Unauthorized. Invalid authentication credentials.

swagger:response deleteApisUnauthorized
*/
type DeleteApisUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteApisUnauthorized creates DeleteApisUnauthorized with default headers values
func NewDeleteApisUnauthorized() *DeleteApisUnauthorized {

	return &DeleteApisUnauthorized{}
}

// WithPayload adds the payload to the delete apis unauthorized response
func (o *DeleteApisUnauthorized) WithPayload(payload *models.Error) *DeleteApisUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apis unauthorized response
func (o *DeleteApisUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApisUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteApisNotFoundCode is the HTTP code returned for type DeleteApisNotFound
const DeleteApisNotFoundCode int = 404

/*DeleteApisNotFound Not Found.
Requested API does not exist.


swagger:response deleteApisNotFound
*/
type DeleteApisNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteApisNotFound creates DeleteApisNotFound with default headers values
func NewDeleteApisNotFound() *DeleteApisNotFound {

	return &DeleteApisNotFound{}
}

// WithPayload adds the payload to the delete apis not found response
func (o *DeleteApisNotFound) WithPayload(payload *models.Error) *DeleteApisNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apis not found response
func (o *DeleteApisNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApisNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteApisInternalServerErrorCode is the HTTP code returned for type DeleteApisInternalServerError
const DeleteApisInternalServerErrorCode int = 500

/*DeleteApisInternalServerError Internal Server Error.

swagger:response deleteApisInternalServerError
*/
type DeleteApisInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteApisInternalServerError creates DeleteApisInternalServerError with default headers values
func NewDeleteApisInternalServerError() *DeleteApisInternalServerError {

	return &DeleteApisInternalServerError{}
}

// WithPayload adds the payload to the delete apis internal server error response
func (o *DeleteApisInternalServerError) WithPayload(payload *models.Error) *DeleteApisInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apis internal server error response
func (o *DeleteApisInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApisInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
