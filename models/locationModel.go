package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	County            string
	Parish            string
	Village           string
	CadastralUnitCode string
	CadastralUnitName string
	FindingsAmount    uint
	FindingActs       []*FindingAct `gorm:"many2many:finding_locations;"`
	Artefacts         []*Artefact   `gorm:"many2many:artefact_locations;"`
	Afacts            []Artefact    `gorm:"-"`
}

func (*Location) IsLocationInputOK() bool {

	//TODO: validate all needed fields
	return true
}
