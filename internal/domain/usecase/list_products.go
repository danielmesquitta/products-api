package usecase

import (
	"context"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/repo"
)

type ListProducts struct {
	productRepo repo.ProductRepo
}

func NewListProducts(productRepo repo.ProductRepo) *ListProducts {
	return &ListProducts{
		productRepo: productRepo,
	}
}

func (l ListProducts) Execute() ([]entity.Product, error) {
	products, err := l.productRepo.ListProducts(context.Background())
	if err != nil {
		return []entity.Product{}, entity.NewErr(err)
	}

	return products, nil
}
