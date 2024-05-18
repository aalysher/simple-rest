package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aalysher/simple-rest/internal/config"
	"github.com/aalysher/simple-rest/internal/models"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(id string) (*models.User, error)
	EditUser(user *models.User) error
}

type UserStorage struct {
	db *sql.DB
}

func NewStorage(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Address())
	if err != nil {
		return nil, fmt.Errorf("database open: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL")

	return db, nil
}

func CreateUser(user *models.User) error {
	return nil
}
