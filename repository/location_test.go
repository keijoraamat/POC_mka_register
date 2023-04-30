package repository_test

import (
	"reflect"
	"testing"

	"github.com/keijoraamat/mka_register/models"
	"github.com/keijoraamat/mka_register/repository"
)

func Test_GetAllLocations_Should_Return_Slice_Of_Locations(t *testing.T) {

	actual, _ := repository.GetAllLocations(db)
	var expt []models.Location

	if !reflect.DeepEqual(reflect.TypeOf(actual), reflect.TypeOf(expt)) {
		t.Errorf("GetAllLocations() returned type %T, expected type %T", actual, expt)
	}

}
