package eleven_labs

type TextToSpeechInput struct {
	ModelID         string  `json:"model_id"`
	Text            string  `json:"text"`
	Stability       float32 `json:"stability,omitempty"`
	SimilarityBoost float32 `json:"similarity_boost,omitempty"`
}
