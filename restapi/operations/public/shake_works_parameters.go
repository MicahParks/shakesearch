// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewShakeWorksParams creates a new ShakeWorksParams object
// with the default values initialized.
func NewShakeWorksParams() ShakeWorksParams {

	var (
		// initialize parameters with default values

		maxLinesDefault = int64(1000)
	)

	return ShakeWorksParams{
		MaxLines: &maxLinesDefault,
	}
}

// ShakeWorksParams contains all the bound params for the shake works operation
// typically these are obtained from a http.Request
//
// swagger:parameters shakeWorks
type ShakeWorksParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The line number to fetch & focus.
	  Required: true
	  In: query
	*/
	Line int64
	/*The maximum number of lines to return. Half before, half after.
	  In: query
	  Default: 1000
	*/
	MaxLines *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewShakeWorksParams() beforehand.
func (o *ShakeWorksParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLine, qhkLine, _ := qs.GetOK("line")
	if err := o.bindLine(qLine, qhkLine, route.Formats); err != nil {
		res = append(res, err)
	}

	qMaxLines, qhkMaxLines, _ := qs.GetOK("maxLines")
	if err := o.bindMaxLines(qMaxLines, qhkMaxLines, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLine binds and validates parameter Line from query.
func (o *ShakeWorksParams) bindLine(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("line", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("line", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("line", "query", "int64", raw)
	}
	o.Line = value

	return nil
}

// bindMaxLines binds and validates parameter MaxLines from query.
func (o *ShakeWorksParams) bindMaxLines(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewShakeWorksParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("maxLines", "query", "int64", raw)
	}
	o.MaxLines = &value

	return nil
}