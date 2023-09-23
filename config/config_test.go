package config

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/nirdosh17/go-sdk-template/client"
	"github.com/nirdosh17/go-sdk-template/test"
)

func TestConfig_NewConfig(t *testing.T) {
	c := NewConfig()
	test.ExpectEqual(t, "HTTPTimeout", client.DefaultHTTPTimeout, c.HTTPTimeout)
	test.ExpectNotNil(t, "HTTPClient", c.HTTPClient)
	test.ExpectNotNil(t, "Retryer", c.Retryer)
	test.ExpectEqual(t, "Endpoint", apiBasePath, c.Endpoint)
}

func TestConfig_WithHttpClient(t *testing.T) {
	config := NewConfig()
	to := 60 * time.Second
	custom := http.Client{Timeout: to}
	config.WithHttpClient(&custom)
	test.ExpectEqual(t, "HTTPClient timeout", to, config.HTTPClient.Timeout)
}

func TestConfig_WithHttpTimeout(t *testing.T) {
	config := NewConfig()
	to := 60 * time.Second
	config.WithHttpTimeout(to)
	test.ExpectEqual(t, "HTTPClient timeout", to, config.HTTPClient.Timeout)
}

func TestConfig_WithEndpoint(t *testing.T) {
	config := NewConfig()
	e := "https://region2.aiservice.com"
	config.WithEndpoint(e)
	test.ExpectEqual(t, "Endpoint", config.Endpoint, e)
}

type MockRetry struct {
	MaxRetries int
}

func (r *MockRetry) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	err := fn(ctx)
	return err
}
func (r *MockRetry) SetMaxRetries(n int) {
	r.MaxRetries = n
}

func TestConfig_WithRetryer(t *testing.T) {
	config := NewConfig()
	n := &MockRetry{}
	config.WithRetryer(n)
	test.ExpectSameType(t, "Retryer", n, config.Retryer)
}

func TestConfig_WithMaxRetries(t *testing.T) {
	config := NewConfig()
	mr := &MockRetry{MaxRetries: 3}
	config.WithRetryer(mr)
	config.WithMaxRetries(2)
	test.ExpectEqual(t, "Retryer.MaxRetries", 2, mr.MaxRetries)
	test.ExpectEqual(t, "Config.MaxRetries", 2, config.MaxRetries)
}
