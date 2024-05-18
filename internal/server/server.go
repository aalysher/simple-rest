package server

import (
	"fmt"

	"github.com/aalysher/simple-rest/internal/config"
	"github.com/aalysher/simple-rest/internal/http/handler"
	"github.com/aalysher/simple-rest/internal/storage"
	"github.com/labstack/echo/v4"
)

type Server struct {
	cfg      *config.Config
	e        *echo.Echo
	userRepo storage.UserRepository
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
	userHandler := handler.NewUserHandler(s.userRepo)
	s.e.POST("/users", userHandler.CreateUser)
	// s.e.GET("/users/:id", userHandler.GetUser)
	// s.e.PATCH("/users/:id", userHandler.UpdateUser)

	return s.e.Start(fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port))
}
