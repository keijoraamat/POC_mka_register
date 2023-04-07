package models

import (
	"gorm.io/gorm"
)

type Act struct {
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
}
