package repository_test

import (
	"fmt"
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
	db, _ = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	db.AutoMigrate(
		&models.FindingAct{},
		&models.Location{},
		&models.FindingLocation{},
	)
}

func cleanup() {
	fmt.Println("Removing DB after tests")
	if err := os.Remove(dbFile); err != nil {
		fmt.Printf("Removing DB file, %s failed.\n", dbFile)
	}
}
