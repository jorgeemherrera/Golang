package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/jorgeemherrera/Golang/internal/entity"
)

// Repository is the interface that wraps the basic CRUD operations
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
