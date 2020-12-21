// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/MicahParks/shakesearch/configure"
	"github.com/MicahParks/shakesearch/endpoints"
	"github.com/MicahParks/shakesearch/restapi/operations"
)

//go:generate swagger generate server --target ../../shakesearch --name Shakesearch --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.ShakesearchAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ShakesearchAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Configure the service.
	logger, shakeSearch, tmpl, err := configure.Configure()
	if err != nil {
		log.Fatalf("Failed to configure the service.\nError: %s", err.Error())
	}

	// Set the generated code logger to a named zap logger.
	api.Logger = logger.Named("Generated code").Infof

	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()

	api.HTMLProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {

		// Assert the data is an io.ReadCloser.
		readCloser, ok := data.(io.ReadCloser)
		if !ok {
			return nil
		}

		// Read in all the data.
		var readIn []byte
		var err error
		if readIn, err = ioutil.ReadAll(readCloser); err != nil {
			return err
		}

		// Write all the data.
		if _, err = w.Write(readIn); err != nil {
			return err
		}

		// Close the io.ReadCloser.
		return readCloser.Close()
	})
	api.JSONProducer = runtime.JSONProducer()

	api.PublicShakeSearchHandler = endpoints.HandleSearch(logger.Named("/api/search"), shakeSearch)
	api.PublicShakeWorksHandler = endpoints.HandleWorks(logger.Named("/api/works"), shakeSearch, tmpl)
	api.SystemAliveHandler = endpoints.HandleAlive()

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {

	// Create an incoming request rate limiter that only allows 1 request per section and forgets about clients after 1
	// hour.
	limit := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	// Find the IP of the client in the X-Forwarded-For header, because Caddy will be the server in front of this.
	limit.SetIPLookups([]string{"X-Forwarded-For"})

	// Follow the HTTP middleware pattern.
	return tollbooth.LimitHandler(limit, handler)
}
