package model

type GenerateContentResponse struct {
	Content string `json:"content"`
}

func (g *GenerateContentResponse) Text() string {
	return g.Content
}
