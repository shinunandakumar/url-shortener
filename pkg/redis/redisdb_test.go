package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	client := NewClient()
	defer client.Close()

	pong, err := client.Ping().Result()

	assert.Nil(t, err)
	assert.Equal(t, "PONG", pong)
}

func TestFilterByKey(t *testing.T) {

	client := NewClient()
	defer client.Close()

	key := "test_key"
	value := "test_value"

	// Set a test key-value pair in Redis
	err := client.Set(key, value, 0).Err()
	assert.Nil(t, err)

	foundKey := FilterByKey(client, value)

	assert.Equal(t, key, foundKey)

	client.Del(key)
}

func TestFilterByKey_NotFound(t *testing.T) {

	client := NewClient()
	defer client.Close()

	// Non-existent value
	value := "non_existent_value"

	foundKey := FilterByKey(client, value)

	assert.Equal(t, "", foundKey)
}
