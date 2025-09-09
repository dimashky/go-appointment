package users

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"github.com/dimashky/go-appointment/app/models"
	"github.com/dimashky/go-appointment/app/repositories"
)

// Handler holds dependencies for user handlers
type Handler struct {
	userRepo repositories.UserRepository
	logger   zerolog.Logger
}

// NewHandler creates a new user handler
func NewHandler(userRepo repositories.UserRepository, logger zerolog.Logger) *Handler {
	return &Handler{
		userRepo: userRepo,
		logger:   logger,
	}
}

// CreateUser creates a new user
func (h *Handler) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			h.logger.Error().Err(err).Msg("Failed to parse user data")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if err := h.userRepo.Create(&user); err != nil {
			h.logger.Error().Err(err).Msg("Failed to create user")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}

		h.logger.Info().Uint("user_id", user.ID).Msg("User created successfully")
		return c.Status(fiber.StatusCreated).JSON(user)
	}
}

// GetUser retrieves a user by ID
func (h *Handler) GetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		user, err := h.userRepo.GetByID(uint(id))
		if err != nil {
			h.logger.Error().Err(err).Uint64("user_id", id).Msg("User not found")
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		return c.JSON(user)
	}
}

// ListUsers retrieves a paginated list of users
func (h *Handler) ListUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))

		users, err := h.userRepo.List(limit, offset)
		if err != nil {
			h.logger.Error().Err(err).Msg("Failed to retrieve users")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve users",
			})
		}

		return c.JSON(fiber.Map{
			"users":  users,
			"limit":  limit,
			"offset": offset,
		})
	}
}
