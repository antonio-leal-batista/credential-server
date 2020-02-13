// Code generated by go-swagger; DO NOT EDIT.

package credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RenewalCredentialHandlerFunc turns a function with the right signature into a renewal credential handler
type RenewalCredentialHandlerFunc func(RenewalCredentialParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RenewalCredentialHandlerFunc) Handle(params RenewalCredentialParams) middleware.Responder {
	return fn(params)
}

// RenewalCredentialHandler interface for that can handle valid renewal credential params
type RenewalCredentialHandler interface {
	Handle(RenewalCredentialParams) middleware.Responder
}

// NewRenewalCredential creates a new http.Handler for the renewal credential operation
func NewRenewalCredential(ctx *middleware.Context, handler RenewalCredentialHandler) *RenewalCredential {
	return &RenewalCredential{Context: ctx, Handler: handler}
}

/*RenewalCredential swagger:route PUT /credential/{credentialId} credential renewalCredential

Renewal a credential

This service updates a verifiable credential. First revoke the old credential with credential id, after generates a new credential with updated data from credential subject

*/
type RenewalCredential struct {
	Context *middleware.Context
	Handler RenewalCredentialHandler
}

func (o *RenewalCredential) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRenewalCredentialParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
