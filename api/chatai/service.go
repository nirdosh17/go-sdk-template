package chatai

import (
	"context"

	"github.com/nirdosh17/go-sdk-template/client"
	"github.com/nirdosh17/go-sdk-template/config"
)

const (
	serviceName    = "chatai"
	MaxInputLength = 200
)

// ChatAPI exposes APIs related to chatAI service.
type ChatAPI struct {
	Config *config.Config
}

type AIAnswer struct {
	Answer          string  `json:"answer"`
	ConfidenceScore float32 `json:"confidenceScore"`
}

// NewService returns an instance of ChatAPI service.
//
// Example:
//
//	// with default configs
//	ai := chatai.New(api.NewConfig())
//
//	// With custom configs:
//	ai := chatai.NewService(
//		config.NewConfig().WithMaxRetries(3),
//	)
func NewService(c *config.Config) *ChatAPI {
	return &ChatAPI{Config: c}
}

type question struct {
	query string `json:"query"`
}

// AskAIWithContext provides answer for input question from ChatAI service.
func (c *ChatAPI) AskAIWithContext(ctx context.Context, input string) (AIAnswer, error) {
	var answer AIAnswer

	// blank answer for blank question
	if input == "" {
		return answer, nil
	}

	if len(input) > MaxInputLength {
		return answer, ErrInputSizeLimitExceeded
	}

	// configure HTTP client
	url := c.Config.Endpoint + "/" + serviceName

	req := client.Request{Client: c.Config.HTTPClient, Logger: c.Config.Logger, Debug: c.Config.Debug}
	q := question{query: input}

	err := c.Config.Retryer.Run(ctx, func(ctx context.Context) error {
		return req.Perform(ctx, url, "POST", q, &answer)
	})

	return answer, err
}

func (c *ChatAPI) AskAI(question string) (AIAnswer, error) {
	return c.AskAIWithContext(context.Background(), question)
}
