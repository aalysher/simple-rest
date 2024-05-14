package storage

import (
	"fmt"
	"log"

	"github.com/aalysher/simple-rest/internal/config"
	"github.com/jmoiron/sqlx"
)

func NewStorage(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", cfg.Address())
	if err != nil {
		return nil, fmt.Errorf("database open: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL")

	return db, nil
}
