package router

import (
	_ "github.com/danielmesquitta/products-api/docs"
	"github.com/danielmesquitta/products-api/internal/app/http/handler"
	"github.com/danielmesquitta/products-api/internal/config"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

type Router struct {
	env            *config.Env
	healthHandler  *handler.HealthHandler
	productHandler *handler.ProductHandler
}

func NewRouter(
	env *config.Env,
	healthHandler *handler.HealthHandler,
	productHandler *handler.ProductHandler,
) *Router {
	return &Router{
		env:            env,
		healthHandler:  healthHandler,
		productHandler: productHandler,
	}
}

func (r *Router) Register(
	app *echo.Echo,
) {
	basePath := "/api/v1"
	apiV1 := app.Group(basePath)

	apiV1.GET("/health", r.healthHandler.Health)

	apiV1.POST("/products", r.productHandler.CreateProduct)
	apiV1.DELETE("/products/:id", r.productHandler.DeleteProduct)
	apiV1.GET("/products/:id", r.productHandler.GetProductByID)
	apiV1.GET("/products", r.productHandler.ListProducts)
	apiV1.PUT("/products/:id", r.productHandler.UpdateProduct)

	apiV1.GET("/docs/*", echoSwagger.WrapHandler)
}
