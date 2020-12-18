// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AliveHandlerFunc turns a function with the right signature into a alive handler
type AliveHandlerFunc func(AliveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AliveHandlerFunc) Handle(params AliveParams) middleware.Responder {
	return fn(params)
}

// AliveHandler interface for that can handle valid alive params
type AliveHandler interface {
	Handle(AliveParams) middleware.Responder
}

// NewAlive creates a new http.Handler for the alive operation
func NewAlive(ctx *middleware.Context, handler AliveHandler) *Alive {
	return &Alive{Context: ctx, Handler: handler}
}

/*Alive swagger:route GET /api/alive system alive

Used by Caddy or other reverse proxy to determine if the service is alive.

Any non-200 response means the service is not alive.

*/
type Alive struct {
	Context *middleware.Context
	Handler AliveHandler
}

func (o *Alive) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAliveParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}