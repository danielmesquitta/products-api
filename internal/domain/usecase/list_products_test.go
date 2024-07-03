package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/repo/inmemoryrepo"
	"github.com/google/uuid"
)

func TestListProducts_Execute(t *testing.T) {
	product1 := entity.Product{
		ID:          uuid.NewString(),
		Name:        "Product 1",
		Description: "Product description 1",
		Price:       1000,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	product2 := entity.Product{
		ID:          uuid.NewString(),
		Name:        "Product 2",
		Description: "Product description 2",
		Price:       500,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	products := []entity.Product{
		product1,
		product2,
	}

	type dependencies struct {
		productRepo *inmemoryrepo.ProductInMemoryRepo
	}
	tests := []struct {
		name         string
		dependencies dependencies
		want         []entity.Product
		wantErr      bool
	}{
		{
			name: "should list products",
			dependencies: dependencies{
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			want:    products,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dependencies.productRepo.Products = products
			l := ListProducts{
				productRepo: tt.dependencies.productRepo,
			}
			got, err := l.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"ListProducts.Execute() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListProducts.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
