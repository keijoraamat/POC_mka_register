package repository

import (
	"log"

	"github.com/keijoraamat/mka_register/models"
	"gorm.io/gorm"
)

func GetAllFindingActs(db *gorm.DB) (acts []models.FindingAct, err error) {
	result := db.Find(&acts)
	if result.Error != nil {
		log.Println("Error getting FindingActs.")
	}

	return
}

func GetFindingActById(act *models.FindingAct, id string, db *gorm.DB) (tx *gorm.DB){
	result := db.Find(&act, id)
	if result.Error != nil {
		log.Println("Error getting Act with id ", id)
	}

	return
}
