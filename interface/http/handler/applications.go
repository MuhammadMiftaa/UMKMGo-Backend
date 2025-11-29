package handler

import (
	"net/http"
	"strconv"

	"UMKMGo-backend/internal/service"
	"UMKMGo-backend/internal/types/dto"

	"github.com/gofiber/fiber/v2"
)

type applicationsHandler struct {
	applicationsService service.ApplicationsService
}

func NewApplicationsHandler(applicationsService service.ApplicationsService) *applicationsHandler {
	return &applicationsHandler{
		applicationsService: applicationsService,
	}
}

func (h *applicationsHandler) GetAllApplications(c *fiber.Ctx) error {
	filterType := c.Query("type", "")
	userIDVal := c.Locals("userID")
	var userID int
	if userIDVal != nil {
		if id, ok := userIDVal.(float64); ok {
			userID = int(id)
		}
	}

	applications, err := h.applicationsService.GetAllApplications(c.Context(), userID, filterType)
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
		"message":    "Get all applications",
		"data":       applications,
	})
}

func (h *applicationsHandler) GetApplicationByID(c *fiber.Ctx) error {
	id := c.Params("id")
	userIDVal := c.Locals("userID")
	var userID int
	if userIDVal != nil {
		if id, ok := userIDVal.(float64); ok {
			userID = int(id)
		}
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid application ID",
		})
	}

	application, err := h.applicationsService.GetApplicationByID(c.Context(), userID, intID)
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
		"message":    "Get application data",
		"data":       application,
	})
}

func (h *applicationsHandler) ScreeningApprove(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid application ID",
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

	application, err := h.applicationsService.ScreeningApprove(c.Context(), int(userData.ID), intID)
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
		"message":    "Application approved by screening",
		"data":       application,
	})
}

func (h *applicationsHandler) ScreeningReject(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid application ID",
		})
	}

	var decision dto.ApplicationDecision
	err = c.BodyParser(&decision)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	decision.ApplicationID = intID

	// Get user data from context
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	application, err := h.applicationsService.ScreeningReject(c.Context(), int(userData.ID), decision)
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
		"message":    "Application rejected by screening",
		"data":       application,
	})
}

func (h *applicationsHandler) ScreeningRevise(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid application ID",
		})
	}

	var decision dto.ApplicationDecision
	err = c.BodyParser(&decision)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	decision.ApplicationID = intID

	// Get user data from context
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	application, err := h.applicationsService.ScreeningRevise(c.Context(), int(userData.ID), decision)
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
		"message":    "Application revision requested",
		"data":       application,
	})
}

func (h *applicationsHandler) FinalApprove(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid application ID",
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

	application, err := h.applicationsService.FinalApprove(c.Context(), int(userData.ID), intID)
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
		"message":    "Application approved by vendor",
		"data":       application,
	})
}

func (h *applicationsHandler) FinalReject(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid application ID",
		})
	}

	var decision dto.ApplicationDecision
	err = c.BodyParser(&decision)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	decision.ApplicationID = intID

	// Get user data from context
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	application, err := h.applicationsService.FinalReject(c.Context(), int(userData.ID), decision)
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
		"message":    "Application rejected by vendor",
		"data":       application,
	})
}
