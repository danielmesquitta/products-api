package http

import (
	"context"

	"github.com/danielmesquitta/products-api/internal/app/http/middleware"
	"github.com/danielmesquitta/products-api/internal/app/http/router"
	"github.com/danielmesquitta/products-api/internal/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewApp(
	lc fx.Lifecycle,
	e *config.Env,
	m *middleware.Middleware,
	r *router.Router,
) *echo.Echo {
	app := echo.New()

	defaultErrorHandler := app.HTTPErrorHandler
	customErrorHandler := m.ErrorHandler(defaultErrorHandler)
	app.HTTPErrorHandler = customErrorHandler

	r.Register(app)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := app.Start(":" + e.Port); err != nil {
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown(context.Background())
		},
	})

	return app
}
