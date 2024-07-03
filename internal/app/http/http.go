package http

import (
	"github.com/danielmesquitta/products-api/internal/app/http/handler"
	"github.com/danielmesquitta/products-api/internal/app/http/middleware"
	"github.com/danielmesquitta/products-api/internal/app/http/router"
	"github.com/danielmesquitta/products-api/internal/config"
	"github.com/danielmesquitta/products-api/internal/domain/usecase"
	"github.com/danielmesquitta/products-api/internal/provider/repo"
	"github.com/danielmesquitta/products-api/internal/provider/repo/mysqlrepo"
	"github.com/danielmesquitta/products-api/pkg/logger"
	"github.com/danielmesquitta/products-api/pkg/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func Start() {
	depsProvider := fx.Provide(
		// Config
		config.LoadEnv,

		// PKGs
		validator.NewValidator,
		logger.NewLogger,

		// Providers
		mysqlrepo.NewMySQLDBConn,
		fx.Annotate(
			mysqlrepo.NewProductMySQLRepo,
			fx.As(new(repo.ProductRepo)),
		),

		// Use cases
		usecase.NewCreateProduct,
		usecase.NewDeleteProduct,
		usecase.NewGetProductByID,
		usecase.NewListProducts,
		usecase.NewUpdateProduct,

		// Handlers
		handler.NewProductHandler,

		// Middleware
		middleware.NewMiddleware,

		// Router
		router.NewRouter,

		// App
		NewApp,
	)

	container := fx.New(
		depsProvider,
		fx.Invoke(func(*echo.Echo) {}),
	)

	container.Run()
}
