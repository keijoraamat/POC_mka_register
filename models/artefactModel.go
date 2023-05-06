package models

import "gorm.io/gorm"

type Artefact struct {
	gorm.Model
	Name        string
	YCoord      string
	XCoord      string
	Amount      uint
	FindingActs []*FindingAct `gorm:"many2many:artefact_locations;"`
}
