package repository

import (
	"log"
	"strconv"

	"github.com/keijoraamat/mka_register/models"
	"gorm.io/gorm"
)

func GetLocationById(id string, db *gorm.DB) models.Location {
	var l models.Location

	result := db.Find(&l, id)
	if result.Error != nil {
		log.Println("Error getting Location with id: ", id)
	}

	return l
}

func GetAllLocations(db *gorm.DB) (loc []models.Location, err error) {
	result := db.Find(&loc)
	if result.Error != nil {
		log.Println("Could not get all locations")
	}
	return
}

func GetLocationsByFindingActID(actID string, db *gorm.DB) ([]models.Location, error) {
	var locs []models.Location

	id, err := strconv.ParseUint(actID, 0, 20)
	if err != nil {
		log.Println("Error converting id from url to int")
	}

	err = db.
		Joins("JOIN finding_locations ON locations.id = finding_locations.location_id").
		Where("finding_locations.finding_act_id = ?", id).
		Find(&locs).Error
	if err != nil {
		log.Printf("Error quering locations for finding act %s", actID)
		return nil, err
	}

	return locs, nil
}

func AddLocation(l models.Location, db *gorm.DB) (models.Location, error) {

	result := db.Create(&l)
	if result.Error != nil {
		log.Println("could no save finding location: ", &l)
		log.Println("location saving error: ", result.Error)
	}

	return l, nil
}

func RemoveLocationByID(id string, db *gorm.DB) error {

	result := db.Delete(&models.Location{}, id)
	if result.Error != nil {
		log.Println("could not remove loction by id: ", id)
		return result.Error
	}

	db.Model(&models.FindingLocation{}).Association("Locations").Delete(&models.Location{}, id)

	return nil
}
