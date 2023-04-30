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
