package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/provider/repo/inmemoryrepo"
	"github.com/google/uuid"
)

func TestGetProductByID_Execute(t *testing.T) {
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
		want         entity.Product
		wantErr      bool
	}{
		{
			name: "should get product by id",
			dependencies: dependencies{
				productRepo: inmemoryrepo.NewProductInMemoryRepo(),
			},
			args: args{
				id: existingProduct.ID,
			},
			want:    existingProduct,
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
			want:    entity.Product{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dependencies.productRepo.Products = []entity.Product{
				existingProduct,
			}
			l := GetProductByID{
				productRepo: tt.dependencies.productRepo,
			}
			got, err := l.Execute(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"GetProductByID.Execute() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByID.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
