package usecase

import (
	"context"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/repo"
	"github.com/danielmesquitta/products-api/pkg/validator"
	"github.com/jinzhu/copier"
)

type UpdateProduct struct {
	validator   *validator.Validator
	productRepo repo.ProductRepo
}

func NewUpdateProduct(
	validator *validator.Validator,
	productRepo repo.ProductRepo,
) *UpdateProduct {
	return &UpdateProduct{
		validator:   validator,
		productRepo: productRepo,
	}
}

type UpdateProductParams struct {
	ID          string `json:"id"          validate:"uuid,required"`
	Name        string `json:"name"        validate:"omitempty,min=3,max=255"`
	Description string `json:"description" validate:"omitempty,min=3"`
	Price       int64  `json:"price"       validate:"omitempty,min=1"`
}

func (l UpdateProduct) Execute(
	params UpdateProductParams,
) error {
	if err := l.validator.Validate(params); err != nil {
		validationErr := entity.ErrValidation
		validationErr.Message = err.Error()
		return validationErr
	}

	product, err := l.productRepo.GetProductByID(
		context.Background(),
		params.ID,
	)
	if err != nil {
		return entity.NewErr(err)
	}

	if product.ID == "" {
		return entity.ErrProductNotFound
	}

	repoParams := repo.UpdateProductParams{}
	if err := copier.Copy(&repoParams, product); err != nil {
		return entity.NewErr(err)
	}

	if err := copier.CopyWithOption(&repoParams, params, copier.Option{
		IgnoreEmpty: true,
	}); err != nil {
		return entity.NewErr(err)
	}

	if err := l.productRepo.UpdateProduct(context.Background(), repoParams); err != nil {
		return entity.NewErr(err)
	}

	return nil
}
