package gpt3

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var prompt = "使用单词：%s 写一篇英语短文。短文最后标注这几个单词的中文意思"

type Client struct {
	*gogpt.Client
}

func NewClient(token string, proxyUrl string) (*Client, error) {
	conf := gogpt.DefaultConfig(token)

	if proxyUrl != "" {
		proxy, err := url.Parse(proxyUrl)
		if err != nil {
			return nil, err
		}
		dialer := net.Dialer{
			Timeout: 30 * time.Second,
		}
		transport := &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			DialContext:     dialer.DialContext,
			MaxIdleConns:    50,
			IdleConnTimeout: 60 * time.Second,
		}
		conf.HTTPClient.Transport = transport
	}
	gpt := gogpt.NewClientWithConfig(conf)
	return &Client{gpt}, nil
}

func (client *Client) Generate(ctx context.Context, words []string) (phrase string, err error) {
	wordStr := strings.Join(words, ",")
	return client.request(ctx, fmt.Sprintf(prompt, wordStr))
}

func (client *Client) request(ctx context.Context, text string) (string, error) {

	req := gogpt.ChatCompletionRequest{
		Model: gogpt.GPT3Dot5Turbo,
		Messages: []gogpt.ChatCompletionMessage{
			{
				Role:    "user",
				Content: text,
			},
		},
		MaxTokens: 400,
	}

	var (
		resp gogpt.ChatCompletionResponse
		err  error
	)

	for i := 0; i < 3; i++ {
		resp, err = client.CreateChatCompletion(ctx, req)
		if err == nil {
			if len(resp.Choices) == 0 {
				return "", err
			}
			fmt.Println("Generate:", resp.Choices[0].Message.Content)
			return resp.Choices[0].Message.Content, nil
		}
	}

	return "", err
}
