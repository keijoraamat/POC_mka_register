package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/keijoraamat/mka_register/helpers"
	"github.com/keijoraamat/mka_register/models"
	"github.com/keijoraamat/mka_register/repository"
	"gorm.io/gorm"
)

type FindingActController struct {
	DB *gorm.DB
}

func (fac *FindingActController) Index(c *fiber.Ctx) error {

	var acts []models.FindingAct
	var viewableActs []models.FindingActView
	var artefacts int
	result := fac.DB.Find(&acts)
	if result.Error != nil {
		log.Println("Error getting Acts.")
	}

	for _, act := range acts {
		viewableActs = append(viewableActs, act.DataToTemplate())
		var locs []models.Location
		locs, _ = repository.GetLocationsByFindingActID(fmt.Sprint(act.ID), fac.DB)
		for _, loc := range locs {
			artefacts = artefacts + int(loc.FindingsAmount)
		}
	}

	return c.Render("findings/index", fiber.Map{
		"Acts":      &viewableActs,
		"Artefacts": &artefacts,
	})
}

func (fac *FindingActController) NewFinding(c *fiber.Ctx) error {

	return c.Render("findings/addFinding", fiber.Map{})
}

func (fac *FindingActController) GetFindingByID(c *fiber.Ctx) error {

	var act models.FindingAct
	result := fac.DB.Find(&act, c.Params("id"))
	if result.Error != nil {
		log.Println("Error getting Act with id ", c.Params("id"))
	}

	var locs []models.Location

	locs, _ = repository.GetLocationsByFindingActID(c.Params("id"), fac.DB)

	for i := range locs {
		locs[i].Afacts = append(locs[i].Afacts, repository.GetAllArtefactsByLocationID(locs[i].ID, fac.DB)...)
	}

	return c.Render("findings/addLocationToFinding", fiber.Map{
		"Act":  act.DataToTemplate(),
		"Locs": &locs,
	})
}

func (fac *FindingActController) CreateFinding(c *fiber.Ctx) error {
	var findingAct models.FindingAct
	var errors = make(map[string]string)

	findingAct.FinderName = c.FormValue("finderName")
	findingAct.FinderIdNumber = c.FormValue("finderIdNumber")
	findingAct.RecieverName = c.FormValue("recieverName")
	findingAct.FindingType = c.FormValue("findingType")
	findingAct.FindersFee = helpers.ParseCheckBox(c.FormValue("findersFee"))
	findingAct.ResiginOwnership = helpers.ParseCheckBox(c.FormValue("resiginOwnership"))
	findingAct.RemainAnonymous = helpers.ParseCheckBox(c.FormValue("remainAnonymous"))
	findingAct.TransferLocation = c.FormValue("transferLocation")
	findingAct.TransferDate = helpers.ParseDate(c.FormValue("transferDate"))
	findingAct.WDActNumber = c.FormValue("wdActNumber")

	//TODO: validate all other fields also
	if len(findingAct.WDActNumber) == 0 {
		log.Println("No WD act number found. Creating finding act failed.")

		errors["wdError"] = "WD akti number puudu"
		return c.Render("findings/addFinding", fiber.Map{
			"Errors": &errors,
			"Act":    &findingAct,
		})
	}

	result := fac.DB.Create(&findingAct)
	if result.Error != nil {
		log.Println("could no save finding act")
	}

	redirectUrl := fmt.Sprintf("/leidmine/akt/%d", findingAct.ID)
	return c.Redirect(redirectUrl)
}

func (fac *FindingActController) AddLocation(c *fiber.Ctx) error {
	var act models.FindingAct
	err := fac.DB.First(&act, c.Params("id")).Error
	if err != nil {
		log.Println("could not get act by id ", err)
	}

	var loc models.Location
	var findingLocation models.FindingLocation
	if err := c.BodyParser(&loc); err != nil {
		log.Println("Could not parse finding location data")
		return err
	}

	if loc.IsLocationInputOK() {
		result := fac.DB.Create(&loc)
		if result.Error != nil {
			log.Println("could no save finding location: ", &loc)
		}
		log.Printf("Location county %s with id %d added to DB", loc.County, loc.ID)

	}

	err = fac.DB.First(&loc, &loc.ID).Error
	if err != nil {
		log.Println("could not get loc by id ", err)
	}

	findingLocation.FindingAct = act
	findingLocation.Location = loc

	result := fac.DB.Create(&findingLocation)
	if result.Error != nil {
		log.Println("Could not save to DB findingLocation:", findingLocation)
	}

	log.Printf("finding_location with id %d added to DB", findingLocation.ID)

	var locs []models.Location

	locs, _ = repository.GetLocationsByFindingActID(c.Params("id"), fac.DB)

	for i := range locs {
		locs[i].Afacts = append(locs[i].Afacts, repository.GetAllArtefactsByLocationID(locs[i].ID, fac.DB)...)
	}

	return c.Render("findings/addLocationToFinding", fiber.Map{
		"Act":  act.DataToTemplate(),
		"Loc":  &loc,
		"Locs": &locs,
	})
}

func (fac *FindingActController) FetchFindingLocationAdding(c *fiber.Ctx) error {
	var act models.FindingAct
	var locs []models.Location

	repository.GetFindingActById(&act, c.Params("id"), fac.DB)
	locs, _ = repository.GetLocationsByFindingActID(c.Params("id"), fac.DB)

	for i := range locs {
		locs[i].Afacts = append(locs[i].Afacts, repository.GetAllArtefactsByLocationID(locs[i].ID, fac.DB)...)
	}

	return c.Render("findings/addLocationToFinding", fiber.Map{
		"Act":  act.DataToTemplate(),
		"Locs": &locs,
	})
}

func (fac *FindingActController) RemoveLocation(c *fiber.Ctx) error {

	log.Printf("Removing location %s from act %s", c.Params("loc_id"), c.Params("id"))
	repository.RemoveLocationByID(c.Params("loc_id"), fac.DB)

	return c.Redirect("/leidmine/akt/" + c.Params("id") + "/lisa_asukoht")
}

func (fac *FindingActController) AddArtefact(c *fiber.Ctx) error {
	log.Println("Getting AddArtefact()")
	params := struct {
		ActId string
		LocId string
	}{
		ActId: c.Params("act_id"),
		LocId: c.Params("loc_id"),
	}
	return c.Render("findings/addArtefact", fiber.Map{
		"data": params,
	})
}

func (fac *FindingActController) SaveArtefact(c *fiber.Ctx) error {
	log.Println("____SaveArtefact() called_________")
	actID := c.Params("act_id")
	locId := c.Params("loc_id")
	var artefactLocation models.ArtefactLocation

	loc := repository.GetLocationById(locId, fac.DB)

	var artefact models.Artefact
	if err := c.BodyParser(&artefact); err != nil {
		log.Println("Could not parse artefact")
		return err
	}

	repository.AddArtefact(&artefact, fac.DB)

	artefactLocation.Artefact = artefact
	artefactLocation.Location = loc

	repository.AddArtefactLocation(&artefactLocation, fac.DB)

	var act models.FindingAct
	var locs []models.Location

	repository.GetFindingActById(&act, actID, fac.DB)
	locs, _ = repository.GetLocationsByFindingActID(actID, fac.DB)
	log.Println("locs len after getting from db: ", len(locs))
	for i := range locs {
		locs[i].Afacts = append(locs[i].Afacts, repository.GetAllArtefactsByLocationID(locs[i].ID, fac.DB)...)
	}

	log.Println(".....SaceArtefact() end.....")
	return c.Render("findings/addLocationToFinding", fiber.Map{
		"Act":  act.DataToTemplate(),
		"Locs": &locs,
	})
}

func (fac *FindingActController) GetFindingActPDF(c *fiber.Ctx) error {
	log.Println("Creating PDF for finding act with id: ", c.Params("id"))
	return nil
}
