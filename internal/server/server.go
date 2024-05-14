package server

import (
	"fmt"

	"github.com/aalysher/simple-rest/internal/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	cfg *config.Config
	e   *echo.Echo
}

func NewServer(cfg *config.Config) (*Server, error) {
	e := echo.New()

	s := &Server{
		e:   e,
		cfg: cfg,
	}

	return s, nil
}

func (s *Server) Start() error {
	// s.logger.Info(fmt.Sprintf("start http server: %s", s.cfg.Address()))
	return s.e.Start(fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port))
}