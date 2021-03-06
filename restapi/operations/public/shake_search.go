// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ShakeSearchHandlerFunc turns a function with the right signature into a shake search handler
type ShakeSearchHandlerFunc func(ShakeSearchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ShakeSearchHandlerFunc) Handle(params ShakeSearchParams) middleware.Responder {
	return fn(params)
}

// ShakeSearchHandler interface for that can handle valid shake search params
type ShakeSearchHandler interface {
	Handle(ShakeSearchParams) middleware.Responder
}

// NewShakeSearch creates a new http.Handler for the shake search operation
func NewShakeSearch(ctx *middleware.Context, handler ShakeSearchHandler) *ShakeSearch {
	return &ShakeSearch{Context: ctx, Handler: handler}
}

/*ShakeSearch swagger:route GET /api/search public shakeSearch

Fuzzy search for a string of text in Shakespeare's works.

The string will be used in conjunction with [this project](https://github.com/sahilm/fuzzy) to perform a fuzzy search on Shakespeare's works.

*/
type ShakeSearch struct {
	Context *middleware.Context
	Handler ShakeSearchHandler
}

func (o *ShakeSearch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewShakeSearchParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
