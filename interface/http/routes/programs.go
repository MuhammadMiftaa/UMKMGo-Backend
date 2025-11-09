package routes

import (
	"sapaUMKM-backend/config/redis"
	"sapaUMKM-backend/config/storage"
	"sapaUMKM-backend/interface/http/handler"
	"sapaUMKM-backend/interface/http/middleware"
	"sapaUMKM-backend/internal/repository"
	"sapaUMKM-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProgramRoutes(version fiber.Router, db *gorm.DB, redis redis.RedisRepository, minio *storage.MinIOManager) {
	// Repository initialization
	programRepo := repository.NewProgramsRepository(db)
	userRepo := repository.NewUsersRepository(db)

	// Service initialization
	programService := service.NewProgramsService(programRepo, userRepo, redis, minio)

	// Handler initialization
	programHandler := handler.NewProgramsHandler(programService)

	// Apply auth middleware
	version.Use(middleware.AuthMiddleware())

	programs := version.Group("/programs")
	{
		programs.Get("/", programHandler.GetAllPrograms)
		programs.Get("/:id", programHandler.GetProgramByID)
		programs.Post("/", programHandler.CreateProgram)
		programs.Put("/:id", programHandler.UpdateProgram)
		programs.Put("/activate/:id", programHandler.ActivateProgram)
		programs.Put("/deactivate/:id", programHandler.DeactivateProgram)
		programs.Delete("/:id", programHandler.DeleteProgram)
	}
}
