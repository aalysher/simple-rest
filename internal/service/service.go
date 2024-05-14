package service

import (
	"github.com/aalysher/simple-rest/internal/config"
	"github.com/aalysher/simple-rest/internal/server"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	cfg    *config.Config
	db     *sqlx.DB
	server *server.Server
}

func New() error {
	s := &Service{}

	if err := s.init(); err != nil {
		return err
	}

	return nil
}

func (s *Service) init() (err error) {
	s.cfg, err = config.LoadConfig()
	if err != nil {
		return err
	}

	s.server, err = server.NewServer(s.cfg)
	if err != nil {
		return err
	}

	if err = s.server.Start(); err != nil {
		return err
	}

	return nil
}
