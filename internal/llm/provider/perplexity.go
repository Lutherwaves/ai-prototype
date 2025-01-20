// internal/llm/provider/perplexity.go
package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PerplexityProvider struct {
	config     Config
	httpClient *http.Client
}

func NewPerplexityProvider(cfg Config) *PerplexityProvider {
	return &PerplexityProvider{
		config: cfg,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (p *PerplexityProvider) Name() ProviderType {
	return Perplexity
}

func (p *PerplexityProvider) ProcessMessage(ctx context.Context, messages []Message) (*Message, error) {
	payload := map[string]interface{}{
		"model":       p.config.Model,
		"messages":    messages,
		"temperature": p.config.Temperature,
		"max_tokens":  p.config.MaxTokens,
	}

	resp, err := p.sendRequest(ctx, "/chat/completions", payload)
	if err != nil {
		return nil, err
	}

	return &resp.Choices[0].Message, nil
}

func (p *PerplexityProvider) sendRequest(ctx context.Context, endpoint string, payload interface{}) (*Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("%s%s", p.config.BaseURL, endpoint),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	fmt.Print(fmt.Sprintf("Bearer %s"))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.config.APIKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
