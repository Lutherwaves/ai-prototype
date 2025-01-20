// cmd/api/main.go
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"imagine-proto/internal/api"
	"imagine-proto/internal/llm"
	"imagine-proto/internal/llm/provider"
	"imagine-proto/internal/platform/config"
	"imagine-proto/internal/platform/logger"

	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	log := logger.NewLogger()
	defer log.Sync()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config", zap.Error(err))
	}

	providerConfig := provider.Config{
		Type:        cfg.LLM.Type,
		BaseURL:     cfg.LLM.BaseURL,
		Model:       cfg.LLM.Model,
		MaxTokens:   cfg.LLM.MaxTokens,
		Temperature: cfg.LLM.Temperature,
	}

	// Set API key based on provider type
	switch cfg.LLM.Type {
	case provider.OpenAI:
		providerConfig.APIKey = os.Getenv("OPENAI_API_KEY")
	case provider.Perplexity:
		log.Info("Choosing perplexity...")
		providerConfig.APIKey = os.Getenv("PERPLEXITY_API_KEY")
	}

	llmProvider, err := llm.NewProvider(providerConfig)
	if err != nil {
		log.Fatal("failed to create LLM provider", zap.Error(err))
	}

	llmService := llm.NewService(llmProvider, log)
	server := api.NewServer(cfg, llmService, log)

	// Handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Info("shutting down server...")

		os.Exit(1)
		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Error("server shutdown error", zap.Error(err))
		}
	}()

	log.Info("starting server", zap.String("port", cfg.Server.Port))
	if err := server.Start(); err != nil {
		log.Fatal("server error", zap.Error(err))
	}
}
