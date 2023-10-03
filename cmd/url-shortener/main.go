package main

// This is a url shortener project for an interview task
// an experimental project
import (
	"log"
	"net/http"

	"github.com/shinunandakumar/url-shortener/controllers"
	"github.com/shinunandakumar/url-shortener/pkg/redis"
)

func main() {

	// check
	redis.NewClient()

	// Registering the routes
	// TODO remove these from the main
	http.HandleFunc("/", controllers.Redirect)
	http.HandleFunc("/api/v1/health-check", controllers.HealthCheck)
	http.HandleFunc("/api/v1/generate-url", controllers.GenerateURL)
	http.HandleFunc("/api/v1/metrics", controllers.Metrics)

	log.Println("Server started at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
