// internal/llm/factory.go
package llm

import (
	"fmt"
	"imagine-proto/internal/llm/provider"
)

func NewProvider(cfg provider.Config) (provider.Provider, error) {
	switch cfg.Type {
	case provider.OpenAI:
		return provider.NewOpenAIProvider(cfg), nil
	case provider.Perplexity:
		return provider.NewPerplexityProvider(cfg), nil
	default:
		return nil, fmt.Errorf("unsupported provider type: %s", cfg.Type)
	}
}
