package model

import (
	"gorm.io/gorm"
)

type Konsulen struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	Spesialis string `gorm:"not null" json:"spesialis"`
	IDUser    uint   `gorm:"not null" json:"iduser"`
}
