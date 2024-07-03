package repo

import (
	"context"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
)

type CreateProductParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

type UpdateProductParams struct {
	ID          string `json:"id"`
	Name        string `json:"name"        validate:"min=3,max=255"`
	Description string `json:"description" validate:"min=3"`
	Price       int64  `json:"price"       validate:"min=1"`
}

type ProductRepo interface {
	GetProductByID(ctx context.Context, id string) (entity.Product, error)
	ListProducts(ctx context.Context) ([]entity.Product, error)
	CreateProduct(ctx context.Context, product CreateProductParams) error
	UpdateProduct(ctx context.Context, product UpdateProductParams) error
	DeleteProduct(ctx context.Context, id string) error
}
