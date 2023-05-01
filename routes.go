package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keijoraamat/mka_register/controllers"
	"github.com/keijoraamat/mka_register/initializer"
) 

func Routes(app *fiber.App) {
	var findingsActController controllers.FindingActController = controllers.FindingActController{DB: initializer.DB}

	app.Get("/", controllers.MainIndex)
	app.Get("/leidmine", findingsActController.FindingsIndex)
	app.Get("/leidmine/akt/:id", controllers.FindingsFetchFinding)
	app.Get("/leidmine/akt/:id/lisa_asukoht", controllers.FindingsFetchFindingLocationAdding)
	app.Post("/leidmine/akt/:id/lisa_asukoht", controllers.FindingsAddLocation)
	app.Post("/leidmine/akt/:id/eemalda_asukoht/:loc_id", controllers.FindingsRemoveLocation)
	app.Get("/leidmine/lisa_akt", controllers.FindingsNewFinding)
	app.Post("/leidmine/lisa", controllers.FindingsCreateFinding)
	app.Get("/konserveerimine", controllers.ConservationsIndex)
	app.Get("/ekspertimine", controllers.AssesmentsIndex)
}
