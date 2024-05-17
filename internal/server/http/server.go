package http

import (
	productHttp "github.com/devkemc/challenge-eulabs/internal/product/infraestructure/port/http"
	"github.com/devkemc/challenge-eulabs/pkg/config"
	"github.com/devkemc/challenge-eulabs/pkg/db/product"
	"github.com/labstack/echo/v4"
)

type Server struct {
	engine  *echo.Echo
	queries *product.Queries
	config  *config.Schema
}

func NewServer(queries *product.Queries) *Server {
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
	productHttp.Routes(s.engine.Group("/v1"), s.queries)
	return nil
}
