package middleware

import (
	"github.com/b0gochort/conv_service/internal/utils/jwt"
	"go.uber.org/zap"
)

type Middleware struct {
	jwtSvc jwt.JWTService
	logger zap.Logger
}

func NewMiddleware(jwtSvc jwt.JWTService) *Middleware {
	return &Middleware{
		jwtSvc: jwtSvc,
	}
}
