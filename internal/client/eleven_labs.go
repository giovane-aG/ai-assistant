package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/giovane-aG/ai-assistant/internal/model/eleven_labs"
)

type ElevenLabsClient struct {
	apiKey string
	client *http.Client
}

func NewElevenLabsClient(apiKey string) *ElevenLabsClient {
	return &ElevenLabsClient{
		apiKey: os.Getenv("ELEVEN_LABS_API_KEY"),
		client: &http.Client{},
	}
}
func (c *ElevenLabsClient) GenerateSpeech(text string) error {
	bodyJSON, err := json.Marshal(eleven_labs.TextToSpeechInput{
		ModelID:         "eleven_flash_v2_5",
		Text:            text,
		Stability:       0.5,
		SimilarityBoost: 0.5,
	})

	if err != nil {
		return fmt.Errorf("failed to marshal body: %w", err)
	}

	byteBody := bytes.NewBuffer(bodyJSON)

	req, err := http.NewRequest("POST", "https://api.elevenlabs.io/v1/text-to-speech/EXAVITQu4vr4xnSDxMaL", byteBody)

	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("xi-api-key", c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to generate speech: %s", resp.Status)
	}

	output, err := os.Create("output.mp3")
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer output.Close()

	_, err = io.Copy(output, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy response body to file: %w", err)
	}

	return nil
}
