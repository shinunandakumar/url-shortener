package controllers

import (
	"fmt"
	"math/rand"

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
