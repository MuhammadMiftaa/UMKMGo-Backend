package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"sapaUMKM-backend/config/env"
	"sapaUMKM-backend/internal/service"
	"sapaUMKM-backend/internal/types/dto"
	"sapaUMKM-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type usersHandler struct {
	usersService service.UsersService
}

func NewUsersHandler(usersService service.UsersService) *usersHandler {
	return &usersHandler{
		usersService: usersService,
	}
}

func (user_handler *usersHandler) Register(c *fiber.Ctx) error {
	var userRequest dto.Users
	err := c.BodyParser(&userRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	user, err := user_handler.usersService.Register(c.Context(), userRequest)
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
		"message":    "Register user data",
		"data":       user,
	})
}

func (user_handler *usersHandler) Login(c *fiber.Ctx) error {
	var userRequest dto.Users
	err := c.BodyParser(&userRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	token, err := user_handler.usersService.Login(c.Context(), userRequest)
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
		"message":    "Login user data",
		"data":       token,
	})
}

func (user_handler *usersHandler) UpdateProfile(c *fiber.Ctx) error {
	var userRequest dto.Users
	err := c.BodyParser(&userRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	// Get user ID from JWT token
	userData, ok := c.Locals("user_data").(dto.UserData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Unauthorized",
		})
	}

	user, err := user_handler.usersService.UpdateProfile(c.Context(), int(userData.ID), userRequest)
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
		"message":    "Update profile success",
		"data":       user,
	})
}

func (user_handler *usersHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := user_handler.usersService.GetAllUsers(c.Context())
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
		"message":    "Get all users data",
		"data":       users,
	})
}

func (user_handler *usersHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid user ID",
		})
	}

	user, err := user_handler.usersService.GetUserByID(c.Context(), intID)
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
		"message":    "Get user data",
		"data":       user,
	})
}

func (user_handler *usersHandler) UpdateUser(c *fiber.Ctx) error {
	var userRequest dto.Users
	err := c.BodyParser(&userRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid user ID",
		})
	}

	user, err := user_handler.usersService.UpdateUser(c.Context(), intID, userRequest)
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
		"message":    "Update user data",
		"data":       user,
	})
}

func (user_handler *usersHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid user ID",
		})
	}

	user, err := user_handler.usersService.DeleteUser(c.Context(), intID)
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
		"message":    "Delete user data",
		"data":       user,
	})
}

func (user_handler *usersHandler) SendOTP(c *fiber.Ctx) error {
	var OTP dto.OTP
	if err := c.BodyParser(&OTP); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid request body",
		})
	}
	OTP.OTP = utils.GenerateOTP()

	// Simpan OTP ke Redis
	if err := user_handler.usersService.SetOTP(c.Context(), OTP.Email, OTP.OTP, 5*time.Minute); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": 500,
			"status":     false,
			"message":    "Failed to save OTP",
		})
	}

	// Kirimkan OTP ke email
	SMTPProvider := utils.NewZohoSMTP(env.Cfg.ZSMTP)
	if err := utils.NewSMTPClient(SMTPProvider).SendSingleEmail(OTP.Email, "OTP Verification", "otp-email-template.html", OTP); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": 500,
			"status":     false,
			"message":    fmt.Sprintf("Failed to send email, error: %v", err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     true,
		"statusCode": 200,
		"message":    "OTP sent successfully",
		"data":       OTP.Email,
	})
}

func (user_handler *usersHandler) VerifyOTP(c *fiber.Ctx) error {
	var OTP dto.OTP
	if err := c.BodyParser(&OTP); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	valid, err := user_handler.usersService.ValidateOTP(c.Context(), OTP.Email, OTP.OTP)
	if err != nil || !valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Invalid or expired OTP",
		})
	}

	user, err := user_handler.usersService.VerifyUser(c.Context(), OTP.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": 500,
			"status":     false,
			"message":    "Failed to verify user",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     true,
		"statusCode": 200,
		"message":    "OTP verified successfully",
		"data":       user,
	})
}

func (user_handler *usersHandler) GetListRolePermissions(c *fiber.Ctx) error {
	rolePermissions, err := user_handler.usersService.GetListRolePermissions(c.Context())
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
		"message":    "Get all role permissions data",
		"data":       rolePermissions,
	})
}

func (user_handler *usersHandler) GetListPermissions(c *fiber.Ctx) error {
	permissions, err := user_handler.usersService.GetListPermissions(c.Context())
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
		"message":    "Get all permissions data",
		"data":       permissions,
	})
}

func (user_handler *usersHandler) UpdateRolePermissions(c *fiber.Ctx) error {
	var rolePermissionsRequest dto.RolePermissions
	err := c.BodyParser(&rolePermissionsRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	err = user_handler.usersService.UpdateRolePermissions(c.Context(), rolePermissionsRequest)
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
		"message":    "Update role permissions",
	})
}
