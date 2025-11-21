package handler

import (
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

// Upload Documents
func (h *MobileHandler) UploadNIB(c *fiber.Ctx) error {
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

	if err := h.mobileService.UploadNIB(c.Context(), int(userData.ID), request.Document); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "NIB document uploaded successfully",
	})
}

func (h *MobileHandler) UploadNPWP(c *fiber.Ctx) error {
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

	if err := h.mobileService.UploadNPWP(c.Context(), int(userData.ID), request.Document); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "NPWP document uploaded successfully",
	})
}

func (h *MobileHandler) UploadRevenueRecord(c *fiber.Ctx) error {
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

	if err := h.mobileService.UploadRevenueRecord(c.Context(), int(userData.ID), request.Document); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Revenue Record document uploaded successfully",
	})
}

func (h *MobileHandler) UploadBusinessPermit(c *fiber.Ctx) error {
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

	if err := h.mobileService.UploadBusinessPermit(c.Context(), int(userData.ID), request.Document); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Business Permit document uploaded successfully",
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

	application, err := h.mobileService.CreateTrainingApplication(c.Context(), int(userData.ID), request)
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
		"message":    "Training application created successfully",
		"data":       application,
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

	application, err := h.mobileService.CreateCertificationApplication(c.Context(), int(userData.ID), request)
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
		"message":    "Certification application created successfully",
		"data":       application,
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

	application, err := h.mobileService.CreateFundingApplication(c.Context(), int(userData.ID), request)
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
		"message":    "Funding application created successfully",
		"data":       application,
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

// GetNotificationsByUMKMID retrieves notifications for a specific UMKM ID.
func (h *MobileHandler) GetNotificationsByUMKMID(c *fiber.Ctx) error {
	umkmID, ok := c.Locals("userID").(int)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UMKM ID"})
	}

	notifications, err := h.mobileService.GetNotificationsByUMKMID(c.Context(), umkmID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve notifications"})
	}

	return c.JSON(notifications)
}

// GetUnreadCount retrieves the count of unread notifications for a specific UMKM ID.
func (h *MobileHandler) GetUnreadCount(c *fiber.Ctx) error {
	umkmID, ok := c.Locals("userID").(int)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UMKM ID"})
	}
	count, err := h.mobileService.GetUnreadCount(c.Context(), umkmID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve unread count"})
	}

	return c.JSON(fiber.Map{"unread_count": count})
}

// MarkNotificationsAsRead marks specified notifications as read for a specific UMKM ID.
func (h *MobileHandler) MarkNotificationsAsRead(c *fiber.Ctx) error {
	var request struct {
		NotificationIDs []int `json:"notification_ids"`
	}
	umkmID, ok := c.Locals("userID").(int)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UMKM ID"})
	}
	if err := h.mobileService.MarkNotificationsAsRead(c.Context(), umkmID, request.NotificationIDs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to mark notifications as read"})
	}

	return c.JSON(fiber.Map{"message": "Notifications marked as read"})
}

// MarkAllNotificationsAsRead marks all notifications as read for a specific UMKM ID.
func (h *MobileHandler) MarkAllNotificationsAsRead(c *fiber.Ctx) error {
	umkmID, ok := c.Locals("userID").(int)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UMKM ID"})
	}

	if err := h.mobileService.MarkAllNotificationsAsRead(c.Context(), umkmID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to mark all notifications as read"})
	}

	return c.JSON(fiber.Map{"message": "All notifications marked as read"})
}
