package models

import "gorm.io/gorm"

type Artefact struct {
	gorm.Model
	Name      string
	YCoord    string
	XCoord    string
	Amount    uint
	State     string
	Locations []*Location `gorm:"many2many:artefact_locations;"`
}
