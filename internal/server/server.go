package server

import (
	"context"
	"github.com/find_property/internal/config"
	v1 "github.com/find_property/internal/delivery/v1"
	"github.com/find_property/pkg/logger"

	"github.com/rs/cors"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handlers *v1.Handler) *Server {

	router := NewRouter(handlers)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*",
		},
	})

	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.HTTP.Port,
			Handler: corsOpts.Handler(router),
		},
	}
}

func (s *Server) Run() error {
	logger.LogInfo("Restart server")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logger.LogInfo("Shutdown server")
	return s.httpServer.Shutdown(ctx)
}
