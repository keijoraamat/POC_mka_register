package models

import (
	"gorm.io/gorm"
)

type FindingAct struct {
	gorm.Model
	FinderName       string
	FinderIdNumber   string
	RecieverName     string
	FindingType      string
	FindersFee       bool
	ResiginOwnership bool
	RemainAnonymous  bool
	TransferLocation string
	WDActNumber      string
	Status           string
	Locations        []*Location `gorm:"many2many:finding_locations;"`
}
