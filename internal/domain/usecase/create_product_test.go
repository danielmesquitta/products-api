package usecase

import (
	"strings"
	"testing"

	"github.com/danielmesquitta/products-api/internal/provider/repo"
	"github.com/danielmesquitta/products-api/internal/provider/repo/inmemoryrepo"
	"github.com/danielmesquitta/products-api/pkg/validator"
)

func TestCreateProduct_Execute(t *testing.T) {
	type dependencies struct {
		validator   *validator.Validator
		productRepo repo.ProductRepo
	}
	type args struct {
		params CreateProductParams
	}
	tests := []struct {
		name         string
		dependencies dependencies
		args         args
		wantErr      bool
	}{
		{
			name: "should create a product",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: CreateProductParams{
					Name:        "Valid product name",
					Description: "Valid product description",
					Price:       1000,
				},
			},
			wantErr: false,
		},
		{
			name: "should return an error when name is empty",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: CreateProductParams{
					Description: "Valid product description",
					Price:       1000,
				},
			},
			wantErr: true,
		},
		{
			name: "should return an error when name is less than 3 characters",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: CreateProductParams{
					Name:        "In",
					Description: "Valid product description",
					Price:       1000,
				},
			},
			wantErr: true,
		},
		{
			name: "should return an error when name is grater than 255 characters",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: CreateProductParams{
					Name:        strings.Repeat("a", 256),
					Description: "Valid product description",
					Price:       1000,
				},
			},
			wantErr: true,
		},
		{
			name: "should return an error when description is empty",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: CreateProductParams{
					Name:  "Valid product name",
					Price: 1000,
				},
			},
			wantErr: true,
		},
		{
			name: "should return an error when description is less than 3 characters",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: CreateProductParams{
					Name:        "Valid product name",
					Description: "In",
					Price:       1000,
				},
			},
			wantErr: true,
		},
		{
			name: "should not return an error with large descriptions",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: CreateProductParams{
					Name:        "Valid product name",
					Description: strings.Repeat("a", 1000),
					Price:       1000,
				},
			},
			wantErr: false,
		},
		{
			name: "should return an error when price is empty",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: CreateProductParams{
					Name:        "Valid product name",
					Description: "Valid product description",
				},
			},
			wantErr: true,
		},
		{
			name: "should return an error when price is lower than 1",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: CreateProductParams{
					Name:        "Valid product name",
					Description: "Valid product description",
					Price:       0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := CreateProduct{
				validator:   tt.dependencies.validator,
				productRepo: tt.dependencies.productRepo,
			}
			if err := l.Execute(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf(
					"CreateProduct.Execute() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
