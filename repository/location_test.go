package repository_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/keijoraamat/mka_register/models"
	"github.com/keijoraamat/mka_register/repository"
)

func Test_AddLocation_Returns_ID_Given_Location_to_Save(t *testing.T) {
	l := models.Location{
		County:            "Almepõis",
		Parish:            "",
		Village:           "Liiva",
		CadastralUnitCode: "44:67",
		CadastralUnitName: "Kiigemäe",
		FindingsAmount:    2,
	}

	loc, err := repository.AddLocation(l, db)
	if err != nil {
		log.Println("Error adding loc in test ", err)
	}

	if loc.ID == l.ID {
		t.Errorf("AddLocation() returned ID %d, expected a bit more", l.ID)
	}

}

func Test_GetAllLocations_Should_Return_Slice_Of_Locations(t *testing.T) {

	actual, _ := repository.GetAllLocations(db)
	var expt []models.Location

	if !reflect.DeepEqual(reflect.TypeOf(actual), reflect.TypeOf(expt)) {
		t.Errorf("GetAllLocations() returned type %T, expected type %T", actual, expt)
	}

}

func Test_GetAllLocations_Should_Return_Slice_Of_Locations_Containg_Given_Location(t *testing.T) {
	l := models.Location{
		County:            "Vagula",
		Parish:            "",
		Village:           "Liiva",
		CadastralUnitCode: "44:67",
		CadastralUnitName: "gemäe",
		FindingsAmount:    2,
	}
	loc, _ := repository.AddLocation(l, db)
	actual, _ := repository.GetAllLocations(db)

	var isInSlice = false
	for _, location := range actual {
		if location.ID == loc.ID {
			isInSlice = true
			return
		}
	}

	if !isInSlice{
		t.Errorf("GetAllLocations() returned slice without inserted location %d", loc.ID)
	}

}
