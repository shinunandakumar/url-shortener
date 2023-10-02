package controllers

// This is a url shortener project for an interview task
// an experimental project
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ErrorResponse struct can be used whenever things go on the wrong way
type ErrorResponse struct {
	Status  int    `json:"status_code"`
	Message string `json:"message"`
	Error   error  `json:"error_code"`
}

// A health check controller to check the health of application
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!")
}

// generateURL is an http function used generate short urls with incoming urls and returns it.
func GenerateURL(w http.ResponseWriter, r *http.Request) {

	// parse the request
	var params struct {
		URL string `json:"url"`
	}

	//bind the request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil || params.URL == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp := ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid request", Error: err}
		json.NewEncoder(w).Encode(resp)
		return
	}
	log.Println("params.URL----", params.URL)
	fmt.Fprintf(w, "Generating the url", params.URL)
}

// Redirect is an http function,will redirected to the original url, When the user is clicked on the short url.
func Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Redirecting the url")
}

// metrics is an http function,it will return the most used url providers
// TODO make a proper description
func Metrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Generating the url")
}
