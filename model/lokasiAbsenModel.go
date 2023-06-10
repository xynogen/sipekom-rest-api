package model

import "gorm.io/gorm"

type Prodi struct {
	gorm.Model
	Prodi string `gorm:"not null" json:"prodi"`
	Uri   string `gorm:"not null; unique" json:"uri"`
}

type Ruangan struct {
	gorm.Model
	Ruangan string `gorm:"not null" json:"ruangan"`
	Uri     string `gorm:"not null; unique" json:"uri"`
}
