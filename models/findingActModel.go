package models

import (
	"time"

	"github.com/keijoraamat/mka_register/helpers"
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
	TransferDate     time.Time
	WDActNumber      string
	Status           string
	Locations        []*Location `gorm:"many2many:finding_locations;"`
}

type FindingActView struct {
	ID               uint
	FinderName       string
	FinderIdNumber   string
	RecieverName     string
	FindingType      string
	FindersFee       bool
	ResiginOwnership bool
	RemainAnonymous  bool
	TransferLocation string
	TransferDate     string
	WDActNumber      string
	Status           string
	WeeksToEnd       float64
	EndTime          string
}

func (fa *FindingAct) DataToTemplate() (fav FindingActView) {
	fav.ID = fa.ID
	fav.FinderName = fa.FinderName
	fav.FinderIdNumber = fa.FinderIdNumber
	fav.RecieverName = fa.RecieverName
	fav.FindingType = fa.FindingType
	fav.FindersFee = fa.FindersFee
	fav.ResiginOwnership = fa.ResiginOwnership
	fav.RemainAnonymous = fa.RemainAnonymous
	fav.TransferLocation = fa.TransferLocation
	fav.WDActNumber = fa.WDActNumber
	fav.Status = fa.Status

	fav.TransferDate = fa.TransferDate.Format("02.01.2006")

	fav.WeeksToEnd = helpers.WeeksToEnd(fa.TransferDate)

	fav.EndTime = fa.TransferDate.AddDate(0, 6, 0).Format("02.01.2006")
	return
}
