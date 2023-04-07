package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keijoraamat/mka_register/controllers"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.MainIndex)
	app.Get("/leidmine", controllers.FindingsIndex)
	app.Get("/leidmine/lisa", controllers.FindingsNewFinding)
	app.Post("/leidmine/lisa", controllers.FindingsCreateFinding)
	app.Get("/konserveerimine", controllers.ConservationsIndex)
	app.Get("/ekspertimine", controllers.AssesmentsIndex)
}
