package config

import (
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	Port     string
	RedisURL string
	RootURL  *url.URL
}

func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8084"
	}
	redisUrl := os.Getenv("REDIS_URL")
	rootUrlRaw := os.Getenv("ROOT_URL")
	if rootUrlRaw == "" {
		return nil, fmt.Errorf("ROOT_URL env must be provided")
	}
	rootUrl, err := url.Parse(rootUrlRaw)
	if err != nil {
		return nil, fmt.Errorf("error parsing ROOT_URL: %v", err.Error())
	}
	return &Config{
		Port:     port,
		RedisURL: redisUrl,
		RootURL:  rootUrl,
	}, nil
}
