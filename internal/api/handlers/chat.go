// internal/api/handlers/chat.go
package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
	"imagine-proto/internal/core/domain"
	"imagine-proto/internal/llm"
)

type ChatHandler struct {
	llmService *llm.Service
	logger     *zap.Logger
}

func NewChatHandler(service *llm.Service, logger *zap.Logger) *ChatHandler {
	return &ChatHandler{
		llmService: service,
		logger:     logger,
	}
}

func (h *ChatHandler) HandleMessage(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ThreadID string `json:"thread_id"`
		Message  string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Error("failed to decode request", zap.Error(err))
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	thread := &domain.Thread{
		ID: input.ThreadID,
		Messages: []domain.Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant for house design and construction.",
			},
		},
	}

	msg := &domain.Message{
		Role:    "user",
		Content: input.Message,
	}

	response, err := h.llmService.ProcessMessage(r.Context(), thread, msg)
	if err != nil {
		h.logger.Error("failed to process message",
			zap.Error(err),
			zap.String("thread_id", input.ThreadID),
		)
		http.Error(w, "processing failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
