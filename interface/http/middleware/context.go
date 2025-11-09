package middleware

import (
	"context"
	"sapaUMKM-backend/internal/utils/constant"

	"github.com/gofiber/fiber/v2"
)

func ContextMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.UserContext(), constant.DefaultConnectionTimeout)
		defer cancel()
		
		ctx = context.WithValue(ctx, "requestID", c.Locals("requestID"))
		ctx = context.WithValue(ctx, "userID", c.Locals("userID"))
		ctx = context.WithValue(ctx, "role", c.Locals("role"))
		ctx = context.WithValue(ctx, "ipAddress", c.IP())
		ctx = context.WithValue(ctx, "userAgent", c.Get("User-Agent"))

		c.SetUserContext(ctx)

		err := c.Next()

		if ctx.Err() == context.DeadlineExceeded {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "request timeout",
			})
		}

		return err
	}
}
