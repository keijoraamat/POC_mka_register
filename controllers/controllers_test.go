package controllers_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/keijoraamat/mka_register/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbFile = "register_test.db"

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func setup() {
	fmt.Println("Setting up DB")
	var err error
	db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Println("Err setting up db", err)
	}
	db.AutoMigrate(
		&models.FindingAct{},
		&models.Location{},
		&models.FindingLocation{},
	)

	seedDB()
}

func cleanup() {
	fmt.Println("Removing DB after tests")
	if err := os.Remove(dbFile); err != nil {
		fmt.Printf("Removing DB file, %s failed.\n", dbFile)
	}
}

func seedDB() {
	var acts = make([]models.FindingAct, 3)

	acts[0] = *firstFindingAct
	acts[1] = *secondFindingAct

	db.Create(&acts)

}

var firstFindingAct = &models.FindingAct{
	FinderName:       "Mari Maasikas",
	FinderIdNumber:   "EE12345667890",
	RecieverName:     "Tuule Tallerma",
	FindingType:      "juhu",
	FindersFee:       false,
	ResiginOwnership: true,
	RemainAnonymous:  false,
	TransferLocation: "Fellin",
	WDActNumber:      "33.66453",
	Status:           "töös",
}
var secondFindingAct = &models.FindingAct{
	FinderName:       "Mati Kurikas",
	FinderIdNumber:   "EE2345678900",
	RecieverName:     "Tuule Tallerma",
	FindingType:      "detektor",
	FindersFee:       true,
	ResiginOwnership: true,
	RemainAnonymous:  false,
	TransferLocation: "Dorpat",
	WDActNumber:      "33.234523",
	Status:           "arhiiv",
}
