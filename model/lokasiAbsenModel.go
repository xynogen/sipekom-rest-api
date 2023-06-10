package model

import "gorm.io/gorm"

type Lokasi struct {
	gorm.Model
	Lokasi string `gorm:"not null" json:"lokasi"`
	Uri    string `gorm:"not null; unique" json:"uri"`
}
