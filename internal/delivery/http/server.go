package http

import (
	"database/sql"
	"log"
	"net/http"

	"test-api-arch/config"
	"test-api-arch/internal/repository"
	"test-api-arch/internal/usecase"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.Config, dbConn *sql.DB) *Server {
	// Initialize repository and usecase
	userRepo := repository.NewUserRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := NewUserHandler(userUsecase)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.RegisterUser)

	// Add middleware (logging, recovery)
	handler := LoggingMiddleware(RecoveryMiddleware(mux))

	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Start() {
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
