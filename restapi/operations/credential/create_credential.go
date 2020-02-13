// Code generated by go-swagger; DO NOT EDIT.

package credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateCredentialHandlerFunc turns a function with the right signature into a create credential handler
type CreateCredentialHandlerFunc func(CreateCredentialParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCredentialHandlerFunc) Handle(params CreateCredentialParams) middleware.Responder {
	return fn(params)
}

// CreateCredentialHandler interface for that can handle valid create credential params
type CreateCredentialHandler interface {
	Handle(CreateCredentialParams) middleware.Responder
}

// NewCreateCredential creates a new http.Handler for the create credential operation
func NewCreateCredential(ctx *middleware.Context, handler CreateCredentialHandler) *CreateCredential {
	return &CreateCredential{Context: ctx, Handler: handler}
}

/*CreateCredential swagger:route POST /credential credential createCredential

Create Credential

This service is responsible for generating verifiable credentials, storing the credential hash in the blockchain. You send the credential data with type of credential, range of valid dates and evidence generated about credential as optional

*/
type CreateCredential struct {
	Context *middleware.Context
	Handler CreateCredentialHandler
}

func (o *CreateCredential) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateCredentialParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
