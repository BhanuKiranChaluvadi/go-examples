package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

/*
client always expect json struct of this stuct
*/
type APIError struct {
	Context string `json:"context"`
	Message string `json:"message"`
}

func httpRespond(w http.ResponseWriter, httpStatusCode int, apiError *APIError) {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	// log.Formatter = new(logrus.TextFormatter) //default
	// log.Formatter.(*logrus.TextFormatter).DisableColors = false    // remove colors
	// log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false // remove timestamp from test output
	// log.Formatter.(*logrus.TextFormatter).FullTimestamp = true
	log.Debug("This is a debug log")
	log.Info("This is a info log")
	log.Error("This is a error log")
}

func main() {

	var w http.ResponseWriter
	statusCode := http.StatusInternalServerError
	apiError := APIError{
		Context: "context",
		Message: "message",
	}
	httpRespond(w, statusCode, &apiError)
}
