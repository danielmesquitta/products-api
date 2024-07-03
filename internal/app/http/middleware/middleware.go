package middleware

import "github.com/danielmesquitta/products-api/pkg/logger"

type Middleware struct {
	log *logger.Logger
}

func NewMiddleware(log *logger.Logger) *Middleware {
	return &Middleware{log: log}
}
