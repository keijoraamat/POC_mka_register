package controllers

import "github.com/gofiber/fiber/v2"

func MainIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
