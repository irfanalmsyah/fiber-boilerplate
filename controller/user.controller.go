package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irfanalmsyah/fiber-boilerplate/model"
)

// GetMe function returns the user object of the authenticated user.
// It takes a fiber.Ctx pointer as a parameter and returns an error.
// The user object is retrieved from the fiber context's locals.
// The function returns a JSON response with the user object and a success status.
func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}