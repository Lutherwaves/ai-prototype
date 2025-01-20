// internal/core/ports/repositories.go
package ports

import (
	"context"
	"imagine-proto/internal/core/domain"
)

type ThreadRepository interface {
	Create(ctx context.Context, thread *domain.Thread) error
	Get(ctx context.Context, id string) (*domain.Thread, error)
	Update(ctx context.Context, thread *domain.Thread) error
	Delete(ctx context.Context, id string) error
}
