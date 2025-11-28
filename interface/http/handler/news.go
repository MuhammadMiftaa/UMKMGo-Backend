package handler

import (
	"net/http"
	"strconv"

	"UMKMGo-backend/internal/service"
	"UMKMGo-backend/internal/types/dto"

	"github.com/gofiber/fiber/v2"
)

type NewsHandler struct {
	newsService service.NewsService
}

func NewNewsHandler(newsService service.NewsService) *NewsHandler {
	return &NewsHandler{
		newsService: newsService,
	}
}

func (h *NewsHandler) GetAllNews(c *fiber.Ctx) error {
	params := dto.NewsQueryParams{
		Page:     c.QueryInt("page", 1),
		Limit:    c.QueryInt("limit", 10),
		Category: c.Query("category"),
		Search:   c.Query("search"),
		Tag:      c.Query("tag"),
	}

	// Parse is_published parameter
	if isPublishedStr := c.Query("is_published"); isPublishedStr != "" {
		isPublished := isPublishedStr == "true"
		params.IsPublished = &isPublished
	}

	news, total, err := h.newsService.GetAllNews(c.Context(), params)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Get all news",
		"data": fiber.Map{
			"news":  news,
			"total": total,
			"page":  params.Page,
			"limit": params.Limit,
		},
	})
}

func (h *NewsHandler) GetNewsByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid news ID",
		})
	}

	news, err := h.newsService.GetNewsByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"statusCode": 404,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Get news detail",
		"data":       news,
	})
}

func (h *NewsHandler) CreateNews(c *fiber.Ctx) error {
	var request dto.NewsRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	// Get user data from context
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	news, err := h.newsService.CreateNews(c.Context(), int(userData.ID), request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"statusCode": 201,
		"status":     true,
		"message":    "News created successfully",
		"data":       news,
	})
}

func (h *NewsHandler) UpdateNews(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid news ID",
		})
	}

	var request dto.NewsRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	news, err := h.newsService.UpdateNews(c.Context(), id, request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "News updated successfully",
		"data":       news,
	})
}

func (h *NewsHandler) DeleteNews(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid news ID",
		})
	}

	if err := h.newsService.DeleteNews(c.Context(), id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "News deleted successfully",
	})
}

func (h *NewsHandler) PublishNews(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid news ID",
		})
	}

	news, err := h.newsService.PublishNews(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "News published successfully",
		"data":       news,
	})
}

func (h *NewsHandler) UnpublishNews(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid news ID",
		})
	}

	news, err := h.newsService.UnpublishNews(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "News unpublished successfully",
		"data":       news,
	})
}