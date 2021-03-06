// Code generated by go-swagger; DO NOT EDIT.

package did

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ValidateDidHandlerFunc turns a function with the right signature into a validate did handler
type ValidateDidHandlerFunc func(ValidateDidParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ValidateDidHandlerFunc) Handle(params ValidateDidParams) middleware.Responder {
	return fn(params)
}

// ValidateDidHandler interface for that can handle valid validate did params
type ValidateDidHandler interface {
	Handle(ValidateDidParams) middleware.Responder
}

// NewValidateDid creates a new http.Handler for the validate did operation
func NewValidateDid(ctx *middleware.Context, handler ValidateDidHandler) *ValidateDid {
	return &ValidateDid{Context: ctx, Handler: handler}
}

/*ValidateDid swagger:route GET /did/{did} did validateDid

Validate a DID

This service validates through JWT if a DID is valid. The JWT contains a did of a person, institution or object that needs be validated. The owner of did sign the payload with its private key

*/
type ValidateDid struct {
	Context *middleware.Context
	Handler ValidateDidHandler
}

func (o *ValidateDid) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewValidateDidParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
