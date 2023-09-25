package config_test

import (
	"context"

	"github.com/nirdosh17/go-sdk-template/api/chatai"
	"github.com/nirdosh17/go-sdk-template/config"
)

// CustomRetry must satisfy client.Retryer interface.
type CustomRetry struct {
	MaxRetries int
}

func (r *CustomRetry) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	// function received in the parameter will be called here
	err := fn(ctx)
	// implement your custom retry logic based on the error
	return err
}

func (r *CustomRetry) SetMaxRetries(n int) {
	r.MaxRetries = n
}

// ExampleConfig_WithRetryer shows how a custom retry function can be created and passed in the config object.
func ExampleConfig_WithRetryer() {
	retryer := CustomRetry{MaxRetries: 3}
	c := config.NewConfig("apiKey").WithRetryer(&retryer)

	s := chatai.NewService(c)
	s.AskAI("some question")
}
