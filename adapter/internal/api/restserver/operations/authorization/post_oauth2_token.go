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

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostOauth2TokenHandlerFunc turns a function with the right signature into a post oauth2 token handler
type PostOauth2TokenHandlerFunc func(PostOauth2TokenParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOauth2TokenHandlerFunc) Handle(params PostOauth2TokenParams) middleware.Responder {
	return fn(params)
}

// PostOauth2TokenHandler interface for that can handle valid post oauth2 token params
type PostOauth2TokenHandler interface {
	Handle(PostOauth2TokenParams) middleware.Responder
}

// NewPostOauth2Token creates a new http.Handler for the post oauth2 token operation
func NewPostOauth2Token(ctx *middleware.Context, handler PostOauth2TokenHandler) *PostOauth2Token {
	return &PostOauth2Token{Context: ctx, Handler: handler}
}

/* PostOauth2Token swagger:route POST /oauth2/token Authorization postOauth2Token

Get an access token

This operation can be used to get an access token by providing the username and password
in the autherization header


*/
type PostOauth2Token struct {
	Context *middleware.Context
	Handler PostOauth2TokenHandler
}

func (o *PostOauth2Token) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostOauth2TokenParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostOauth2TokenOKBody post oauth2 token o k body
//
// swagger:model PostOauth2TokenOKBody
type PostOauth2TokenOKBody struct {

	// access token
	AccessToken string `json:"accessToken,omitempty"`
}

// Validate validates this post oauth2 token o k body
func (o *PostOauth2TokenOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post oauth2 token o k body based on context it is used
func (o *PostOauth2TokenOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOauth2TokenOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOauth2TokenOKBody) UnmarshalBinary(b []byte) error {
	var res PostOauth2TokenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
