package controllers

import "github.com/gofiber/fiber/v2"

func AssesmentsIndex(c *fiber.Ctx) error {
	return c.Render("assesments/index", fiber.Map{})
}
