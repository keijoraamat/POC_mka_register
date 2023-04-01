package controllers

import "github.com/gofiber/fiber/v2"

func ConservationsIndex(c *fiber.Ctx) error {
	return c.Render("conservations/index", fiber.Map{})
}
