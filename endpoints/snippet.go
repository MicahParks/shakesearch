package endpoints

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/MicahParks/shakesearch"
	"github.com/MicahParks/shakesearch/models"
	"github.com/MicahParks/shakesearch/restapi/operations/public"
)

var (

	// defaultMaxLi8nes is the largest number of surrounding lines to return if none was given.
	defaultMaxLines = int64(1000)
)

// HandleWorks creates a /api/works endpoint handler via a closure. It will create an HTML document of a snippet of the
// request line number of Shakespeare's works.
func HandleWorks(logger *zap.SugaredLogger, shakeSearcher *shakesearch.ShakeSearcher, tmpl *template.Template) public.ShakeWorksHandlerFunc {
	return func(params public.ShakeWorksParams) middleware.Responder {

		// Debug info.
		logger.Debugw("Gathering work snippet.",
			"lineNumber", params.Line,
			"maxLines", params.MaxLines,
		)

		// Use the default number of max lines to use if none was given.
		if params.MaxLines == nil || *params.MaxLines < 0 {
			params.MaxLines = &defaultMaxLines
		}

		// Gather the surrounding lines from the given line.
		lines := shakeSearcher.SurroundingWorks(int(params.Line), uint(*params.MaxLines))

		// Create a buffer to write the populated HTML template to.
		buf := bytes.NewBuffer(nil)

		// Write the data to the HTML template.
		if err := tmpl.Execute(buf, lines); err != nil {

			// Log at the appropriate level.
			message := "Failed to execute the HTML template."
			logger.Errorw(message,
				"lineNumber", params.Line,
				"error", err.Error(),
			)

			// Report the error to the client.
			code := int64(500)
			return &public.ShakeWorksDefault{Payload: &models.Error{
				Code:    &code,
				Message: &message,
			}}
		}

		// Return the HTML document to the client.
		payload := ioutil.NopCloser(buf)
		return &public.ShakeWorksOK{
			Payload: payload,
		}
	}
}
