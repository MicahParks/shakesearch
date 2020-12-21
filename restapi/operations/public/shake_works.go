// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ShakeWorksHandlerFunc turns a function with the right signature into a shake works handler
type ShakeWorksHandlerFunc func(ShakeWorksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ShakeWorksHandlerFunc) Handle(params ShakeWorksParams) middleware.Responder {
	return fn(params)
}

// ShakeWorksHandler interface for that can handle valid shake works params
type ShakeWorksHandler interface {
	Handle(ShakeWorksParams) middleware.Responder
}

// NewShakeWorks creates a new http.Handler for the shake works operation
func NewShakeWorks(ctx *middleware.Context, handler ShakeWorksHandler) *ShakeWorks {
	return &ShakeWorks{Context: ctx, Handler: handler}
}

/*ShakeWorks swagger:route GET /api/works public shakeWorks

TODO.

TODO

*/
type ShakeWorks struct {
	Context *middleware.Context
	Handler ShakeWorksHandler
}

func (o *ShakeWorks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewShakeWorksParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}