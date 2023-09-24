package chatai

import (
	"context"

	"github.com/nirdosh17/go-sdk-template/client"
	"github.com/nirdosh17/go-sdk-template/config"
)

const (
	ServiceName    = "chatai"
	MaxInputLength = 200
)

// exposes APIs related to chatAI service
type ChatAPI struct {
	Config *config.Config
}

type AIAnswer struct {
	Answer          string  `json:"answer"`
	ConfidenceScore float32 `json:"confidenceScore"`
}

// Example:
//
//	// with default configs
//	ai := chatai.New(api.NewConfig())
//
//	// With custom configs:
//	ai := chatai.New(
//		config.NewConfig().WithHTTPTimeout(5 * time.Second).WithMaxRetries(1),
//	)
func NewService(c *config.Config) *ChatAPI {
	return &ChatAPI{Config: c}
}

type Question struct {
	Query string `json:"query"`
}

// AskAIWithContext provides answer for input question from ChatAI service.
func (c *ChatAPI) AskAIWithContext(ctx context.Context, question string) (AIAnswer, error) {
	var answer AIAnswer

	// blank answer for blank question
	if question == "" {
		return answer, nil
	}

	if len(question) > MaxInputLength {
		return answer, ErrInputSizeLimitExceeded
	}

	// configure HTTP client
	url := c.Config.Endpoint + "/" + ServiceName

	req := client.Request{Client: c.Config.HTTPClient, Logger: c.Config.Logger, Debug: c.Config.Debug}
	q := Question{Query: question}

	err := c.Config.Retryer.Run(ctx, func(ctx context.Context) error {
		return req.Perform(ctx, url, "POST", q, &answer)
	})

	return answer, err
}

func (c *ChatAPI) AskAI(question string) (AIAnswer, error) {
	return c.AskAIWithContext(context.Background(), question)
}
