package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProgramRoutes(version fiber.Router, db *gorm.DB, redis *redis.Client) {
	version.Get("programs/")
	version.Get("programs/:id")
	version.Post("programs/")
	version.Put("programs/:id")
	version.Put("programs/activate/:id")
	version.Put("programs/deactivate/:id")
	version.Delete("programs/:id")
}
