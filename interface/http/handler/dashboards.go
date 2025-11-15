package handler

import (
	"net/http"

	"UMKMGo-backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type dashboardHandler struct {
	dashboardService service.DashboardService
}

func NewDashboardHandler(dashboardService service.DashboardService) *dashboardHandler {
	return &dashboardHandler{
		dashboardService: dashboardService,
	}
}

func (h *dashboardHandler) GetUMKMByCardType(c *fiber.Ctx) error {
	data, err := h.dashboardService.GetUMKMByCardType(c.Context())
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
		"message":    "UMKM statistics by card type",
		"data":       data,
	})
}

func (h *dashboardHandler) GetApplicationStatusSummary(c *fiber.Ctx) error {
	data, err := h.dashboardService.GetApplicationStatusSummary(c.Context())
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
		"message":    "Application status summary",
		"data":       data,
	})
}

func (h *dashboardHandler) GetApplicationStatusDetail(c *fiber.Ctx) error {
	data, err := h.dashboardService.GetApplicationStatusDetail(c.Context())
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
		"message":    "Application status detail",
		"data":       data,
	})
}

func (h *dashboardHandler) GetApplicationByType(c *fiber.Ctx) error {
	data, err := h.dashboardService.GetApplicationByType(c.Context())
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
		"message":    "Applications by type",
		"data":       data,
	})
}
