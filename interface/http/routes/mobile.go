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

	// Service initialization
	mobileService := service.NewMobileService(mobileRepo, programRepo, minio)

	// Handler initialization
	mobileHandler := handler.NewMobileHandler(mobileService)

	// Apply auth middleware for all mobile routes
	version.Use(middleware.AuthMiddleware(), middleware.ContextMiddleware())

	mobile := version.Group("/mobile")
	{
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
			documents.Post("/nib", mobileHandler.UploadNIB)
			documents.Post("/npwp", mobileHandler.UploadNPWP)
			documents.Post("/revenue-record", mobileHandler.UploadRevenueRecord)
			documents.Post("/business-permit", mobileHandler.UploadBusinessPermit)
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
	}
}