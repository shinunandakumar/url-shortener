package config

// RedisAddress will return the redis address from the env
func RedisAddress() string {
	// TODO change to env
	return "redis"
}

// RedisPort will return the redis port from the env
func RedisPort() string {
	// TODO move to env
	return "6379"
}

// RedisPass will return the redis password from the env
func RedisPass() string {
	// TODO move to env
	return "testpass"
}

// RedisPort will return the redis port from the env
func ShortenerLength() int {
	// TODO move to env
	return 6
}

// MetricsLenth will return the number of data to be shown in the metrics from the env
func MetricsLenth() int {
	// TODO move to env
	return 3
}

// RedisPort will return the redis port from the env
func Host() string {
	// TODO move to env
	return "localhost:8080"
}

// MetricsKey will return the key name of metrics data to be stored in the redis from the env
func MetricsKey() string {
	// TODO move to env
	return "metrics"
}
