package handler

import (
	"net/http"
	"strconv"

	"UMKMGo-backend/internal/service"
	"UMKMGo-backend/internal/types/dto"

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

// ====================== Mobile Auth =================================

func (user_handler *usersHandler) GetMeta(c *fiber.Ctx) error {
	meta, err := user_handler.usersService.MetaCityAndProvince(c.Context())
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
		"message":    "Get meta data",
		"data":       meta,
	})
}

func (user_handler *usersHandler) RegisterMobile(c *fiber.Ctx) error {
	var userRequest dto.RegisterMobile
	err := c.BodyParser(&userRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	if err := user_handler.usersService.RegisterMobile(c.Context(), userRequest.Email, userRequest.Phone); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "User registered successfully, please send OTP to verify",
	})
}

func (user_handler *usersHandler) VerifyOTP(c *fiber.Ctx) error {
	var userRequest dto.RegisterMobile
	err := c.BodyParser(&userRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	tempToken, err := user_handler.usersService.VerifyOTP(c.Context(), userRequest.Phone, userRequest.OTPCode)
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
		"message":    "OTP sent successfully",
		"data": fiber.Map{
			"temp_token": tempToken,
			"phone":      userRequest.Phone,
		},
	})
}

func (user_handler *usersHandler) RegisterMobileProfile(c *fiber.Ctx) error {
	var userRequest dto.UMKMMobile
	err := c.BodyParser(&userRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	tempToken := c.Query("temp_token")
	if tempToken == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Temporary token is required",
		})
	}

	user, err := user_handler.usersService.RegisterMobileProfile(c.Context(), userRequest, tempToken)
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
		"message":    "User profile registered successfully",
		"data":       user,
	})
}

func (user_handler *usersHandler) LoginMobile(c *fiber.Ctx) error {
	var userRequest dto.UMKMMobile
	err := c.BodyParser(&userRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	token, err := user_handler.usersService.LoginMobile(c.Context(), userRequest)
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

func (user_handler *usersHandler) ForgotPassword(c *fiber.Ctx) error {
	phone := c.Query("phone")

	if err := user_handler.usersService.ForgotPassword(c.Context(), phone); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": 500,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Reset password request sent, please send OTP to verify",
	})
}

func (user_handler *usersHandler) ResetPassword(c *fiber.Ctx) error {
	var resetRequest dto.ResetPasswordMobile
	err := c.BodyParser(&resetRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	tempToken := c.Query("temp_token")
	if err := user_handler.usersService.ResetPassword(c.Context(), resetRequest, tempToken); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": 500,
			"status":     false,
			"message":    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"status":     true,
		"message":    "Password reset successfully",
	})
}
