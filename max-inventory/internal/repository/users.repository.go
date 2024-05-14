package repository

import (
	"context"

	"github.com/jorgeemherrera/Golang/internal/entity"
)

const (
	queryInsertUser = `
	INSERT INTO USERS (email, name, password)
	VALUES (?, ?, ?);
	`
	queryGetUserByEmail = `
	SELECT 
		id, 
		email,
		name,
		password
	FROM USERS
	WHERE email=?;
	`

	queryInsertUserRole = `
	INSERT INTO USER_ROLES (user_id, role_id) values (:user_id, :role_id);
	`

	queryRemoveUserRole = `
	DELETE FROM USER_ROLES WHERE user_id = :user_id  and role_id = :role_id;
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	// Ejecutar el queryInsertUser
	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	// crear el destination para GetContext
	userDestination := &entity.User{}
	// Ejecutar el queryGetUserByEmail
	err := r.db.GetContext(ctx, userDestination, queryGetUserByEmail, email)

	return userDestination, err
}

func (r *repo) SaveUserRole(ctx context.Context, userID, RoleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: RoleID,
	}
	_, err := r.db.NamedExecContext(ctx, queryInsertUserRole, data)

	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, RoleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: RoleID,
	}
	_, err := r.db.NamedExecContext(ctx, queryRemoveUserRole, data)

	return err
}
func (r *repo) GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error) {
	roles := []entity.UserRole{}
	err := r.db.SelectContext(ctx, &roles, "select user_id, role_id from USER_ROLES where user_id = ?", userID)

	if err != nil {
		return nil, err
	}
	return roles, nil
}
