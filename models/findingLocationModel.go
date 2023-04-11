package models

import (
	"time"

	"gorm.io/gorm"
)

type FindingLocation struct {
	gorm.Model
	FindingActID uint `gorm:"primaryKey"`
	FindingAct   FindingAct
	LocationID   uint `gorm:"primaryKey"`
	Location     Location
	From         time.Time
	Until        time.Time
}
