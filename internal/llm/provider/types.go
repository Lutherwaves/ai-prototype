// internal/llm/provider/types.go
package provider

import (
	"context"
)

type ProviderType string

const (
	OpenAI     ProviderType = "openai"
	Perplexity ProviderType = "perplexity"
)

type Provider interface {
	ProcessMessage(ctx context.Context, messages []Message) (*Message, error)
	Name() ProviderType
}

type Config struct {
	Type        ProviderType `mapstructure:"type"`
	BaseURL     string       `mapstructure:"baseUrl"`
	Model       string       `mapstructure:"model"`
	MaxTokens   int          `mapstructure:"maxTokens"`
	Temperature float64      `mapstructure:"temperature"`
	APIKey      string
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name,omitempty"`
}

type Response struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
