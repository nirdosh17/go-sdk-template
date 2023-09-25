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
	"github.com/nirdosh17/go-sdk-template/test"
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
	json := `{"answer":"in 50 years","confidenceScore":80}`
	c := test.MockHTTPClient{
		JSONBody:   &json,
		StatusCode: 200,
	}
	ai := chatai.NewService(config.NewConfig().WithHTTPClient(&c))

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	ans, err := ai.AskAIWithContext(ctx, "When will the world end?")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Answer: %v | Confidence Score: %v", ans.Answer, ans.ConfidenceScore)
	// Output:
	// Answer: in 50 years | Confidence Score: 80
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
