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

func NewsRoutes(version fiber.Router, db *gorm.DB, minio *storage.MinIOManager) {
	// Repository initialization
	newsRepo := repository.NewNewsRepository(db)

	// Service initialization
	newsService := service.NewNewsService(newsRepo, minio)

	// Handler initialization
	newsHandler := handler.NewNewsHandler(newsService)

	// Web - Admin News Management (protected)
	news := version.Group("/news")
	news.Use(middleware.AuthMiddleware())
	{
		news.Get("/", newsHandler.GetAllNews)
		news.Get("/:id", newsHandler.GetNewsByID)
		news.Post("/", newsHandler.CreateNews)
		news.Put("/:id", newsHandler.UpdateNews)
		news.Delete("/:id", newsHandler.DeleteNews)
		news.Put("/publish/:id", newsHandler.PublishNews)
		news.Put("/unpublish/:id", newsHandler.UnpublishNews)
	}
}