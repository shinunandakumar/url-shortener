package controllers

// This is a url shortener project for an interview task
// an experimental project
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/shinunandakumar/url-shortener/config"
	"github.com/shinunandakumar/url-shortener/pkg/redis"
)

// ErrorResponse struct can be used whenever things go on the wrong way
type ErrorResponse struct {
	Status  int    `json:"status_code"`
	Message string `json:"message"`
	Error   error  `json:"error_code"`
}

// will be used for both binding and retuning the success state
type Params struct {
	URL    string `json:"url"`
	Status int    `json:"status_code,omitempty"`
}

// A health check controller to check the health of application
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!")
}

// generateURL is an http function used generate short urls with incoming urls and returns it.
func GenerateURL(w http.ResponseWriter, r *http.Request) {

	log.Println("Generating url")

	// parse the request
	params := Params{}
	//bind the request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil || params.URL == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp := ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid request", Error: err}
		json.NewEncoder(w).Encode(resp)
		return
	}

	// create redis client
	rdb := redis.NewClient()
	defer rdb.Close()

	// add referenece type if it is missing
	if !strings.HasPrefix(strings.ToLower(params.URL), "http") {
		params.URL = "http://" + params.URL
	}

	// check if the url already registered or not
	// TODO errcheck
	val, _ := rdb.Get(params.URL).Result()
	if val != "" {
		resp := Params{URL: ShortURL(val), Status: http.StatusOK}
		json.NewEncoder(w).Encode(resp)
		return
	}

	// generate a random string
	randStr := generateRandString(config.ShortenerLength())
	shorturl := ShortURL(randStr)
	err := rdb.Set(params.URL, shorturl, 0).Err()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := ErrorResponse{Status: http.StatusBadRequest, Message: "Redis failure", Error: err}
		json.NewEncoder(w).Encode(resp)
		return
	}
	// TODO insert to metrics

	log.Println("url generation was successful", shorturl)
	resp := Params{URL: shorturl, Status: http.StatusOK}
	json.NewEncoder(w).Encode(resp)
	return
}

// Redirect is an http function,will redirected to the original url, When the user is clicked on the short url.
func Redirect(w http.ResponseWriter, r *http.Request) {

	inUrl := r.URL.Path[1:]
	log.Println("Redirecting", inUrl)

	// create redis client
	rdb := redis.NewClient()
	defer rdb.Close()

	val, err := rdb.Get(inUrl).Result()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := ErrorResponse{Status: http.StatusBadRequest, Message: "Redis failure", Error: err}
		json.NewEncoder(w).Encode(resp)
		return
	}
	fmt.Fprintf(w, "Redirecting the url")
}

// metrics is an http function,it will return the most used url providers
// TODO make a proper description
func Metrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Generating the url")
}
