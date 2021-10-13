// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreatestudentHandlerFunc turns a function with the right signature into a createstudent handler
type CreatestudentHandlerFunc func(CreatestudentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreatestudentHandlerFunc) Handle(params CreatestudentParams) middleware.Responder {
	return fn(params)
}

// CreatestudentHandler interface for that can handle valid createstudent params
type CreatestudentHandler interface {
	Handle(CreatestudentParams) middleware.Responder
}

// NewCreatestudent creates a new http.Handler for the createstudent operation
func NewCreatestudent(ctx *middleware.Context, handler CreatestudentHandler) *Createstudent {
	return &Createstudent{Context: ctx, Handler: handler}
}

/* Createstudent swagger:route POST /Student/newstudent/add Student createstudent

Add student recoard

*/
type Createstudent struct {
	Context *middleware.Context
	Handler CreatestudentHandler
}

func (o *Createstudent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreatestudentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}