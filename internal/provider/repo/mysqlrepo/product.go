package mysqlrepo

import (
	"context"
	"database/sql"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/db/mysqldb"
	"github.com/danielmesquitta/products-api/internal/provider/repo"
	"github.com/jinzhu/copier"
)

type ProductMySQLRepo struct {
	db *mysqldb.Queries
}

func NewProductMySQLRepo(db *mysqldb.Queries) *ProductMySQLRepo {
	return &ProductMySQLRepo{
		db: db,
	}
}

func (p *ProductMySQLRepo) GetProductByID(
	ctx context.Context,
	id string,
) (entity.Product, error) {
	result, err := p.db.GetProductByID(ctx, id)

	if err == sql.ErrNoRows {
		return entity.Product{}, nil
	}

	if err != nil {
		return entity.Product{}, entity.NewErr(err)
	}

	product := entity.Product{}
	if err := copier.Copy(&product, result); err != nil {
		return entity.Product{}, entity.NewErr(err)
	}

	return product, nil
}

func (p *ProductMySQLRepo) ListProducts(
	ctx context.Context,
) ([]entity.Product, error) {
	results, err := p.db.ListProducts(ctx)
	if err != nil {
		return []entity.Product{}, entity.NewErr(err)
	}

	products := make([]entity.Product, 0, len(results))
	if err := copier.Copy(&products, results); err != nil {
		return []entity.Product{}, entity.NewErr(err)
	}

	return products, nil
}

func (p *ProductMySQLRepo) CreateProduct(
	ctx context.Context,
	product repo.CreateProductParams,
) error {
	params := mysqldb.CreateProductParams{}
	if err := copier.Copy(&params, product); err != nil {
		return entity.NewErr(err)
	}

	err := p.db.CreateProduct(ctx, params)
	if err != nil {
		return entity.NewErr(err)
	}

	return nil
}

func (p *ProductMySQLRepo) UpdateProduct(
	ctx context.Context,
	product repo.UpdateProductParams,
) error {
	params := mysqldb.UpdateProductParams{}
	if err := copier.Copy(&params, product); err != nil {
		return entity.NewErr(err)
	}

	err := p.db.UpdateProduct(ctx, params)
	if err != nil {
		return entity.NewErr(err)
	}

	return nil
}

func (p *ProductMySQLRepo) DeleteProduct(ctx context.Context, id string) error {
	err := p.db.DeleteProduct(ctx, id)
	if err != nil {
		return entity.NewErr(err)
	}

	return nil
}
