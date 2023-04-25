package initializer

import (
	"os"

	"github.com/keijoraamat/mka_register/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	var dsn string

	if os.Getenv("APP_ENV") == "dev" {
		dsn = os.Getenv("DEV_DB_URL")
		DB, err = gorm.Open(sqlite.Open("register.db"), &gorm.Config{})
	} else {
		dsn = os.Getenv("DB_URL")
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("DB connection failed")
	}
}

func SyncDB() {
	DB.AutoMigrate(
		&models.FindingAct{},
		&models.Location{},
		&models.FindingLocation{},
	)
}
