package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	County        string
	Farmsted      string
	CadastralUnit string
}
