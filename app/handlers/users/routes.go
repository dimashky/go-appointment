package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"github.com/dimashky/go-appointment/app/repositories"
)

// CreateUserRoute creates a new user endpoint
func CreateUserRoute(userRepo repositories.UserRepository, logger zerolog.Logger) fiber.Handler {
	handler := NewHandler(userRepo, logger)
	return handler.CreateUser()
}

// GetUserRoute gets a user by ID endpoint
func GetUserRoute(userRepo repositories.UserRepository, logger zerolog.Logger) fiber.Handler {
	handler := NewHandler(userRepo, logger)
	return handler.GetUser()
}

// ListUsersRoute lists users endpoint
func ListUsersRoute(userRepo repositories.UserRepository, logger zerolog.Logger) fiber.Handler {
	handler := NewHandler(userRepo, logger)
	return handler.ListUsers()
}
