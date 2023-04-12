package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/keijoraamat/mka_register/initializer"
	"github.com/keijoraamat/mka_register/models"
	"github.com/keijoraamat/mka_register/repository"
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

	var act models.FindingAct
	result := initializer.DB.Find(&act, c.Params("id"))
	if result.Error != nil {
		log.Println("Error getting Act with id ", c.Params("id"))
	}

	return c.Render("findings/findingActOverview", fiber.Map{
		"Act": &act,
	})
}

func FindingsCreateFinding(c *fiber.Ctx) error {
	var body models.FindingAct
	var errors = make(map[string]string)

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	//TODO: validate all other fields also
	if len(body.WDActNumber) == 0 {
		log.Println("No WD act number found. Creating finding act failed.")
		log.Println("leidja: ", body.FinderName)

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

	log.Printf("Finding act %d added to DB", body.ID)
	r := "/leidmine/akt/" + strconv.FormatUint(uint64(body.ID), 10)
	log.Printf("Redirecting to %s", r)
	c.Redirect(r)

	return c.Render("findings/addLocationToFinding", fiber.Map{
		"Act": &body,
	})
}

func FindingsAddLocation(c *fiber.Ctx) error {
	var act models.FindingAct
	/*result := initializer.DB.Find(&act, c.Params("id"))
	if result.Error != nil {
		log.Println("Error getting Act with id ", c.Params("id"))
	}*/
	err := initializer.DB.First(&act, c.Params("id")).Error
	if err != nil {
		log.Println("could not get act by id ", err)
	}

	log.Println("Got location data: ", string(c.Body()))

	var loc models.Location
	var findingLocation models.FindingLocation
	if err := c.BodyParser(&loc); err != nil {
		log.Println("Could not parse finding location data")
		return err
	}

	if loc.IsLocationInputOK() {
		result := initializer.DB.Create(&loc)
		if result.Error != nil {
			log.Println("could no save finding location: ", &loc)
		}
		log.Printf("Location county %s with id %d added to DB", loc.County, loc.ID)

	}

	err = initializer.DB.First(&loc, &loc.ID).Error
	if err != nil {
		log.Println("could not get loc by id ", err)
	}

	//err = initializer.DB.Model(&act).Association("Locations").Append(&loc).Error
	if err != nil {
		log.Println("could not get act by id ", err)
	}
	findingLocation.FindingAct = act
	findingLocation.Location = loc

	result := initializer.DB.Create(&findingLocation)
	if result.Error != nil {
		log.Println("Could not save to DB findingLocation:", findingLocation)
	}

	log.Printf("finding_location with id %d added to DB", findingLocation.ID)

	var locs []models.Location

	locs, _ = repository.GetLocationsByFindingActID(c.Params("id"))

	return c.Render("findings/addLocationToFinding", fiber.Map{
		"Act":  &act,
		"Loc":  &loc,
		"Locs": &locs,
	})
}

func FindingsFetchFindingLocationAdding(c *fiber.Ctx) error {
	var act models.FindingAct
	var locs []models.Location

	result := initializer.DB.Find(&act, c.Params("id"))
	if result.Error != nil {
		log.Println("Error getting Act with id ", c.Params("id"))
		return result.Error
	}

	locs, _ = repository.GetLocationsByFindingActID(c.Params("id"))

	log.Println("Act with location: ", &act.ID)

	return c.Render("findings/addLocationToFinding", fiber.Map{
		"Act":  &act,
		"Locs": &locs,
	})
}

func FindingsRemoveLocation(c *fiber.Ctx) error {

	log.Printf("Removing location %s from act %s", c.Params("loc_id"), c.Params("id"))

	err := repository.RemoveLocationByID(c.Params("loc_id"))
	if err != nil {
		return err
	}

	return c.Redirect("/leidmine/akt/" + c.Params("id") + "/lisa_asukoht")
}
