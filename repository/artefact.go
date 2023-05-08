package repository

import (
	"log"

	"github.com/keijoraamat/mka_register/models"
	"gorm.io/gorm"
)

func AddArtefactToLocation(a *models.Artefact, l *models.Location, db *gorm.DB) {
	result := db.Create(&a)
	if result.Error != nil {
		log.Println("Could no save artefact: ", &a)
		log.Println("artefact saving error: ", result.Error)
	}

	db.Model(l).Association("Artefacts").Append(a)

	db.Commit()
}

func AddArtefact(a *models.Artefact, db *gorm.DB) *models.Artefact {

	result := db.Create(&a)
	if result.Error != nil {
		log.Println("Could no save finding location: ", &a)
		log.Panic("location saving error: ", result.Error)
	}
	return a
}

func AddArtefactLocation(al *models.ArtefactLocation, db *gorm.DB) {

	result := db.Create(&al)
	if result.Error != nil {
		log.Println("Could not save artefactLocation: ", &al)
		log.Panic("AddArtefactLocation() error: ", result.Error)
	}
}

func GetAllArtefactsByLocationID(loc_id uint, db *gorm.DB) []models.Artefact {
	log.Println("GetAllArtefactsByLocationID() called")
	var a []models.Artefact
	err := db.
		Joins("JOIN artefact_locations ON artefact_id = artefact_locations.artefact_id").
		Where("artefact_locations.location_id = ?", loc_id).
		Find(&a).Error
	if err != nil {
		log.Printf("Error quering location artefacts with ID: %d", loc_id)
		log.Panic(err)
	}

	log.Println("----------------\n")
	log.Println("Got aretfacts for loc: ", loc_id, len(a))
	/*for i, art := range a {
		log.Println(i, art)
	}*/

	return a
}
