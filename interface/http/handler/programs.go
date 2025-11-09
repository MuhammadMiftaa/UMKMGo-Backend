package handler

import (
	"net/http"
	"strconv"

	"sapaUMKM-backend/internal/service"
	"sapaUMKM-backend/internal/types/dto"

	"github.com/gofiber/fiber/v2"
)

type programsHandler struct {
	programsService service.ProgramsService
}

func NewProgramsHandler(programsService service.ProgramsService) *programsHandler {
	return &programsHandler{
		programsService: programsService,
	}
}

func (h *programsHandler) GetAllPrograms(c *fiber.Ctx) error {
	programs, err := h.programsService.GetAllPrograms(c.Context())
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
		"message":    "Get all programs data",
		"data":       programs,
	})
}

func (h *programsHandler) GetProgramByID(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid program ID",
		})
	}

	program, err := h.programsService.GetProgramByID(c.Context(), intID)
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
		"message":    "Get program data",
		"data":       program,
	})
}

func (h *programsHandler) CreateProgram(c *fiber.Ctx) error {
	var programRequest dto.Programs
	err := c.BodyParser(&programRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	// Get user data from context (set by auth middleware)
	userData, ok := c.Locals("user_data").(dto.UserData)
	if ok {
		programRequest.CreatedBy = int(userData.ID)
	}

	program, err := h.programsService.CreateProgram(c.Context(), programRequest)
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
		"message":    "Create program data",
		"data":       program,
	})
}

func (h *programsHandler) UpdateProgram(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid program ID",
		})
	}

	var programRequest dto.Programs
	err = c.BodyParser(&programRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    err.Error(),
		})
	}

	program, err := h.programsService.UpdateProgram(c.Context(), intID, programRequest)
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
		"message":    "Update program data",
		"data":       program,
	})
}

func (h *programsHandler) DeleteProgram(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid program ID",
		})
	}

	program, err := h.programsService.DeleteProgram(c.Context(), intID)
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
		"message":    "Delete program data",
		"data":       program,
	})
}

func (h *programsHandler) ActivateProgram(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid program ID",
		})
	}

	program, err := h.programsService.ActivateProgram(c.Context(), intID)
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
		"message":    "Activate program",
		"data":       program,
	})
}

func (h *programsHandler) DeactivateProgram(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": 400,
			"status":     false,
			"message":    "Invalid program ID",
		})
	}

	program, err := h.programsService.DeactivateProgram(c.Context(), intID)
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
		"message":    "Deactivate program",
		"data":       program,
	})
}
