package routes

import (
	"UMKMGo-backend/interface/http/handler"
	"UMKMGo-backend/interface/http/middleware"
	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SLARoutes(version fiber.Router, db *gorm.DB) {
	slaRepo := repository.NewSLARepository(db)

	slaService := service.NewSLAService(slaRepo)

	slaHandler := handler.NewSLAHandler(slaService)

	version.Use(middleware.AuthMiddleware())

	sla := version.Group("/sla")
	{
		sla.Get("/screening", slaHandler.GetSLAScreening)
		sla.Get("/final", slaHandler.GetSLAFinal)
		sla.Put("/screening", slaHandler.UpdateSLAScreening)
		sla.Put("/final", slaHandler.UpdateSLAFinal)
		sla.Post("/export-applications", slaHandler.ExportApplications)
		sla.Post("/export-programs", slaHandler.ExportPrograms)
	}
}
