package router

import (
	"sapaUMKM-backend/config/db"
	"sapaUMKM-backend/config/redis"
	"sapaUMKM-backend/interface/http/middleware"
	"sapaUMKM-backend/interface/http/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {
	router := fiber.New(fiber.Config{
		// Prefork: true,
	})

	router.Use(middleware.CORS(), middleware.Logger())

	router.Get("test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World!"})
	})

	version := router.Group("/v1")

	routes.UserRoutes(version, db.DB, redis.RDB)
	routes.ProgramRoutes(version, db.DB, redis.RDB)

	return router
}
