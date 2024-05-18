package service

import (
	"database/sql"
	"fmt"

	"github.com/aalysher/simple-rest/internal/config"
	"github.com/aalysher/simple-rest/internal/server"
	"github.com/aalysher/simple-rest/internal/storage"
)

type Service struct {
	cfg    *config.Config
	db     *sql.DB
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
		return fmt.Errorf("error while load config: %w", err)
	}

	s.server, err = server.NewServer(s.cfg)
	if err != nil {
		return fmt.Errorf("error while run server: %w", err)
	}

	s.db, err = storage.NewStorage(s.cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err = s.server.Start(); err != nil {
		return fmt.Errorf("error while starting web server: %w", err)
	}

	return nil
}
