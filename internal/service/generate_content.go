package service

import (
	"context"

	"github.com/giovane-aG/ai-assistant/internal/model"
)

type IGenerateContentClient interface {
	GenerateContent(
		ctx context.Context,
		model string,
		content string,
	) (*model.GenerateContentResponse, error)
}

type GenerateContentService struct {
	generateContentClient IGenerateContentClient
}

func NewGenerateContentService(generateContentClient IGenerateContentClient) *GenerateContentService {
	return &GenerateContentService{
		generateContentClient: generateContentClient,
	}
}

func (s *GenerateContentService) Generate(ctx context.Context, model string, content string) (*model.GenerateContentResponse, error) {
	return s.generateContentClient.GenerateContent(ctx, model, content)
}
