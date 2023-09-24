// chatai_test includes examples for the service
package chatai_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/nirdosh17/go-sdk-template/api/chatai"
	"github.com/nirdosh17/go-sdk-template/config"
)

func Example() {
	// with default configs
	ai := chatai.NewService(config.NewConfig())
	ans, err := ai.AskAI("tell me one memory optimization technique in Go")
	if err != nil {
		// handle err
	}
	fmt.Println("Answer: '", ans.Answer, "'| Confidence score:", ans.ConfidenceScore)
}

func ExampleChatAPI_AskAI() {
	// with default configs
	ai := chatai.NewService(config.NewConfig())
	ans, err := ai.AskAI("memory optimization technique in Go")
	if err != nil {
		// handle err
	}
	fmt.Println("Answer: '", ans.Answer, "'| Confidence score:", ans.ConfidenceScore)
}

func ExampleChatAPI_AskAIWithContext() {
	ai := chatai.NewService(config.NewConfig())

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ans, err := ai.AskAIWithContext(ctx, "When will the world end?")
	if err != nil {
		// handle err
	}
	fmt.Println("Answer: ", ans.Answer, "Confidence score:", ans.ConfidenceScore)
}

func ExampleChatAPI_AskAIWithContext_withAdditionalConfigs() {
	proxyURL, err := url.Parse("https://example.com")
	if err != nil {
		log.Println("error parsing proxy url", err)
	}
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	// custom client with proxy settings and timeout
	customClient := &http.Client{Transport: transport, Timeout: 10 * time.Second}

	// config object with custom settings
	config := config.NewConfig().
		WithHTTPClient(customClient).
		WithMaxRetries(3)

	ai := chatai.NewService(config)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ans, err := ai.AskAIWithContext(ctx, "When will the world end?")
	if err != nil {
		// handle err
	}
	fmt.Println("Answer: ", ans.Answer, "Confidence score:", ans.ConfidenceScore)
}
