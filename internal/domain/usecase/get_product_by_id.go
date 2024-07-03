package usecase

import (
	"context"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/repo"
)

type GetProductByID struct {
	productRepo repo.ProductRepo
}

func NewGetProductByID(
	productRepo repo.ProductRepo,
) *GetProductByID {
	return &GetProductByID{
		productRepo: productRepo,
	}
}

func (l GetProductByID) Execute(
	id string,
) (entity.Product, error) {
	product, err := l.productRepo.GetProductByID(context.Background(), id)
	if err != nil {
		return entity.Product{}, entity.NewErr(err)
	}

	if product.ID == "" {
		return entity.Product{}, entity.ErrProductNotFound
	}

	return product, nil
}
