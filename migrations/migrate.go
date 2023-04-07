package main

import (
	"github.com/keijoraamat/mka_register/initializer"
	"github.com/keijoraamat/mka_register/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
}

func main() {
	initializer.DB.AutoMigrate(&models.Act{})
}
