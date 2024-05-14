package service

import (
	"os"
	"testing"

	"github.com/jorgeemherrera/Golang/encryption"
	"github.com/jorgeemherrera/Golang/internal/entity"
	"github.com/jorgeemherrera/Golang/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository

var srv Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("ValidPassword"))
	encrytedPassword := encryption.ToBase64(validPassword)
	user := &entity.User{Email: "test@exists.com", Password: encrytedPassword}
	adminUser := &entity.User{ID: 1, Email: "admin@email.com", Password: encrytedPassword}
	customerUser := &entity.User{ID: 2, Email: "customer@email.com", Password: encrytedPassword}

	//Mock Repo
	repo = &repository.MockRepository{}

	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(user, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("GetUserRoles", mock.Anything, int64(1)).Return([]entity.UserRole{
		{
			UserID: 1,
			RoleID: 1,
		},
	}, nil)

	repo.On("GetUserByEmail", mock.Anything, "admin@email.com").Return(adminUser, nil)
	repo.On("GetUserByEmail", mock.Anything, "customer@email.com").Return(customerUser, nil)
	repo.On("GetUserRoles", mock.Anything, int64(2)).Return([]entity.UserRole{
		{
			UserID: 2,
			RoleID: 3,
		},
	}, nil)
	repo.On("SaveProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	srv = New(repo)

	code := m.Run()
	os.Exit(code)
}
