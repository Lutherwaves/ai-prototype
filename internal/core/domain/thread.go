// internal/core/domain/thread.go
package domain

import (
	"time"
)

type Thread struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	Messages  []Message `json:"messages"`
}

type Message struct {
	Role     string                 `json:"role"`
	Content  string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
