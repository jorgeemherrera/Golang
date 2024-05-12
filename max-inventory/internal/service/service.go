package service

import (
	"context"

	"github.com/jorgeemherrera/Golang/internal/models"
	"github.com/jorgeemherrera/Golang/internal/repository"
)

// Service is the business logic of the application
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	AddUserRole(ctx context.Context, userID, RoleID int64) error
	RemoveUserRole(ctx context.Context, userID, RoleID int64) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
