package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisAddress(t *testing.T) {
	os.Setenv("REDISADDRESS", "redisnew")
	address := RedisAddress()
	assert.Equal(t, "redisnew", address)

	os.Unsetenv("REDISADDRESS")
}

func TestRedisPort(t *testing.T) {
	os.Setenv("REDISPORT", "1234")
	port := RedisPort()
	assert.Equal(t, "1234", port)

	os.Unsetenv("REDISPORT")
}

func TestRedisPass(t *testing.T) {
	os.Setenv("REDISPASS", "test123")
	pass := RedisPass()
	assert.Equal(t, "test123", pass)

	os.Unsetenv("REDISPASS")
}

func TestShortenerLength(t *testing.T) {
	os.Setenv("SHORTENERLENGTH", "10")
	length := ShortenerLength()
	assert.Equal(t, 10, length)

	os.Unsetenv("SHORTENERLENGTH")
}

func TestMetricsLenth(t *testing.T) {
	os.Setenv("METRICSLENTH", "5")
	length := MetricsLenth()
	assert.Equal(t, 5, length)

	os.Unsetenv("METRICSLENTH")
}

func TestHost(t *testing.T) {
	os.Setenv("HOST", "testhost:8081")
	host := Host()
	assert.Equal(t, "testhost:8081", host)

	os.Unsetenv("HOST")
}

func TestMetricsKey(t *testing.T) {
	os.Setenv("METRICSKEY", "test_metrics_key")
	key := MetricsKey()
	assert.Equal(t, "test_metrics_key", key)

	os.Unsetenv("METRICSKEY")
}
