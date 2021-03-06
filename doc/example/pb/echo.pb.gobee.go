// Code generated by protoc-gen-gobee
// source: pb/echo.proto
// DO NOT EDIT!

/*
Package example is a self-registering gRPC and JSON+Swagger service definition.

It conforms to the github.com/daewood/gobee Service interface.
*/
package example

import (
	"net/http"
	"strings"

	"github.com/daewood/gobee/transport"
	"github.com/daewood/gobee/transport/httpruntime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// EchoServerDesc is a descriptor/registrator for the EchoServerServer.
type EchoServerDesc struct {
	svc EchoServerServer
}

// NewEchoServerServiceDesc creates new registrator for the EchoServerServer.
func NewEchoServerServiceDesc(svc EchoServerServer) *EchoServerDesc {
	return &EchoServerDesc{svc: svc}
}

// RegisterGRPC implements service registrator interface.
func (d *EchoServerDesc) RegisterGRPC(s *grpc.Server) {
	RegisterEchoServerServer(s, d.svc)
}

// SwaggerDef returns this file's Swagger definition.
func (d *EchoServerDesc) SwaggerDef() []byte {
	return _swaggerDef_pb_echo_proto
}

// RegisterHTTP registers this service's HTTP handlers/bindings.
func (d *EchoServerDesc) RegisterHTTP(mux transport.Router) {

	// Handlers for Echo

	mux.HandleFunc("/"+pattern_gobee_EchoServer_Echo_0, func(w http.ResponseWriter, r *http.Request) {
		//TODO only POST is supported atm

		inbound, outbound := httpruntime.MarshalerForRequest(r)
		var req StringMessage
		err := inbound.Unmarshal(r.Body, &req)
		if err != nil {
			httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't read request JSON"), nil)
			return
		}
		ret, err := d.svc.Echo(r.Context(), &req)
		if err != nil {
			httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "returned from handler"), nil)
			return
		}

		w.Header().Set("Content-Type", outbound.ContentType())
		err = outbound.Marshal(w, ret)
		if err != nil {
			httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"), nil)
			return
		}
	})

}

var _swaggerDef_pb_echo_proto = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "pb/echo.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/example/echo": {
      "post": {
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/exampleStringMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/exampleStringMessage"
            }
          }
        ],
        "tags": [
          "EchoServer"
        ]
      }
    }
  },
  "definitions": {
    "exampleStringMessage": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    }
  }
}

`)

var (
	pattern_gobee_EchoServer_Echo_0 = strings.Join([]string{"v1", "example", "echo"}, "/")
)
