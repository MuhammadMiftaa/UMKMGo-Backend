package routes

import (
	"UMKMGo-backend/config/storage"
	"UMKMGo-backend/interface/http/handler"
	"UMKMGo-backend/interface/http/middleware"
	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MobileRoutes(version fiber.Router, db *gorm.DB, minio *storage.MinIOManager) {
	// Repository initialization
	mobileRepo := repository.NewMobileRepository(db)
	programRepo := repository.NewProgramsRepository(db)
	notificationRepo := repository.NewNotificationRepository(db)
	applicationRepo := repository.NewApplicationsRepository(db)
	slaRepo := repository.NewSLARepository(db)
	vaultDecryptLogRepo := repository.NewVaultDecryptLogRepository(db)

	// Service initialization
	mobileService := service.NewMobileService(mobileRepo, programRepo, notificationRepo, vaultDecryptLogRepo, applicationRepo, slaRepo, minio)

	// Handler initialization
	mobileHandler := handler.NewMobileHandler(mobileService)

	// Apply auth middleware for all mobile routes
	version.Use(middleware.MobileAuthMiddleware())

	mobile := version.Group("/mobile")
	{
		mobile.Get("/dashboard", mobileHandler.GetDashboard)

		// Programs
		programs := mobile.Group("/programs")
		{
			programs.Get("/training", mobileHandler.GetTrainingPrograms)
			programs.Get("/certification", mobileHandler.GetCertificationPrograms)
			programs.Get("/funding", mobileHandler.GetFundingPrograms)
			programs.Get("/:id", mobileHandler.GetProgramDetail)
		}

		// UMKM Profile
		profile := mobile.Group("/profile")
		{
			profile.Get("/", mobileHandler.GetUMKMProfile)
			profile.Put("/", mobileHandler.UpdateUMKMProfile)
		}

		// Documents
		documents := mobile.Group("/documents")
		{
			documents.Get("/", mobileHandler.GetUMKMDocuments)
			documents.Post("/upload", mobileHandler.UploadDocument)
		}

		// Applications
		applications := mobile.Group("/applications")
		{
			applications.Post("/training", mobileHandler.CreateTrainingApplication)
			applications.Post("/certification", mobileHandler.CreateCertificationApplication)
			applications.Post("/funding", mobileHandler.CreateFundingApplication)
			applications.Get("/", mobileHandler.GetApplicationList)
			applications.Get("/:id", mobileHandler.GetApplicationDetail)
		}

		// Notifications
		notifications := mobile.Group("/notifications")
		{
			notifications.Get("/", mobileHandler.GetNotificationsByUMKMID)
			notifications.Get("/unread-count", mobileHandler.GetUnreadCount)
			notifications.Put("/mark-as-read/:id", mobileHandler.MarkNotificationsAsRead)
			notifications.Put("/mark-all-as-read", mobileHandler.MarkAllNotificationsAsRead)
		}

		// News
		news := mobile.Group("/news")
		{
			news.Get("/", mobileHandler.GetPublishedNews)
			news.Get("/:slug", mobileHandler.GetNewsDetailBySlug)
		}
	}
}
