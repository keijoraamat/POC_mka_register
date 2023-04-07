package models

import (
	"time"

	"gorm.io/gorm"
)

type FindingLocation struct {
	gorm.Model
	FindingAct FindingAct
	Location   Location
	From       time.Time
	Until      time.Time
}
