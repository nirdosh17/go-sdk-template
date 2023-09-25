// Package config provides main configuration object to initialize the apis. The config object is used while creating new services.
//
// long description with examples
package config

import (
	"github.com/nirdosh17/go-sdk-template/client"
	"github.com/nirdosh17/go-sdk-template/logger"
)

const (
	apiBasePath = "http://localhost:8000"
)

type Config struct {
	// Endpoint is optional URL that overrides default service endpoint.
	// Some services offer regional endpoints which you can choose based on proximity for minimal latency.
	Endpoint string
	// HTTP client to use while sending requests. Defaults to `http.DefaultClient`
	HTTPClient client.HTTPClient
	// Retryer function
	Retryer client.Retryer
	// The maximum number of times a request will be retried before it is considered failed. Defaults to 3.
	MaxRetries int
	// logger function for sdk
	Logger logger.Logger
	// Debug enables verbose logging if set to true
	Debug bool
}

// NewConfig return a instance of config with default settings.
func NewConfig() *Config {
	return &Config{
		HTTPClient: client.DefaultClient(),
		Retryer:    client.DefaultRetryer(),
		Endpoint:   apiBasePath,
		Logger:     logger.NewDefaultLogger(),
		Debug:      false,
	}
}

// WithHTTPClient overrides default http client `http.DefaultClient`.
func (c *Config) WithHTTPClient(hc client.Client) *Config {
	c.HTTPClient = hc
	return c
}

// WithEndpoint overrides default endpoint.
func (c *Config) WithEndpoint(endpoint string) *Config {
	c.Endpoint = endpoint
	return c
}

// WithRetryer allows to override default retry function.
func (c *Config) WithRetryer(r client.Retryer) *Config {
	c.Retryer = r
	return c
}

// WithMaxRetries overrides default max retry count of default retryer.
func (c *Config) WithMaxRetries(n int) *Config {
	if n > 0 {
		// TODO: save max retries in single place
		c.MaxRetries = n
		c.Retryer.SetMaxRetries(n)
	}
	return c
}

// WithLogger overrides default logger.
func (c *Config) WithLogger(logger logger.Logger) *Config {
	c.Logger = logger
	return c
}

// WithDebugEnabled enables debug flag which for verbose logging.
func (c *Config) WithDebugEnabled() *Config {
	c.Debug = true
	return c
}
