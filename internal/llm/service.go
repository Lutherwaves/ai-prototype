// internal/llm/service.go
package llm

import (
	"context"
	"go.uber.org/zap"
	"imagine-proto/internal/core/domain"
	"imagine-proto/internal/llm/provider"
)

type Service struct {
	provider provider.Provider
	logger   *zap.Logger
}

func NewService(p provider.Provider, logger *zap.Logger) *Service {
	return &Service{
		provider: p,
		logger:   logger,
	}
}

func (s *Service) ProcessMessage(ctx context.Context, thread *domain.Thread, msg *domain.Message) (*domain.Message, error) {
	messages := make([]provider.Message, len(thread.Messages)+1)
	for i, m := range thread.Messages {
		messages[i] = provider.Message{
			Role:    m.Role,
			Content: m.Content,
		}
	}
	messages[len(messages)-1] = provider.Message{
		Role:    msg.Role,
		Content: msg.Content,
	}

	response, err := s.provider.ProcessMessage(ctx, messages)
	if err != nil {
		s.logger.Error("failed to process message with provider",
			zap.Error(err),
			zap.String("provider", string(s.provider.Name())),
		)
		return nil, err
	}

	return &domain.Message{
		Role:    response.Role,
		Content: response.Content,
	}, nil
}
