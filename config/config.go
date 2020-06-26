package config

import "os"

type Config struct {
	RedisAddress  string
	RedisPassword string
	HttpPort      string
}

func New() *Config {
	return &Config{
		RedisAddress:  ParseEnv("REDIS_ADDRESS", "redis:6379"),
		RedisPassword: ParseEnv("REDIS_PASSWORD", ""),
		HttpPort:      ParseEnv("SHORT_HTTP_PORT", "8080"),
	}
}

func ParseEnv(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}
