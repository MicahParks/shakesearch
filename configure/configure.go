package configure

import (
	"html/template"
	"io/ioutil"
	"os"

	"go.uber.org/zap"

	"github.com/MicahParks/shakesearch"
)

const (

	// defaultWorksPath is the default path to find the file that contains the complete works of Shakespeare in text.
	defaultWorksPath = "completeworks.txt"

	// defaultTemplatePath is the default path to find the HTML template for a snippet of Shakespeare's works.
	defaultTemplatePath = "snippet.gohtml"

	// templateEnvVar represents the environment variable used to find the path to the file that contains the HTML
	// template for a snippet of Shakespeare's works.
	templateEnvVar = "SNIPPET_TEMPLATE"

	// worksPathEnvVar represent the environment variable used to find the path to the file that contains complete works
	// of Shakespeare in text. It will use the default value if none is given.
	worksPathEnvVar = "SHAKESPEARES_WORKS"
)

// Configure gathers all required information and creates the required Go structs to run the service.
func Configure() (logger *zap.SugaredLogger, shakeSearcher *shakesearch.ShakeSearcher, tmpl *template.Template, err error) {

	// Create a logger.
	var zapLogger *zap.Logger
	if zapLogger, err = zap.NewDevelopment(); err != nil { // TODO Make NewProduction.
		return nil, nil, nil, err
	}
	logger = zapLogger.Sugar()
	logger.Info("Logger created. Starting configuration.")

	// Get the complete works of Shakespeare's file path from an environment variable.
	worksPath := os.Getenv(worksPathEnvVar)
	if worksPath == "" {
		worksPath = defaultWorksPath
	}

	// Create the ShakeSearcher.
	if shakeSearcher, err = shakesearch.NewShakeSearcher(worksPath); err != nil {
		logger.Fatalw("Failed to create the ShakeSearcher.",
			"error", err.Error(),
		)
		return nil, nil, nil, err // Should be unreachable.
	}

	// Get the HTML template for a snippet of Shakespeare's works' file path from an environment variable.
	templatePath := os.Getenv(templateEnvVar)
	if templatePath == "" {
		templatePath = defaultTemplatePath
	}

	// Read in the HTML template file.
	var fileData []byte
	if fileData, err = ioutil.ReadFile(templatePath); err != nil {
		return nil, nil, nil, err
	}

	// Create the HTML template.
	if tmpl, err = template.New("").Parse(string(fileData)); err != nil {
		return nil, nil, nil, err
	}

	return logger, shakeSearcher, tmpl, nil
}
