package usecase

import (
	"context"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/repo"
)

type DeleteProduct struct {
	productRepo repo.ProductRepo
}

func NewDeleteProduct(
	productRepo repo.ProductRepo,
) *DeleteProduct {
	return &DeleteProduct{
		productRepo: productRepo,
	}
}

func (l DeleteProduct) Execute(
	id string,
) error {
	product, err := l.productRepo.GetProductByID(context.Background(), id)
	if err != nil {
		return entity.NewErr(err)
	}

	if product.ID == "" {
		return entity.ErrProductNotFound
	}

	if err := l.productRepo.DeleteProduct(context.Background(), id); err != nil {
		return entity.NewErr(err)
	}

	return nil
}
