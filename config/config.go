// Package api contains main configuration object for sdk consumers to use while making request to all services.

package config

import (
	"net/http"
	"time"

	"github.com/nirdosh17/go-sdk-template/client"
)

const (
	apiBasePath = "http://localhost:8000"
)

type Config struct {
	// Endpoint is optional URL that overrides default endpoint for the service. e.g. you can use endpoint which is situated near to you to improve latency.
	Endpoint string
	// HTTP client to use while sending requests. Defaults to `http.DefaultClient`
	HTTPClient *http.Client
	// Default HTTP timeout can be overridden using this. Defaults to 30 seconds.
	HTTPTimeout time.Duration
	// Retryer function
	Retryer client.Retryer
	// The maximum number of times a request will be retried before it is considered failed. Defaults to 3.
	MaxRetries int
}

// NewConfig return a instance of config with default settings.
func NewConfig() *Config {
	return &Config{
		HTTPClient:  client.DefaultClient(),
		Retryer:     client.DefaultRetryer(),
		HTTPTimeout: client.DefaultHTTPTimeout,
		Endpoint:    apiBasePath,
	}
}

// WithHttpClient overrides default http client `http.DefaultClient`.
func (c *Config) WithHttpClient(hc *http.Client) *Config {
	c.HTTPClient = hc
	return c
}

// WithHttpTimeout overrides default timeout duration `client.DefaultHTTPTimeout`.
func (c *Config) WithHttpTimeout(t time.Duration) *Config {
	c.HTTPClient.Timeout = t
	// validate this
	return c
}

// WithServiceEndpoint overrides default xcloud endpoint.
func (c *Config) WithEndpoint(endpoint string) *Config {
	c.Endpoint = endpoint
	return c
}

// WithRetryer allows to override default retry function
func (c *Config) WithRetryer(r client.Retryer) *Config {
	c.Retryer = r
	return c
}

// WithMaxRetries overrides default max retry count of default retryer
func (c *Config) WithMaxRetries(n int) *Config {
	if n > 0 {
		// TODO: save max retries in single place
		c.MaxRetries = n
		c.Retryer.SetMaxRetries(n)
	}
	return c
}
