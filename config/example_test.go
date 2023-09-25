package config_test

import (
	"github.com/nirdosh17/go-sdk-template/api/chatai"
	"github.com/nirdosh17/go-sdk-template/config"
)

func Example() {
	// use this for default working config
	// c := config.NewConfig("apiKey")

	c := config.NewConfig("apiKey").
		WithMaxRetries(5).
		WithEndpoint("https://region2.serviceapi.com")

	s := chatai.NewService(c)
	s.AskAI("some question")
}
