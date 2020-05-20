package http

import (
	"net/http"
)

type Config struct {
}

func (c *Config) Client() (*http.Client, error) {
	client := &http.Client{}
	return client, nil
}
