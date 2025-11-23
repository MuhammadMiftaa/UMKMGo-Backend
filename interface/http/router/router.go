package router

import (
	"UMKMGo-backend/config/db"
	"UMKMGo-backend/config/log"
	"UMKMGo-backend/config/redis"
	"UMKMGo-backend/config/storage"
	"UMKMGo-backend/interface/http/middleware"
	"UMKMGo-backend/interface/http/routes"

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

	routes.UserRoutes(version, db.DB, redis.RDB, storage.MinioClient)
	routes.ProgramRoutes(version, db.DB, redis.RDB, storage.MinioClient)
	routes.ApplicationRoutes(version, db.DB, redis.RDB)
	routes.DashboardRoutes(version, db.DB)
	routes.SLARoutes(version, db.DB)
	routes.MobileRoutes(version, db.DB, storage.MinioClient)

	for _, routes := range router.Stack() {
		for _, r := range routes {
			log.Log.Infof("Registered route - METHOD: %s, PATH: %s", r.Method, r.Path)
		}
	}

	return router
}
