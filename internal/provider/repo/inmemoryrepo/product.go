package inmemoryrepo

import (
	"context"
	"time"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/repo"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type ProductInMemoryRepo struct {
	Products []entity.Product
}

func NewProductInMemoryRepo() *ProductInMemoryRepo {
	return &ProductInMemoryRepo{
		Products: []entity.Product{},
	}
}

func (p *ProductInMemoryRepo) GetProductByID(
	ctx context.Context,
	id string,
) (entity.Product, error) {
	for _, product := range p.Products {
		if product.ID == id {
			return product, nil
		}
	}

	return entity.Product{}, nil
}

func (p *ProductInMemoryRepo) ListProducts(
	ctx context.Context,
) ([]entity.Product, error) {
	return p.Products, nil
}

func (p *ProductInMemoryRepo) CreateProduct(
	ctx context.Context,
	params repo.CreateProductParams,
) error {
	product := entity.Product{}
	if err := copier.Copy(&product, params); err != nil {
		return entity.NewErr(err)
	}

	product.ID = uuid.NewString()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	p.Products = append(p.Products, product)

	return nil
}

func (p *ProductInMemoryRepo) UpdateProduct(
	ctx context.Context,
	params repo.UpdateProductParams,
) error {
	for i, product := range p.Products {
		if product.ID != params.ID {
			continue
		}

		if err := copier.CopyWithOption(
			&product,
			params,
			copier.Option{IgnoreEmpty: true},
		); err != nil {
			return entity.NewErr(err)
		}

		product.UpdatedAt = time.Now()

		p.Products[i] = product
		break
	}

	return nil
}

func (p *ProductInMemoryRepo) DeleteProduct(
	ctx context.Context,
	id string,
) error {
	for i, product := range p.Products {
		if product.ID == id {
			p.Products = append(p.Products[:i], p.Products[i+1:]...)
			break
		}
	}

	return nil
}
