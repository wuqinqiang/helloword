package gpt3

import (
	"context"
	"errors"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"io"
)

type Client struct {
	*gogpt.Client
}

func NewClient(token string) *Client {
	gpt := gogpt.NewClient(token)
	return &Client{gpt}
}

func (c *Client) Send(ctx context.Context, text string) (string, error) {
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 1000,
		Prompt:    text,
		Stream:    true,
	}
	stream, err := c.CreateCompletionStream(ctx, req)
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
	return res, nil
}
