package config

import (
	"os"
	"strconv"
)

func getEnv(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// RedisAddress will return the redis address from the env
func RedisAddress() string {
	return getEnv("REDISADDRESS", "localhost")
}

// RedisPort will return the redis port from the env
func RedisPort() string {
	return getEnv("REDISPORT", "6379")
}

// RedisPass will return the redis password from the env
func RedisPass() string {
	return getEnv("REDISPASS", "testpass")
}

// RedisPort will return the redis port from the env
func ShortenerLength() int {
	x, _ := strconv.Atoi(getEnv("SHORTENERLENGTH", "6"))
	return x
}

// MetricsLenth will return the number of data to be shown in the metrics from the env
func MetricsLenth() int {
	x, _ := strconv.Atoi(getEnv("METRICSLENTH", "3"))
	return x
}

// RedisPort will return the redis port from the env
func Host() string {
	return getEnv("HOST", "localhost:8080")
}

// MetricsKey will return the key name of metrics data to be stored in the redis from the env
func MetricsKey() string {
	return getEnv("METRICSKEY", "metrics")
}
