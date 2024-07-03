package usecase

import (
	"testing"
	"time"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/repo/inmemoryrepo"
	"github.com/google/uuid"
)

func TestDeleteProduct_Execute(t *testing.T) {
	existingProduct := entity.Product{
		ID:          uuid.NewString(),
		Name:        "Product name",
		Description: "Product description",
		Price:       1000,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	type dependencies struct {
		productRepo *inmemoryrepo.ProductInMemoryRepo
	}
	type args struct {
		id string
	}
	tests := []struct {
		name         string
		dependencies dependencies
		args         args
		wantErr      bool
	}{
		{
			name: "should delete a product",
			dependencies: dependencies{
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				id: existingProduct.ID,
			},
			wantErr: false,
		},
		{
			name: "should return an error when product not found",
			dependencies: dependencies{
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				id: uuid.NewString(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dependencies.productRepo.Products = []entity.Product{
				existingProduct,
			}

			l := DeleteProduct{
				productRepo: tt.dependencies.productRepo,
			}
			if err := l.Execute(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf(
					"DeleteProduct.Execute() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
