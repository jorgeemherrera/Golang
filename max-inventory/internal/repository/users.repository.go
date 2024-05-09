package repository

import (
	"context"

	"github.com/jorgeemherrera/Golang/internal/entity"
)

const (
	queryInsertUser = `
	INSERT INTO users (email, name, password)
	VALUES (?, ?,?);
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
