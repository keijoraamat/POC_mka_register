package controllers

import "github.com/gofiber/fiber/v2"

func FindingsIndex(c *fiber.Ctx) error {
	return c.Render("findings/index", fiber.Map{})
}
