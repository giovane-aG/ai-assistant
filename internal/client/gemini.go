package client

import (
	"context"
	"fmt"

	"github.com/giovane-aG/ai-assistant/internal/model"
	"google.golang.org/genai"
)

type GeminiClient struct {
	Client *genai.Client
}

func NewGeminiClient(ctx context.Context, apiKey string) (*GeminiClient, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: apiKey})
	if err != nil {
		return nil, err
	}
	return &GeminiClient{Client: client}, nil
}

func (c *GeminiClient) GenerateContent(ctx context.Context, aiModel string, content string) (*model.GenerateContentResponse, error) {
	result, err := c.Client.Models.GenerateContent(ctx, aiModel, genai.Text(content), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	return &model.GenerateContentResponse{Content: result.Text()}, nil
}
