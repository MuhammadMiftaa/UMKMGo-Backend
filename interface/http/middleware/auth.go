package middleware

import (
	"strings"

	"sapaUMKM-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwt := c.Get("Authorization")
		if jwt == "" {
			if c.Path() != "/" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"statusCode": 401,
					"status":     false,
					"error":      "Unauthorized",
				})
			}
			return c.Redirect("/login", fiber.StatusSeeOther)
		}

		tokenParts := strings.Split(jwt, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": 401,
				"status":     false,
				"error":      "Invalid Authorization format",
			})
		}

		token := tokenParts[1]
		userData, err := utils.VerifyToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": 401,
				"status":     false,
				"error":      "Unauthorized",
			})
		}

		c.Locals("user_data", userData)
		return c.Next()
	}
}
