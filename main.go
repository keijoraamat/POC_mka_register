package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/keijoraamat/mka_register/initializer"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
	if os.Getenv("APP_ENV") == "dev" {

		initializer.SyncDB()
		initializer.SeedFindingsActs()
	}
}

func main() {
	engine := html.New("./views", ".tmpl")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	Routes(app)

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
