package models

import (
	"time"

	"gorm.io/gorm"
)

type ArtefactLocation struct {
	gorm.Model
	ArtefactID uint `gorm:"primaryKey"`
	Artefact   Artefact
	LocationID uint `gorm:"primaryKey"`
	Location   Location
	From       time.Time
	Until      time.Time
}
