package initializer

import (
	"log"
	"os"
	"time"

	"github.com/keijoraamat/mka_register/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	var dsn string

	if os.Getenv("APP_ENV") == "dev" {
		dsn = os.Getenv("DEV_DB_URL")
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Warn, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,        // Don't include params in the SQL log
				Colorful:                  false,       // Disable color
			},
		)
		DB, err = gorm.Open(sqlite.Open("register.db"), &gorm.Config{Logger: newLogger})
		if err != nil {
			panic("DB connection failed")
		}
	}

	if os.Getenv("APP_ENV") == "test" {
		dsn = os.Getenv("DB_URL")
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("DB connection failed")
		}
	}

	_, ok := os.LookupEnv("APP_ENV")
	if !ok {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Warn, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,        // Don't include params in the SQL log
				Colorful:                  false,       // Disable color
			},
		)
		DB, err = gorm.Open(sqlite.Open("register.db"), &gorm.Config{Logger: newLogger})
		if err != nil {
			panic("DB connection failed")
		}
	}

}

func SyncDB() {
	DB.AutoMigrate(
		&models.Artefact{},
		&models.FindingAct{},
		&models.Location{},
		&models.FindingLocation{},
		&models.ArtefactLocation{},
	)
}
