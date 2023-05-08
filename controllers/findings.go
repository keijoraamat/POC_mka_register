package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/keijoraamat/mka_register/models"
	"github.com/keijoraamat/mka_register/repository"
	"gorm.io/gorm"
)

type FindingActController struct {
	DB *gorm.DB
}

func (fac *FindingActController) Index(c *fiber.Ctx) error {

	var acts []models.FindingAct
	result := fac.DB.Find(&acts)
	if result.Error != nil {
		log.Println("Error getting Acts.")
	}

	return c.Render("findings/index", fiber.Map{
		"Acts": &acts,
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

	return c.Render("findings/findingActOverview", fiber.Map{
		"Act": &act,
	})
}

func (fac *FindingActController) CreateFinding(c *fiber.Ctx) error {
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

	result := fac.DB.Create(&body)
	if result.Error != nil {
		log.Println("could no save finding act")
	}

	redirectUrl := fmt.Sprintf("/leidmine/akt/%d", body.ID)
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
	log.Println("locs len after getting from db: ", len(locs))
	for i, l := range locs {
		log.Printf("%d. Looking for location: %d artefacts\n", i, l.ID)
		var artecats = repository.GetAllArtefactsByLocationID(l.ID, fac.DB)
		l.Afacts = artecats
		/*log.Println("loc esimene artefact: ", l.Artefacts[0])
		log.Println("esimene artefact: ", artecats[0])
		for j, art := range l.Artefacts {
			log.Println("Loc art", art.ID, j)
		}*/

	}

	return c.Render("findings/addLocationToFinding", fiber.Map{
		"Act":  &act,
		"Loc":  &loc,
		"Locs": &locs,
	})
}

func (fac *FindingActController) FetchFindingLocationAdding(c *fiber.Ctx) error {
	var act models.FindingAct
	var locs []models.Location

	repository.GetFindingActById(&act, c.Params("id"), fac.DB)
	locs, _ = repository.GetLocationsByFindingActID(c.Params("id"), fac.DB)

	log.Println("locs len after getting from db: ", len(locs))
	/*	for i, l := range locs {
				err := fac.DB.
					Joins("JOIN artefact_locations ON artefact_id = artefact_locations.artefact_id").
					Where("artefact_locations.location_id = ?", l.ID).
					Find(&l.Afacts).Error
				if err != nil {
					log.Printf("Error quering location artefacts with ID: %d", l.ID)
					log.Panic(err)
				}
				log.Println("i", i)
				//log.Println(l.Afacts)
				log.Println("ii", i)
			log.Printf("%d. Looking for location: %d artefacts\n", i, l.ID)
				var artecats = repository.GetAllArtefactsByLocationID(l.ID, fac.DB)
				l.Afacts = artecats
				log.Println("loc esimene artefact: ", l.Artefacts[0])
				log.Println("esimene artefact: ", artecats[0])
				for j, art := range l.Artefacts {
					log.Println("Loc art", art.ID, j)
				}

			}
			log.Println("loc 1", locs)
		//log.Println("loc 1, arf", repository.GetAllArtefactsByLocationID(locs[0].ID, fac.DB))
		locs[0].Afacts[0] = models.Artefact{Name: "Kee",
			YCoord: "YYY",
			XCoord: "XXX",
			Amount: 2,
		}*/

	return c.Render("findings/addLocationToFinding", fiber.Map{
		"Act":  &act,
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
	log.Println("locs len before getting from db: ", len(locs))

	repository.GetFindingActById(&act, actID, fac.DB)
	locs, _ = repository.GetLocationsByFindingActID(actID, fac.DB)
	log.Println("locs len after getting from db: ", len(locs))
	for i, l := range locs {
		log.Printf("%d. Looking for location: %d artefacts\n", i, l.ID)
		//var artecats = repository.GetAllArtefactsByLocationID(l.ID, fac.DB)
		l.Afacts = append(l.Afacts, models.Artefact{Name: "Kee",
			YCoord: "YYY",
			XCoord: "XXX",
			Amount: 2,
		})
		/*log.Println("loc esimene artefact: ", l.Artefacts[0])
		log.Println("esimene artefact: ", artecats[0])
		for j, art := range l.Artefacts {
			log.Println("Loc art", art.ID, j)
		}*/

	}

	log.Println(".....SaceArtefact() end.....")
	return c.Render("findings/addLocationToFinding", fiber.Map{
		"Act":  &act,
		"Locs": &locs,
	})
}
