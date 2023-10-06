package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	httpreq := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)

	handler.ServeHTTP(httpreq, req)

	assert.Equal(t, http.StatusOK, httpreq.Code)
	assert.Equal(t, "OK!", httpreq.Body.String())
}

func TestGenerateURL(t *testing.T) {
	// Create a sample request body
	params := Params{URL: "https://youtube.com"}

	// Marshal the request body to JSON
	reqBody, _ := json.Marshal(params)

	req, err := http.NewRequest("POST", "/api/v1/generate", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	httpreq := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateURL)
	handler.ServeHTTP(httpreq, req)

	assert.Equal(t, http.StatusOK, httpreq.Code)

}

func TestRedirect(t *testing.T) {
	// TODO complete the redirection
	log.Println("testing redirecting")
}

// TODO TestMetrics needs some more workarounds
// func TestMetrics(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/api/v1/metrics", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	httpreq := httptest.NewRecorder()
// 	handler := http.HandlerFunc(Metrics)

// 	handler.ServeHTTP(httpreq, req)
// 	fmt.Println("http.StatusAccepted", http.StatusAccepted)
// 	fmt.Println("http.StatusAccepted", httpreq.Code)

// 	assert.Equal(t, http.StatusAccepted, httpreq.Code)

// }
