package service

import (
	"context"
	"os"
	"testing"

	"github.com/jorgeemherrera/Golang/encryption"
	"github.com/jorgeemherrera/Golang/internal/entity"
	"github.com/jorgeemherrera/Golang/internal/repository"
	"github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository

var srv Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("ValidPassword"))
	encrytedPassword := encryption.ToBase64(validPassword)
	user := &entity.User{Email: "test@exists.com", Password: encrytedPassword}

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

	srv = New(repo)

	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			Email:         "test@test.com",
			UserName:      "test",
			Password:      "ValidPassword",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_UserAlreadyExist",
			Email:         "test@exists.com",
			UserName:      "user1",
			Password:      "ValidPassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := srv.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "LoginUser_Success",
			Email:         "test@exists.com",
			Password:      "ValidPassword",
			ExpectedError: nil,
		},
		{
			Name:          "LoginUser_InvalidPassword",
			Email:         "test@exists.com",
			Password:      "InvalidPassword",
			ExpectedError: ErrInvalidCredentials,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := srv.LoginUser(ctx, tc.Email, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestAddUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserID        int64
		RoleID        int64
		ExpectedError error
	}{
		{
			Name:          "AddUserRole_Success",
			UserID:        1,
			RoleID:        2,
			ExpectedError: nil,
		},
		{
			Name:          "AddUserRole_UserAlreadyHasRole",
			UserID:        1,
			RoleID:        1,
			ExpectedError: ErrRoleAlreadyAdded,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := srv.AddUserRole(ctx, tc.UserID, tc.RoleID)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestRemoveUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserID        int64
		RoleID        int64
		ExpectedError error
	}{
		{
			Name:          "RemoveUserRole_Success",
			UserID:        1,
			RoleID:        1,
			ExpectedError: nil,
		},
		{
			Name:          "RemoveUserRole_UserDoesNotHaveRole",
			UserID:        1,
			RoleID:        3,
			ExpectedError: ErrRoleNotFound,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := srv.RemoveUserRole(ctx, tc.UserID, tc.RoleID)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
