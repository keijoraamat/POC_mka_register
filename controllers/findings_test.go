package controllers_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/keijoraamat/mka_register/controllers"
	"github.com/keijoraamat/mka_register/models"
	"github.com/stretchr/testify/assert"
)

func TestFindings_Index_Should_Contain_List_Of_Findig_Acts(t *testing.T) {
	engine := html.New("../views", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	// Setup routes
	faC := controllers.FindingActController{DB: db}
	app.Get("/leidmine", faC.Index)

	// Create request
	res := getResponse(app, t, "/leidmine")

	// Was the request success?
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body := getResponseBody(res, t)

	// Is the response body containing initialli added acts?
	assert.Contains(t, string(body), firstFindingAct.WDActNumber)
	assert.Contains(t, string(body), firstFindingAct.FinderIdNumber)
	assert.Contains(t, string(body), secondFindingAct.FinderIdNumber)
	assert.Contains(t, string(body), secondFindingAct.FinderIdNumber)

}

func TestFinding_Act_Should_Be_Inserted_Into_DB_When_Form_Is_Submitted(t *testing.T) {
	act := &models.FindingAct{FinderName: "Helka", WDActNumber: "666", RecieverName: "Mikk"}

	// Setup new Fiber app
	engine := html.New("../views", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	// Setup routes
	faC := controllers.FindingActController{DB: db}
	app.Get("/leidmine/lisa", faC.NewFinding)
	app.Post("/leidmine/lisa", faC.CreateFinding)

	// Setup request body
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("finderName", act.FinderName)
	writer.WriteField("WDActNumber", act.WDActNumber)
	writer.WriteField("RecieverName", act.RecieverName)
	writer.Close()

	// Create request
	req := httptest.NewRequest(http.MethodPost, "/leidmine/lisa", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := app.Test(req, -1)

	// Was the request succesful?
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, res.StatusCode)

	// Is the new act in DB?
	var a models.FindingAct
	result := db.Where("finder_name = ? AND wd_act_number = ? and reciever_name = ?",
			   act.FinderName, act.WDActNumber, act.RecieverName).First(&a)
	if result.Error != nil {
		t.Error("Could not find given act in DB")
	}

}

func getResponseBody(res *http.Response, t *testing.T) []byte {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	return body
}

func getResponse(app *fiber.App, t *testing.T, url string) *http.Response {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	res, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}
	return res
}
