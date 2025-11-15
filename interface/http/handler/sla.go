package handler

import (
	"net/http"

	"UMKMGo-backend/internal/service"
	"UMKMGo-backend/internal/types/dto"

	"github.com/gofiber/fiber/v2"
)

type slaHandler struct {
	slaService service.SLAService
}

func NewSLAHandler(slaService service.SLAService) *slaHandler {
	return &slaHandler{
		slaService: slaService,
	}
}

func (h *slaHandler) GetSLAScreening(c *fiber.Ctx) error {
	result, err := h.slaService.GetSLAScreening(c.Context())
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
		"message":    "SLA screening retrieved successfully",
		"data":       result,
	})
}

func (h *slaHandler) GetSLAFinal(c *fiber.Ctx) error {
	result, err := h.slaService.GetSLAFinal(c.Context())
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
		"message":    "SLA final retrieved successfully",
		"data":       result,
	})
}

func (h *slaHandler) UpdateSLAScreening(c *fiber.Ctx) error {
	var slaRequest dto.SLA
	err := c.BodyParser(&slaRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	result, err := h.slaService.UpdateSLAScreening(c.Context(), slaRequest)
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
		"message":    "SLA screening updated successfully",
		"data":       result,
	})
}

func (h *slaHandler) UpdateSLAFinal(c *fiber.Ctx) error {
	var slaRequest dto.SLA
	err := c.BodyParser(&slaRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	result, err := h.slaService.UpdateSLAFinal(c.Context(), slaRequest)
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
		"message":    "SLA final updated successfully",
		"data":       result,
	})
}

func (h *slaHandler) ExportApplications(c *fiber.Ctx) error {
	var exportRequest dto.ExportRequest
	err := c.BodyParser(&exportRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	fileData, filename, err := h.slaService.ExportApplications(c.Context(), exportRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	// Set content type based on file type
	contentType := "text/plain"
	if exportRequest.FileType == "excel" {
		contentType = "text/csv"
	}

	c.Set("Content-Type", contentType)
	c.Set("Content-Disposition", "attachment; filename="+filename)

	return c.Send(fileData)
}

func (h *slaHandler) ExportPrograms(c *fiber.Ctx) error {
	var exportRequest dto.ExportRequest
	err := c.BodyParser(&exportRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	fileData, filename, err := h.slaService.ExportPrograms(c.Context(), exportRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	// Set content type based on file type
	contentType := "text/plain"
	if exportRequest.FileType == "excel" {
		contentType = "text/csv"
	}

	c.Set("Content-Type", contentType)
	c.Set("Content-Disposition", "attachment; filename="+filename)

	return c.Send(fileData)
}
