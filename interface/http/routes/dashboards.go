package routes

import (
	"sapaUMKM-backend/interface/http/handler"
	"sapaUMKM-backend/interface/http/middleware"
	"sapaUMKM-backend/internal/repository"
	"sapaUMKM-backend/internal/service"
	
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DashboardRoutes(version fiber.Router, db *gorm.DB) {
	dashboardRepo := repository.NewDashboardRepository(db)
	
	dashboardService := service.NewDashboardService(dashboardRepo)
	
	dashboardHandler := handler.NewDashboardHandler(dashboardService)
	
	version.Use(middleware.AuthMiddleware())
	
	dashboard := version.Group("/dashboard")
	{
		dashboard.Get("/umkm-by-card-type", dashboardHandler.GetUMKMByCardType)
		dashboard.Get("/application-status-summary", dashboardHandler.GetApplicationStatusSummary)
		dashboard.Get("/application-status-detail", dashboardHandler.GetApplicationStatusDetail)
		dashboard.Get("/application-by-type", dashboardHandler.GetApplicationByType)
	}
}