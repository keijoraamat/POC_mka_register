package repository

import (
	"log"
	"strconv"

	"github.com/keijoraamat/mka_register/initializer"
	"github.com/keijoraamat/mka_register/models"
)

func GetLocationsByFindingActID(actID string) ([]models.Location, error) {
	var locs []models.Location

	id, err := strconv.ParseUint(actID, 0, 20)
	if err != nil {
		log.Println("Error converting id from url to int")
	}

	err = initializer.DB.
		Joins("JOIN finding_locations ON locations.id = finding_locations.location_id").
		Where("finding_locations.finding_act_id = ?", id).
		Find(&locs).Error
	if err != nil {
		return nil, err
	}
	log.Println("Locations by id:", locs)

	return locs, nil
}

func AddLocation(l models.Location) (models.Location, error) {

	result := initializer.DB.Create(&l)
	if result.Error != nil {
		log.Println("could no save finding location: ", &l)
		log.Println("location saving error: ", result.Error)
	}
	log.Printf("Location county %s with id %d added to DB", l.County, l.ID)

	return l, nil
}
