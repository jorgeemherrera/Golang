package repository

import (
	context "context"

	entity "github.com/jorgeemherrera/Golang/internal/entity"
)

const (
	queryInsertProduct = `
	INSERT INTO PRODUCTS(name, description, price, created_by) values (?,?,?,?);
	`

	queryGetAllProducts = `
	SELECT 
	id,
	name,
	description,
	price
	FROM PRODUCTS;
	`

	queryGetProductById = `
	SELECT 
	id,
	name,
	description,
	price
	FROM PRODUCTS
	where id = ?;
	`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, queryInsertProduct, name, description, price, createdBy)
	return err
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {
	prods := []entity.Product{}

	err := r.db.SelectContext(ctx, &prods, queryGetAllProducts)
	if err != nil {
		return nil, err
	}
	return prods, nil
}

func (r *repo) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	prod := &entity.Product{}

	err := r.db.GetContext(ctx, prod, queryGetProductById, id)
	if err != nil {
		return nil, err
	}
	return prod, nil
}
