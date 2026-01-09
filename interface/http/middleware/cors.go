package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173,https://umkmgo.miftech.web.id,https://umkmgo-staging.miftech.web.id,https://pwa-umkmgo.miftech.web.id,https://pwa-umkmgo-staging.miftech.web.id",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type,Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 hours
	})
}
