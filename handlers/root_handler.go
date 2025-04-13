package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func HandleRoot(c *fiber.Ctx) error {
	var tokenString string

	// Try to get token from cookie
	tokenString = c.Cookies("jwt")

	// Or from Authorization header
	if tokenString == "" {
		authHeader := c.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}
	}

	// Or from query param
	if tokenString == "" {
		tokenString = c.Query("token")
	}

	// No token found → redirect to login
	if tokenString == "" {
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	// Parse and verify token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	// Valid token → redirect to todos
	return c.Redirect("/todos", fiber.StatusSeeOther)
}
