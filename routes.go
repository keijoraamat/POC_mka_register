package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keijoraamat/mka_register/controllers"
	"github.com/keijoraamat/mka_register/initializer"
)

func Routes(app *fiber.App) {
	var findingsActController controllers.FindingActController = controllers.FindingActController{DB: initializer.DB}

	app.Get("/", controllers.MainIndex)
	app.Get("/leidmine", findingsActController.Index)
	app.Get("/leidmine/akt/:id", findingsActController.GetFindingByID)
	app.Get("/leidmine/akt/:id/lisa_asukoht", findingsActController.FetchFindingLocationAdding)
	app.Post("/leidmine/akt/:id/lisa_asukoht", findingsActController.AddLocation)
	app.Post("/leidmine/akt/:id/eemalda_asukoht/:loc_id", findingsActController.RemoveLocation)
	app.Get("/leidmine/lisa_akt", findingsActController.NewFinding)
	app.Post("/leidmine/lisa", findingsActController.CreateFinding)
	app.Get("/konserveerimine", controllers.ConservationsIndex)
	app.Get("/ekspertimine", controllers.AssesmentsIndex)
}
