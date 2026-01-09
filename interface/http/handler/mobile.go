package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"UMKMGo-backend/internal/service"
	"UMKMGo-backend/internal/types/dto"

	"github.com/gofiber/fiber/v2"
)

type MobileHandler struct {
	mobileService service.MobileService
}

func NewMobileHandler(mobileService service.MobileService) *MobileHandler {
	return &MobileHandler{
		mobileService: mobileService,
	}
}

func (h *MobileHandler) GetDashboard(c *fiber.Ctx) error {
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	data, err := h.mobileService.GetDashboard(c.Context(), int(userData.ID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": 500,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Get user data",
		"data":       data,
	})
}

// Programs - Training
func (h *MobileHandler) GetTrainingPrograms(c *fiber.Ctx) error {
	programs, err := h.mobileService.GetTrainingPrograms(c.Context())
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
		"message":    "Get training programs",
		"data":       programs,
	})
}

// Programs - Certification
func (h *MobileHandler) GetCertificationPrograms(c *fiber.Ctx) error {
	programs, err := h.mobileService.GetCertificationPrograms(c.Context())
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
		"message":    "Get certification programs",
		"data":       programs,
	})
}

// Programs - Funding
func (h *MobileHandler) GetFundingPrograms(c *fiber.Ctx) error {
	programs, err := h.mobileService.GetFundingPrograms(c.Context())
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
		"message":    "Get funding programs",
		"data":       programs,
	})
}

// Program Detail
func (h *MobileHandler) GetProgramDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid program ID",
		})
	}

	program, err := h.mobileService.GetProgramDetail(c.Context(), intID)
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
		"message":    "Get program detail",
		"data":       program,
	})
}

// UMKM Profile
func (h *MobileHandler) GetUMKMProfile(c *fiber.Ctx) error {
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	profile, err := h.mobileService.GetUMKMProfile(c.Context(), int(userData.ID))
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
		"message":    "Get UMKM profile",
		"data":       profile,
	})
}

func (h *MobileHandler) UpdateUMKMProfile(c *fiber.Ctx) error {
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	var request dto.UpdateUMKMProfile
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	profile, err := h.mobileService.UpdateUMKMProfile(c.Context(), int(userData.ID), request)
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
		"message":    "Update UMKM profile success",
		"data":       profile,
	})
}

func (h *MobileHandler) GetUMKMDocuments(c *fiber.Ctx) error {
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	documents, err := h.mobileService.GetUMKMDocuments(c.Context(), int(userData.ID))
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
		"message":    "Get UMKM documents",
		"data":       documents,
	})
}

// Upload Document with dynamic type
func (h *MobileHandler) UploadDocument(c *fiber.Ctx) error {
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	var request dto.UploadDocumentRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	// Call service with dynamic document type
	if err := h.mobileService.UploadDocument(c.Context(), int(userData.ID), request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	// Generate dynamic success message
	docTypeNames := map[string]string{
		"nib":             "NIB",
		"npwp":            "NPWP",
		"revenue-record":  "Revenue Record",
		"business-permit": "Business Permit",
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    fmt.Sprintf("%s document uploaded successfully", docTypeNames[request.Type]),
	})
}

// Applications
func (h *MobileHandler) CreateTrainingApplication(c *fiber.Ctx) error {
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	var request dto.CreateApplicationTraining
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	if err := h.mobileService.CreateTrainingApplication(c.Context(), int(userData.ID), request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"statusCode": 201,
		"status":     true,
		"message":    "Training application created successfully",
	})
}

func (h *MobileHandler) CreateCertificationApplication(c *fiber.Ctx) error {
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	var request dto.CreateApplicationCertification
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	if err := h.mobileService.CreateCertificationApplication(c.Context(), int(userData.ID), request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"statusCode": 201,
		"status":     true,
		"message":    "Certification application created successfully",
	})
}

func (h *MobileHandler) CreateFundingApplication(c *fiber.Ctx) error {
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	var request dto.CreateApplicationFunding
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	if err := h.mobileService.CreateFundingApplication(c.Context(), int(userData.ID), request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"statusCode": 201,
		"status":     true,
		"message":    "Funding application created successfully",
	})
}

func (h *MobileHandler) GetApplicationList(c *fiber.Ctx) error {
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	applications, err := h.mobileService.GetApplicationList(c.Context(), int(userData.ID))
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
		"message":    "Get application list",
		"data":       applications,
	})
}

func (h *MobileHandler) GetApplicationDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid application ID",
		})
	}

	application, err := h.mobileService.GetApplicationDetail(c.Context(), intID)
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
		"message":    "Get application detail",
		"data":       application,
	})
}

func (h *MobileHandler) ReviseApplication(c *fiber.Ctx) error {
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid application ID",
		})
	}

	var request []dto.UploadDocumentRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	if err := h.mobileService.ReviseApplication(c.Context(), int(userData.ID), intID, request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Application revised successfully",
	})
}

// GetNotificationsByUMKMID retrieves notifications for a specific UMKM ID.
// $ For Backward Compatibility
func (h *MobileHandler) GetNotificationsByUMKMID(c *fiber.Ctx) error {
	umkmID, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UMKM ID"})
	}

	notifications, err := h.mobileService.GetNotificationsByUMKMID(c.Context(), int(umkmID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve notifications"})
	}

	return c.JSON(notifications)
}

// $ Get List Notifications Fix
func (h *MobileHandler) GetListNotificationsByUMKMID(c *fiber.Ctx) error {
	umkmID, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UMKM ID"})
	}

	notifications, err := h.mobileService.GetNotificationsByUMKMID(c.Context(), int(umkmID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve notifications"})
	}

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Get list notifications",
		"data":       notifications,
	})
}

// GetUnreadCount retrieves the count of unread notifications for a specific UMKM ID.
func (h *MobileHandler) GetUnreadCount(c *fiber.Ctx) error {
	umkmID, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UMKM ID"})
	}
	count, err := h.mobileService.GetUnreadCount(c.Context(), int(umkmID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve unread count"})
	}

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Get application detail",
		"data":       count,
	})
}

// MarkNotificationsAsRead marks specified notifications as read for a specific UMKM ID.
func (h *MobileHandler) MarkNotificationsAsRead(c *fiber.Ctx) error {
	umkmID, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UMKM ID"})
	}
	notificationID := c.Params("id")
	notificationIDInt, err := strconv.Atoi(notificationID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid notification ID"})
	}
	if err := h.mobileService.MarkNotificationsAsRead(c.Context(), int(umkmID), notificationIDInt); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to mark notifications as read"})
	}

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Notification marked as read",
		"data":       nil,
	})
}

// MarkAllNotificationsAsRead marks all notifications as read for a specific UMKM ID.
func (h *MobileHandler) MarkAllNotificationsAsRead(c *fiber.Ctx) error {
	umkmID, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UMKM ID"})
	}

	if err := h.mobileService.MarkAllNotificationsAsRead(c.Context(), int(umkmID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to mark all notifications as read"})
	}

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "All notifications marked as read",
		"data":       nil,
	})
}

// GetPublishedNews retrieves published news with optional filters.
func (h *MobileHandler) GetPublishedNews(c *fiber.Ctx) error {
	params := dto.NewsQueryParams{
		Page:     c.QueryInt("page", 1),
		Limit:    c.QueryInt("limit", 10),
		Category: c.Query("category"),
		Search:   c.Query("search"),
		Tag:      c.Query("tag"),
	}

	news, _, err := h.mobileService.GetPublishedNews(c.Context(), params)
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
		"message":    "Get published news",
		"data":       news,
	})
}

func (h *MobileHandler) GetNewsDetailBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "News slug is required",
		})
	}

	news, err := h.mobileService.GetNewsDetail(c.Context(), slug)
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
		"message":    "Get news detail",
		"data":       news,
	})
}
