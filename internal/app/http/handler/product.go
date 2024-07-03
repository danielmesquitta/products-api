package handler

import (
	"net/http"

	"github.com/danielmesquitta/products-api/internal/app/http/dto"
	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/danielmesquitta/products-api/internal/domain/usecase"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	createProductUseCase  *usecase.CreateProduct
	deleteProductUseCase  *usecase.DeleteProduct
	getProductByIDUseCase *usecase.GetProductByID
	listProductsUseCase   *usecase.ListProducts
	updateProductUseCase  *usecase.UpdateProduct
}

func NewProductHandler(
	createProductUseCase *usecase.CreateProduct,
	deleteProductUseCase *usecase.DeleteProduct,
	getProductByIDUseCase *usecase.GetProductByID,
	listProductsUseCase *usecase.ListProducts,
	updateProductUseCase *usecase.UpdateProduct,
) *ProductHandler {
	return &ProductHandler{
		createProductUseCase:  createProductUseCase,
		deleteProductUseCase:  deleteProductUseCase,
		getProductByIDUseCase: getProductByIDUseCase,
		listProductsUseCase:   listProductsUseCase,
		updateProductUseCase:  updateProductUseCase,
	}
}

// @Summary Create a new product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductRequestDTO true "Request body"
// @Success 201
// @Failure 400 {object} dto.ErrorResponseDTO
// @Failure 500 {object} dto.ErrorResponseDTO
// @Router /products [post]
func (p ProductHandler) CreateProduct(c echo.Context) error {
	requestData := dto.CreateProductRequestDTO{}
	if err := c.Bind(&requestData); err != nil {
		return entity.NewErr(err)
	}

	useCaseParams := usecase.CreateProductParams{}
	if err := copier.Copy(&useCaseParams, requestData); err != nil {
		return entity.NewErr(err)
	}

	if err := p.createProductUseCase.Execute(useCaseParams); err != nil {
		return entity.NewErr(err)
	}

	return c.NoContent(http.StatusCreated)
}

// @Summary Delete a product
// @Description Delete a product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200
// @Failure 400 {object} dto.ErrorResponseDTO
// @Failure 404 {object} dto.ErrorResponseDTO
// @Failure 500 {object} dto.ErrorResponseDTO
// @Router /products/{id} [delete]
func (p ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	if err := p.deleteProductUseCase.Execute(id); err != nil {
		return entity.NewErr(err)
	}

	return c.NoContent(http.StatusOK)
}

// @Summary Get a product by id
// @Description Get a product by id
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 400 {object} dto.ErrorResponseDTO
// @Failure 404 {object} dto.ErrorResponseDTO
// @Failure 500 {object} dto.ErrorResponseDTO
// @Router /products/{id} [get]
func (p ProductHandler) GetProductByID(c echo.Context) error {
	id := c.Param("id")

	product, err := p.getProductByIDUseCase.Execute(id)
	if err != nil {
		return entity.NewErr(err)
	}

	return c.JSON(http.StatusOK, product)
}

// @Summary List products
// @Description List products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} []entity.Product
// @Failure 400 {object} dto.ErrorResponseDTO
// @Failure 500 {object} dto.ErrorResponseDTO
// @Router /products [get]
func (p ProductHandler) ListProducts(c echo.Context) error {
	products, err := p.listProductsUseCase.Execute()
	if err != nil {
		return entity.NewErr(err)
	}

	return c.JSON(http.StatusOK, products)
}

// @Summary Update product
// @Description Update product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body dto.UpdateProductRequestDTO true "Request body"
// @Success 200
// @Failure 400 {object} dto.ErrorResponseDTO
// @Failure 404 {object} dto.ErrorResponseDTO
// @Failure 500 {object} dto.ErrorResponseDTO
// @Router /products/{id} [put]
func (p ProductHandler) UpdateProduct(c echo.Context) error {
	requestData := dto.UpdateProductRequestDTO{}
	if err := c.Bind(&requestData); err != nil {
		return entity.NewErr(err)
	}

	requestData.ID = c.Param("id")

	useCaseParams := usecase.UpdateProductParams{}
	if err := copier.Copy(&useCaseParams, requestData); err != nil {
		return entity.NewErr(err)
	}

	if err := p.updateProductUseCase.Execute(useCaseParams); err != nil {
		return entity.NewErr(err)
	}

	return c.NoContent(http.StatusOK)
}
