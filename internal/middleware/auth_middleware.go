package middleware

import (
	"net/http"
	"strings"

	"github.com/daytrip-idn-api/pkg/utils"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing Authorization header"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
		}

		tokenString := parts[1]
		token, err := utils.VerifyToken(tokenString, false)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or expired token"})
		}

		claims, ok := utils.ExtractClaims(token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Failed to extract claims from token"})
		}

		userId, ok := claims["id"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token claims"})
		}

		c.Set("userId", int(userId))

		return next(c)
	}
}
