package middleware

import (
	"strings"

	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func authenticate(c *fiber.Ctx) (*dto.UserData, error) {
	jwt := c.Get("Authorization")
	if jwt == "" {
		if c.Path() != "/" {
			return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": 401,
				"status":     false,
				"error":      "Unauthorized",
			})
		}
		return nil, c.Redirect("/login", fiber.StatusSeeOther)
	}

	tokenParts := strings.Split(jwt, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"error":      "Invalid Authorization format",
		})
	}

	token := tokenParts[1]
	userData, err := utils.VerifyToken(token)
	if err != nil {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"error":      "Unauthorized",
		})
	}

	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		requestID = utils.GenerateRequestID()
		c.Set("X-Request-ID", requestID)
	}

	c.Locals("user_data", userData)
	c.Locals("userID", userData.ID)
	c.Locals("requestID", requestID)
	c.Locals("ipAddress", c.IP())
	c.Locals("userAgent", c.Get("User-Agent"))

	return &userData, nil
}

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userData, err := authenticate(c)
		if err != nil {
			return err
		}
		c.Locals("role", userData.Role)
		return c.Next()
	}
}

func MobileAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userData, err := authenticate(c)
		if err != nil {
			return err
		}
		c.Locals("phone", userData.Phone)
		c.Locals("kartuType", userData.KartuType)
		return c.Next()
	}
}
