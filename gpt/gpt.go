package gpt

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/qingconglaixueit/wechatbot/config"
	"github.com/sashabaranov/go-openai"
)

func Chat(msg []openai.ChatCompletionMessage) (string, error) {
	cfg := config.LoadConfig()
	if cfg.GPTApiKey == "" {
		log.Printf("GPT api key required\n")
		return "", errors.New("GPT api key required")
	}
	//代理可通过文件配置
	var client = openai.NewClient(cfg.GPTApiKey)
	if cfg.Proxy {
		config := openai.DefaultConfig(cfg.GPTApiKey)
		proxyUrl, err := url.Parse(cfg.ProxyUrl) //"http://localhost:1080"
		if err != nil {
			panic(err)
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		config.HTTPClient = &http.Client{
			Transport: transport,
		}
		client = openai.NewClientWithConfig(config)
	}

	log.Printf("Request already send")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            cfg.Model,
			Messages:         msg,
			MaxTokens:        cfg.MaxTokens,
			Temperature:      cfg.Temperature,
			TopP:             1,
			FrequencyPenalty: 0,
			PresencePenalty:  0,
		},
	)
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}
	content := resp.Choices[0].Message.Content
	log.Printf("GPT Response: %s\n", content)
	return content, nil
}
