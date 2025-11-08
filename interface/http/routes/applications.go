package routes

import (
	"sapaUMKM-backend/interface/http/middleware"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ApplicationRoutes(version fiber.Router, db *gorm.DB, redis *redis.Client) {
	// Apply auth middleware
	version.Use(middleware.AuthMiddleware())

	trainingApps := version.Group("/applications/training")
	{
		trainingApps.Get("/", )
		trainingApps.Get("/:id", )
		trainingApps.Put("/screening-approve/:id", )
		trainingApps.Put("/screening-reject/:id", )
		trainingApps.Put("/screening-revise/:id", )
		trainingApps.Put("/final-approve/:id", )
		trainingApps.Put("/final-reject/:id", )
	}
}
