package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/keijoraamat/mka_register/initializer"
	"github.com/keijoraamat/mka_register/models"
)

func FindingsIndex(c *fiber.Ctx) error {

	var acts []models.FindingAct
	result := initializer.DB.Find(&acts)
	if result.Error != nil {
		log.Println("Error getting Acts.")
	}

	return c.Render("findings/index", fiber.Map{
		"Acts": &acts,
	})
}

func FindingsNewFinding(c *fiber.Ctx) error {

	var acts []models.FindingAct
	result := initializer.DB.Find(&acts)
	if result.Error != nil {
		log.Println("Error getting Acts.")
	}

	return c.Render("findings/addFinding", fiber.Map{})
}

func FindingsFetchFinding(c *fiber.Ctx) error {

	var acts []models.FindingAct
	result := initializer.DB.Find(&acts)
	if result.Error != nil {
		log.Println("Error getting Acts.")
	}

	return c.Render("findings/addFinding", fiber.Map{
		"Acts": &acts,
	})
}

func FindingsCreateFinding(c *fiber.Ctx) error {
	var body models.FindingAct
	var errors = make(map[string]string)

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	if len(body.WDActNumber) == 0 {
		log.Printf("No WD act number found. Creating finding act failed.")
		errors["wdError"] = "WD akti number puudu"
		return c.Render("findings/addFinding", fiber.Map{
			"Errors": &errors,
			"Act":    &body,
		})
	}

	result := initializer.DB.Create(&body)
	if result.Error != nil {
		log.Println("could no save finding act")
	}

	return c.Render("findings/addFinding", fiber.Map{
		"Act": &body,
	})
}

func FindingsAddLocation(c *fiber.Ctx) error {

}
