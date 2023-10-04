package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"

	"github.com/shinunandakumar/url-shortener/config"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// ShortURL is used to concat shorturl with host for now
// will change later if anything is required
func ShortURL(s string) string {
	return fmt.Sprintf("%s/%s", config.Host(), s)
}

// GetSite will return the domain from the url string
func GetDomain(s string) string {

	u, err := url.Parse(s)
	if err != nil {
		log.Println("error occured while url parsing")
		panic(err)
	}
	return u.Host
}

// ReConstructMetrics function is used to generate map data for insert/update ops
// error check is not planning to do
func ReConstructMetrics(metricsArr []byte, domain string) []byte {
	domainArr := make(map[string]int)
	_ = json.Unmarshal(metricsArr, &domainArr)

	if domainArr[domain] != 0 {
		domainArr[domain] += 1
	} else {
		domainArr[domain] = 1
	}
	metricsMarshelled, _ := json.Marshal(domainArr)

	return metricsMarshelled
}
