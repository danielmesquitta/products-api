package usecase

import (
	"context"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/repo"
	"github.com/danielmesquitta/products-api/pkg/validator"
	"github.com/jinzhu/copier"
)

type CreateProduct struct {
	validator   *validator.Validator
	productRepo repo.ProductRepo
}

func NewCreateProduct(
	validator *validator.Validator,
	productRepo repo.ProductRepo,
) *CreateProduct {
	return &CreateProduct{
		validator:   validator,
		productRepo: productRepo,
	}
}

type CreateProductParams struct {
	Name        string `json:"name"        validate:"min=3,max=255,required"`
	Description string `json:"description" validate:"min=3,required"`
	Price       int64  `json:"price"       validate:"min=1,required"`
}

func (l CreateProduct) Execute(
	params CreateProductParams,
) error {
	if err := l.validator.Validate(params); err != nil {
		validationErr := entity.ErrValidation
		validationErr.Message = err.Error()
		return validationErr
	}

	repoParams := repo.CreateProductParams{}
	if err := copier.Copy(&repoParams, params); err != nil {
		return entity.NewErr(err)
	}

	if err := l.productRepo.CreateProduct(context.Background(), repoParams); err != nil {
		return entity.NewErr(err)
	}

	return nil
}
