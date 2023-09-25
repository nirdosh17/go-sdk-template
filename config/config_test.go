package config

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/nirdosh17/go-sdk-template/test"
)

func TestConfig_NewConfig(t *testing.T) {
	c := NewConfig()
	test.ExpectNotNil(t, "HTTPClient", c.HTTPClient)
	test.ExpectNotNil(t, "Retryer", c.Retryer)
	test.ExpectEqual(t, "Endpoint", apiBasePath, c.Endpoint)
}

func TestConfig_WithHTTPClient(t *testing.T) {
	config := NewConfig()
	to := 60 * time.Second
	custom := http.Client{Timeout: to}
	config.WithHTTPClient(&custom)
	test.ExpectEqual(t, "HTTPClient", &custom, config.HTTPClient)
}

func TestConfig_WithEndpoint(t *testing.T) {
	config := NewConfig()
	e := "https://region2.aiservice.com"
	config.WithEndpoint(e)
	test.ExpectEqual(t, "Endpoint", config.Endpoint, e)
}

type mockRetry struct {
	MaxRetries int
}

func (r *mockRetry) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	err := fn(ctx)
	return err
}
func (r *mockRetry) SetMaxRetries(n int) {
	r.MaxRetries = n
}

func TestConfig_WithRetryer(t *testing.T) {
	config := NewConfig()
	n := &mockRetry{}
	config.WithRetryer(n)
	test.ExpectSameType(t, "Retryer", n, config.Retryer)
}

func TestConfig_WithMaxRetries(t *testing.T) {
	config := NewConfig()
	mr := &mockRetry{MaxRetries: 3}
	config.WithRetryer(mr)
	config.WithMaxRetries(2)
	test.ExpectEqual(t, "Retryer.MaxRetries", 2, mr.MaxRetries)
	test.ExpectEqual(t, "Config.MaxRetries", 2, config.MaxRetries)
}

type mockLogger struct {
}

func (l *mockLogger) Log(args ...interface{}) {
}

func TestConfig_WithLogger(t *testing.T) {
	config := NewConfig()
	l := &mockLogger{}
	config.WithLogger(l)
	test.ExpectSameType(t, "Logger", l, config.Logger)
}

func TestConfig_WithDebugEnabled(t *testing.T) {
	config := NewConfig()
	test.ExpectEqual(t, "config.Debug", false, config.Debug)

	config.WithDebugEnabled()
	test.ExpectEqual(t, "config.Debug", true, config.Debug)
}
