package http

import (
	"github.com/devkemc/challenge-eulabs/internal/database"
	"github.com/devkemc/challenge-eulabs/pkg/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	engine  *echo.Echo
	queries *database.Queries
	config  *config.Schema
}

func NewServer(queries *database.Queries) *Server {
	return &Server{
		engine:  echo.New(),
		queries: queries,
		config:  config.GetConfig(),
	}
}

func (s *Server) Start() {
	if err := s.MapRoutes(); err != nil {
		s.engine.Logger.Fatal(err)
		panic(err)
	}
	s.engine.Logger.Info("Http server started on port " + s.config.HttpPort)
	if err := s.engine.Start(s.config.HttpPort); err != nil {
		s.engine.Logger.Fatal(err)
		panic(err)
	}
}

func (s *Server) MapRoutes() error {
	s.engine.Group("/v1")
	return nil
}
