package jwt

import (
	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userID int64) (token string, err error)
	ValidateToken(tokenString string) (token *jwt.Token, err error)
}
