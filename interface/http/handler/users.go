package handler

import (
	"fmt"
	"net/http"
	"time"

	"sapaUMKM-backend/config/env"
	"sapaUMKM-backend/internal/service"
	"sapaUMKM-backend/internal/types/dto"
	"sapaUMKM-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type usersHandler struct {
	usersService service.UsersService
	otpService   service.OTPService
}

func NewUsersHandler(usersService service.UsersService, otpService service.OTPService) *usersHandler {
	return &usersHandler{
		usersService: usersService,
		otpService:   otpService,
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

	user, err := user_handler.usersService.Register(userRequest)
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

	token, err := user_handler.usersService.Login(userRequest)
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

func (user_handler *usersHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := user_handler.usersService.GetAllUsers()
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

	user, err := user_handler.usersService.GetUserByID(id)
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

	user, err := user_handler.usersService.UpdateUser(id, userRequest)
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

	user, err := user_handler.usersService.DeleteUser(id)
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
	if err := user_handler.otpService.SetOTP(OTP.Email, OTP.OTP, 5*time.Minute); err != nil {
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

	valid, err := user_handler.otpService.ValidateOTP(OTP.Email, OTP.OTP)
	if err != nil || !valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"status":     false,
			"message":    "Invalid or expired OTP",
		})
	}

	user, err := user_handler.usersService.VerifyUser(OTP.Email)
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