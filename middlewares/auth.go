package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte("secret")},
		TokenLookup: "header:Authorization,cookie:jwt,query:token",

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Redirect("/login", fiber.StatusSeeOther)
		},

		SuccessHandler: func(c *fiber.Ctx) error {
			// Extract token & claims and store in context
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			// Safely extract user_id
			userID, ok := claims["user_id"].(float64)
			if !ok {
				// return c.Status(fiber.StatusUnauthorized).SendString("Invalid token payload")
				return c.Redirect("/login", fiber.StatusSeeOther)
			}

			// Set useful values in context
			c.Locals("username", claims["username"])
			c.Locals("user_id", uint(userID)) // Set as uint for later use

			return c.Next()
		},
	})
}
