package controllers

import (
	"log"

	"github.com/keijoraamat/mka_register/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func SaveAct(act *models.Act) (ID uint, err error) {
	log.Println("SaveAct called")
	result := db.Create(&act)

	err = result.Error
	ID = act.ID

	return
}
