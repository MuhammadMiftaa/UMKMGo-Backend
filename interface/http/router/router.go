package router

import (
	"sort"

	"sapaUMKM-backend/config/db"
	"sapaUMKM-backend/config/log"
	"sapaUMKM-backend/config/redis"
	"sapaUMKM-backend/config/storage"
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
	routes.ProgramRoutes(version, db.DB, redis.RDB, storage.MinioClient)
	routes.ApplicationRoutes(version, db.DB, redis.RDB)

	for _, routes := range router.Stack() {
		for _, r := range routes {
			fields := map[string]interface{}{
				"METHOD": r.Method,
				"PATH":   r.Path,
			}

			keys := make([]string, 0, len(fields))
			for k := range fields {
				keys = append(keys, k)
			}

			sort.Strings(keys)

			log.Info("Registered route - ", fields)
		}
	}

	return router
}
