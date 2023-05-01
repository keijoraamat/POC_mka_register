package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/keijoraamat/mka_register/controllers"
	"github.com/stretchr/testify/assert"
)

func TestFindings_Index_Should_Contain_List_Of_Findig_Acts(t *testing.T) {
	engine := html.New(".. /views", ".tmpl")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	var findingsActController controllers.FindingActController = controllers.FindingActController{DB: db}

	app.Get("/leidmine", findingsActController.FindingsIndex)

	req := httptest.NewRequest(http.MethodGet, "/leidmine", nil)
	res, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)

}
