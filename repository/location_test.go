package repository_test

import (
	"reflect"
	"testing"

	"github.com/keijoraamat/mka_register/models"
	"github.com/keijoraamat/mka_register/repository"
)

func TestGetLocationsByFindingActID(t *testing.T) {
	type args struct {
		actID string
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Location
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repository.GetLocationsByFindingActID(tt.args.actID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocationsByFindingActID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocationsByFindingActID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddLocation(t *testing.T) {
	type args struct {
		l models.Location
	}
	tests := []struct {
		name    string
		args    args
		want    models.Location
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repository.AddLocation(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
