package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) JWTAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authorizationHeader := c.Request().Header.Get("Authorization")
			bearerToken := strings.Split(authorizationHeader, " ")

			if len(bearerToken) != 2 {
				return c.JSON(http.StatusUnauthorized, "invalid authorization token")
			}

			tokenStr := bearerToken[1]
			token, err := m.jwtSvc.ValidateToken(tokenStr)
			if err != nil {
				return c.JSON(
					http.StatusUnauthorized,
					"invalid authorization token",
				)
			}

			if !token.Valid {
				return c.JSON(
					http.StatusUnauthorized,
					"invalid authorization token",
				)
			}

			claims := token.Claims.(jwt.MapClaims)
			c.Set("user_id", int64(claims["user_id"].(float64)))
			return next(c)
		}
	}
}
