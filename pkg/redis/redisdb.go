package redis

// redis package is used for the transactions in the redis
import (
	"fmt"

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

	redis.Strings(rdb.Do("MGET", "k1", "k2", "k3"))

	return rdb
}
