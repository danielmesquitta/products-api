package middleware

import (
	"net/http"

	"github.com/danielmesquitta/products-api/internal/app/http/dto"
	"github.com/danielmesquitta/products-api/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

var mapErrTypeToStatusCode = map[entity.ErrType]int{
	entity.ErrTypeValidation: http.StatusBadRequest,
	entity.ErrTypeUnknown:    http.StatusInternalServerError,
	entity.ErrTypeNotFound:   http.StatusNotFound,
}

func (m *Middleware) ErrorHandler(
	defaultErrorHandler echo.HTTPErrorHandler,
) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		if appErr, ok := err.(*entity.Err); ok {
			statusCode := mapErrTypeToStatusCode[appErr.Type]
			if internalServerError := statusCode >= 500 || statusCode == 0; internalServerError {
				req := c.Request()

				requestData := map[string]any{}
				_ = c.Bind(&requestData)

				m.log.Errorln(
					appErr.Error(),
					"url",
					req.URL.Path,
					"body",
					requestData,
					"query",
					c.QueryParams(),
					"params",
					c.ParamValues(),
					appErr.StackTrace,
				)

				err := c.JSON(
					statusCode,
					dto.ErrorResponseDTO{Message: "internal server error"},
				)
				if err != nil {
					m.log.Errorln(err)
				}
				return
			}

			err := c.JSON(
				statusCode,
				dto.ErrorResponseDTO{Message: appErr.Error()},
			)
			if err != nil {
				m.log.Errorln(err)
			}
			return
		}

		defaultErrorHandler(err, c)
	}
}
