package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
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
		initializer.SeedDatabase()
	}
}

//go:embed views/*
var viewsFS embed.FS

func main() {
	viewsRoot, err := fs.Sub(viewsFS, "views")
	if err != nil {
		log.Fatal(err)
	}
	engine := html.NewFileSystem(http.FS(viewsRoot), ".tmpl")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	Routes(app)

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
