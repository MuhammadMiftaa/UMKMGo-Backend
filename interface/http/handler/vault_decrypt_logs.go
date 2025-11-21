package handler

import (
	"fmt"
	"strconv"

	"UMKMGo-backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type VaultDecryptLogHandler struct {
	vaultDecryptLogService service.VaultDecryptLogService
}

func NewVaultDecryptLogHandler(vaultDecryptLogService service.VaultDecryptLogService) *VaultDecryptLogHandler {
	return &VaultDecryptLogHandler{
		vaultDecryptLogService: vaultDecryptLogService,
	}
}

// GetLogs retrieves all vault decrypt logs with pagination.
func (h *VaultDecryptLogHandler) GetLogs(c *fiber.Ctx) error {
	logs, err := h.vaultDecryptLogService.GetLogs(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"status":     false,
			"error":      "Failed to retrieve logs",
		})
	}
	return c.JSON(logs)
}

// GetLogsByUserID retrieves vault decrypt logs for a specific user with pagination.
func (h *VaultDecryptLogHandler) GetLogsByUserID(c *fiber.Ctx) error {
	userIDVal := c.Locals("userID")
	userID, ok := userIDVal.(int)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"status":     false,
			"error":      "Invalid or missing userID",
		})
	}
	logs, err := h.vaultDecryptLogService.GetLogsByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"status":     false,
			"error":      "Failed to retrieve logs",
		})
	}
	return c.JSON(logs)
}

// GetLogsByUMKMID retrieves vault decrypt logs for a specific UMKM with pagination.
func (h *VaultDecryptLogHandler) GetLogsByUMKMID(c *fiber.Ctx) error {
	umkmID := c.Params("umkm_id")
	if umkmID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"status":     false,
			"error":      "UMKM ID is required",
		})
	}

	// Convert UMKM ID to integer
	umkmIDInt, err := strconv.Atoi(umkmID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"status":     false,
			"error":      fmt.Sprintf("Invalid UMKM ID: %s", umkmID),
		})
	}

	logs, err := h.vaultDecryptLogService.GetLogsByUMKMID(c.Context(), umkmIDInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"status":     false,
			"error":      "Failed to retrieve logs",
		})
	}
	return c.JSON(logs)
}
