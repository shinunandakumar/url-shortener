package redis

// redis package is used for the transactions in the redis
import (
	"log"

	"github.com/go-redis/redis/v7"
)

func NewClient() error {

	// Change to env
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{"redis:6379"},
		Password: "testpass",
	})

	log.Println(rdb.Ping())
	// try to ping data base
	if _, err := rdb.Ping().Result(); err != nil {
		return err
	}

	err := rdb.Set("python", "interpreter", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("python").Result()
	if err != nil {
		panic(err)
	}
	log.Println("python", val)

	val2, err := rdb.Get("key2").Result()
	if err == redis.Nil {
		log.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		log.Println("key2", val2)
	}

	return nil
}
