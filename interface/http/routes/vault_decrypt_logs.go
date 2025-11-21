package routes

import (
	"UMKMGo-backend/interface/http/handler"
	"UMKMGo-backend/interface/http/middleware"
	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func VaultDecryptLogRoutes(version fiber.Router, db *gorm.DB) {
	// Repository initialization
	vaultDecryptLogRepo := repository.NewVaultDecryptLogRepository(db)

	// Service initialization
	vaultDecryptLogService := service.NewVaultDecryptLogService(vaultDecryptLogRepo)

	// Handler initialization
	vaultDecryptLogHandler := handler.NewVaultDecryptLogHandler(vaultDecryptLogService)

	// Apply auth middleware for all vault decrypt log routes
	version.Use(middleware.AuthMiddleware(), middleware.ContextMiddleware())

	vaultDecrypt := version.Group("/vault-decrypt-logs")
	{
		vaultDecrypt.Get("/", vaultDecryptLogHandler.GetLogs)             // Get all logs with pagination
		vaultDecrypt.Get("/user", vaultDecryptLogHandler.GetLogsByUserID) // Get logs by user ID with pagination
		vaultDecrypt.Get("/umkm/:umkm_id", vaultDecryptLogHandler.GetLogsByUMKMID) // Get logs by UMKM ID with pagination
	}
}
