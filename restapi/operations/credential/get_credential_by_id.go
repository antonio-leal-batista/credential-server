// Code generated by go-swagger; DO NOT EDIT.

package credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCredentialByIDHandlerFunc turns a function with the right signature into a get credential by Id handler
type GetCredentialByIDHandlerFunc func(GetCredentialByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCredentialByIDHandlerFunc) Handle(params GetCredentialByIDParams) middleware.Responder {
	return fn(params)
}

// GetCredentialByIDHandler interface for that can handle valid get credential by Id params
type GetCredentialByIDHandler interface {
	Handle(GetCredentialByIDParams) middleware.Responder
}

// NewGetCredentialByID creates a new http.Handler for the get credential by Id operation
func NewGetCredentialByID(ctx *middleware.Context, handler GetCredentialByIDHandler) *GetCredentialByID {
	return &GetCredentialByID{Context: ctx, Handler: handler}
}

/*GetCredentialByID swagger:route GET /credential/{credentialId} credential getCredentialById

Get a Credential by ID

This service returns a credential verifiable

*/
type GetCredentialByID struct {
	Context *middleware.Context
	Handler GetCredentialByIDHandler
}

func (o *GetCredentialByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCredentialByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
