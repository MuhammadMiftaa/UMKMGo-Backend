package routes

import (
	"sapaUMKM-backend/interface/http/handler"
	"sapaUMKM-backend/interface/http/middleware"
	"sapaUMKM-backend/internal/repository"
	"sapaUMKM-backend/internal/service"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ApplicationRoutes(version fiber.Router, db *gorm.DB, redis *redis.Client) {
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
