package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keijoraamat/mka_register/controllers"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.MainIndex)
}