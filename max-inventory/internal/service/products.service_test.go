package service

import (
	"context"
	"testing"

	"github.com/jorgeemherrera/Golang/internal/models"
)

func TestAddProduct(t *testing.T) {
	testCases := []struct {
		Name          string
		Product       models.Product
		Email         string
		ExpectedError error
	}{
		{
			Name: "AddProduct_Success",
			Product: models.Product{
				ID:          1,
				Name:        "Test Product",
				Description: "This is a product description",
				Price:       10.00,
			},
			Email:         "admin@email.com",
			ExpectedError: nil,
		},
		{
			Name: "AddProduct_InvalidPermissions",
			Product: models.Product{
				ID:          1,
				Name:        "Test Product",
				Description: "This is a product description",
				Price:       10.00,
			},
			Email:         "customer@email.com",
			ExpectedError: ErrInvalidPermissions,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := srv.AddProduct(ctx, tc.Product, tc.Email)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
