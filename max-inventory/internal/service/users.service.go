package service

import (
	"context"
	"errors"

	"github.com/jorgeemherrera/Golang/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("User already exists")
	ErrInvalidCredentials = errors.New("Invalid Password")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	user, _ := s.repo.GetUserByEmail(ctx, email)
	if user != nil {
		return ErrUserAlreadyExists
	}
	//TODO: hash password
	return s.repo.SaveUser(ctx, email, name, password)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if user != nil {
		return nil, err
	}
	//TODO: decrypt password
	if user.Password != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
