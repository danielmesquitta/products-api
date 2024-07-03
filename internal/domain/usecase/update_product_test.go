package usecase

import (
	"strings"
	"testing"
	"time"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/repo/inmemoryrepo"
	"github.com/danielmesquitta/products-api/pkg/validator"
	"github.com/google/uuid"
)

func TestUpdateProduct_Execute(t *testing.T) {
	existingProduct := entity.Product{
		ID:          uuid.NewString(),
		Name:        "Product name",
		Description: "Product description",
		Price:       1000,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	type dependencies struct {
		validator   *validator.Validator
		productRepo *inmemoryrepo.ProductInMemoryRepo
	}
	type args struct {
		params UpdateProductParams
	}
	tests := []struct {
		name         string
		dependencies dependencies
		args         args
		wantErr      bool
	}{
		{
			name: "should update product with all fields filled in",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: UpdateProductParams{
					ID:          existingProduct.ID,
					Name:        "Updated product name",
					Description: "Updated product description",
					Price:       500,
				},
			},
			wantErr: false,
		},
		{
			name: "should update product with only name filled in",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: UpdateProductParams{
					ID:   existingProduct.ID,
					Name: "Updated product name",
				},
			},
			wantErr: false,
		},
		{
			name: "should update product with only description filled in",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: UpdateProductParams{
					ID:          existingProduct.ID,
					Description: "Updated product description",
				},
			},
			wantErr: false,
		},
		{
			name: "should update product with only price filled in",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: UpdateProductParams{
					ID:    existingProduct.ID,
					Price: 500,
				},
			},
			wantErr: false,
		},
		{
			name: "should return an error when not found",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: UpdateProductParams{
					ID:          uuid.NewString(),
					Name:        "Updated product name",
					Description: "Updated product description",
					Price:       500,
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
				params: UpdateProductParams{
					ID:          existingProduct.ID,
					Name:        "In",
					Description: "Updated product description",
					Price:       500,
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
				params: UpdateProductParams{
					ID:          existingProduct.ID,
					Name:        strings.Repeat("a", 256),
					Description: "Updated product description",
					Price:       500,
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
				params: UpdateProductParams{
					ID:          existingProduct.ID,
					Name:        "Updated product name",
					Description: "In",
					Price:       500,
				},
			},
			wantErr: true,
		},
		{
			name: "should update product with large descriptions",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: UpdateProductParams{
					ID:          existingProduct.ID,
					Name:        "Updated product name",
					Description: strings.Repeat("a", 1000),
					Price:       500,
				},
			},
			wantErr: false,
		},
		{
			name: "should return an error when price is lower than 1",
			dependencies: dependencies{
				validator:   validator.NewValidator(),
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				params: UpdateProductParams{
					ID:          existingProduct.ID,
					Name:        "Updated product name",
					Description: "Updated product description",
					Price:       -1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dependencies.productRepo.Products = []entity.Product{
				existingProduct,
			}
			l := UpdateProduct{
				validator:   tt.dependencies.validator,
				productRepo: tt.dependencies.productRepo,
			}
			if err := l.Execute(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf(
					"UpdateProduct.Execute() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
