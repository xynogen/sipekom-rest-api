package entity

import (
	"gorm.io/gorm"
)

type Konsul struct {
	IDMahasiswa uint `gorm:"not null" json:"id_mahasiswa"`
	IDKonsulen  uint `gorm:"not null" json:"id_konsulen"`
}

type Konsulen struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	Spesialis string `gorm:"not null" json:"spesialis"`
	IDUser    uint   `gorm:"not null" json:"id_user"`
}
