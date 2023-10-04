package redis

// redis package is used for the transactions in the redis
import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v7"
	"github.com/shinunandakumar/url-shortener/config"
)

// NewClient function will create a client and return
func NewClient() redis.UniversalClient {

	// TODO Change to env
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{fmt.Sprintf("%s:%s", config.RedisAddress(), config.RedisPort())},
		Password: config.RedisPass(),
	})

	// try to ping data base
	if _, err := rdb.Ping().Result(); err != nil {
		panic(err)
	}

	return rdb
}

// FilterByKey is an experiment function, need to do some reworks on this function
// just using it now for the completion purpose
func FilterByKey(rdb redis.UniversalClient, value string) string {
	// Retrieve all keys and values
	keys, err := rdb.Keys("*").Result()
	if err != nil {
		log.Fatal(err)
	}

	for _, key := range keys {
		val, err := rdb.Get(key).Result()

		if err != nil {
			log.Println("Error occured", err)
			continue
		}

		if val == value {
			return key
		}
	}
	return ""
}

// SortByDESC is an experiment function, need to do some reworks on this function
// just using it now for the completion purpose
func SortByDESC(rdb redis.UniversalClient) map[string]int {
	return nil
}
