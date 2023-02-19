package gpt3

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var EmptyErr = errors.New("the result is empty")

var prompt = "请你用以下单词：%s 写一篇英语短文。此外，在生成的短文后面，说明以上单词的中文意思"

type Client struct {
	*gogpt.Client
}

func NewClient(token string) *Client {
	gpt := gogpt.NewClient(token)
	return &Client{gpt}
}

func (client *Client) Generate(ctx context.Context, words []string) (phrase string, err error) {
	wordStr := strings.Join(words, ",")
	return client.request(ctx, fmt.Sprintf(prompt, wordStr))
}

func (client *Client) request(ctx context.Context, text string) (string, error) {
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 1000,
		Prompt:    text,
		Stream:    true,
	}
	stream, err := client.CreateCompletionStream(ctx, req)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	var res string
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			break
		}
		if len(response.Choices) > 0 {
			res += response.Choices[0].Text
		}
	}
	if res == "" {
		return res, EmptyErr
	}
	return res, nil
}
