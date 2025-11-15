package routes

import (
	"UMKMGo-backend/config/redis"
	"UMKMGo-backend/interface/http/handler"
	"UMKMGo-backend/interface/http/middleware"
	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ApplicationRoutes(version fiber.Router, db *gorm.DB, redis redis.RedisRepository) {
	// Repository initialization
	applicationRepo := repository.NewApplicationsRepository(db)
	userRepo := repository.NewUsersRepository(db)

	// Service initialization
	applicationService := service.NewApplicationsService(applicationRepo, userRepo)

	// Handler initialization
	applicationHandler := handler.NewApplicationsHandler(applicationService)

	// Apply auth middleware
	version.Use(middleware.AuthMiddleware())

	applications := version.Group("/applications")
	{
		applications.Get("/", applicationHandler.GetAllApplications)
		applications.Get("/:id", applicationHandler.GetApplicationByID)
		// applications.Post("/", applicationHandler.CreateApplication)
		// applications.Put("/:id", applicationHandler.UpdateApplication)
		// applications.Delete("/:id", applicationHandler.DeleteApplication)

		// Screening decisions
		applications.Put("/screening-approve/:id", applicationHandler.ScreeningApprove)
		applications.Put("/screening-reject/:id", applicationHandler.ScreeningReject)
		applications.Put("/screening-revise/:id", applicationHandler.ScreeningRevise)

		// Final decisions
		applications.Put("/final-approve/:id", applicationHandler.FinalApprove)
		applications.Put("/final-reject/:id", applicationHandler.FinalReject)
	}
}
