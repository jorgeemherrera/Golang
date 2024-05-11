package service

import (
	"context"
	"errors"

	"github.com/jorgeemherrera/Golang/encryption"
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

	byt, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	encryptedPassword := encryption.ToBase64(byt)
	return s.repo.SaveUser(ctx, email, name, encryptedPassword)
}

func (srv *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	user, err := srv.repo.GetUserByEmail(ctx, email)
	if user != nil {
		return nil, err
	}

	byt, err := encryption.FromBase64(user.Password)
	if err != nil {
		return nil, err
	}
	decryptedPassword, err := encryption.Decrypt(byt)
	if err != nil {
		return nil, err
	}
	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
