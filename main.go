package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/keijoraamat/mka_register/initializer"
)

func init() {
	initializer.LoadEnvVariables()
}

func main() {
	engine := html.New("./views", ".tmpl")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "/public")

	Routes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
