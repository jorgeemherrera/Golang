package service

import (
	context "context"
	"errors"

	models "github.com/jorgeemherrera/Golang/internal/models"
)

func (srv *serv) GetProducts(ctx context.Context) ([]models.Product, error) {

	//Ejecutar el query
	prods, err := srv.repo.GetProducts(ctx)

	if err != nil {
		return nil, err
	}

	products := []models.Product{}

	for _, p := range prods {
		products = append(products, models.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}
	return products, nil
}
func (srv *serv) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	prod, err := srv.repo.GetProduct(ctx, id)

	if err != nil {
		return nil, err
	}

	product := &models.Product{
		ID:          prod.ID,
		Name:        prod.Name,
		Description: prod.Description,
		Price:       prod.Price,
	}

	return product, nil
}

var validRolesToAddProduct []int64 = []int64{1, 2}
var ErrInvalidPermissions = errors.New("User does not have permission to add a Product")

func (srv *serv) AddProduct(ctx context.Context, product models.Product, email string) error {

	userEmail, err := srv.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	roles, err := srv.repo.GetUserRoles(ctx, userEmail.ID)
	if err != nil {
		return err
	}

	userCanAddProduct := false

	for _, r := range roles {
		for _, validRole := range validRolesToAddProduct {
			if validRole == r.RoleID {
				userCanAddProduct = true
			}
		}
	}

	if !userCanAddProduct {
		return ErrInvalidPermissions
	}
	return srv.repo.SaveProduct(
		ctx,
		product.Name,
		product.Description,
		product.Price,
		userEmail.ID,
	)
}
