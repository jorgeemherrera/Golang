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
	ErrUserNotFound       = errors.New("User not found")
	ErrRoleAlreadyAdded   = errors.New("Role was already added for this user")
	ErrRoleNotFound       = errors.New("Role not found")
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
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
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

func (srv *serv) AddUserRole(ctx context.Context, userID, RoleID int64) error {

	roles, err := srv.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.RoleID == RoleID {
			return ErrRoleAlreadyAdded
		}
	}

	return srv.repo.SaveUserRole(ctx, userID, RoleID)
}

func (srv *serv) RemoveUserRole(ctx context.Context, userID, RoleID int64) error {
	roles, err := srv.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}
	roleFound := false
	for _, r := range roles {
		if r.RoleID == RoleID {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}
	return srv.repo.RemoveUserRole(ctx, userID, RoleID)
}
