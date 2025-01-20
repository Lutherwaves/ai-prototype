// internal/api/router.go
package api

import (
	"context"
	"net/http"

	"imagine-proto/internal/api/handlers"
	"imagine-proto/internal/api/middleware"
	"imagine-proto/internal/llm"
	"imagine-proto/internal/platform/config"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Server struct {
	router     *chi.Mux
	config     *config.Config
	logger     *zap.Logger
	llmService *llm.Service
}

func NewServer(cfg *config.Config, llmService *llm.Service, logger *zap.Logger) *Server {
	s := &Server{
		router:     chi.NewRouter(),
		config:     cfg,
		logger:     logger,
		llmService: llmService,
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

func (s *Server) setupMiddleware() {
	s.router.Use(chimiddleware.RequestID)
	s.router.Use(chimiddleware.RealIP)
	s.router.Use(middleware.Logging(s.logger))
	s.router.Use(chimiddleware.Recoverer)
}

func (s *Server) setupRoutes() {
	chatHandler := handlers.NewChatHandler(s.llmService, s.logger)

	s.router.Get("/health", handlers.HealthCheck)
	s.router.Post("/chat", chatHandler.HandleMessage)
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.config.Server.Port, s.router)
}

func (s *Server) Shutdown(ctx context.Context) error {
	// Implement any cleanup needed
	return nil
}
