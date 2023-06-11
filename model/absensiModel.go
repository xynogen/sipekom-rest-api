package model

import (
	"time"

	"gorm.io/gorm"
)

type Absensi struct {
	gorm.Model
	Absen     time.Time `gorm:"not null" json:"absen"`
	AbsenFlag uint8     `gorm:"not null" json:"absen_flag"`
	Lokasi    string    `gorm:"not null" json:"lokasi"`
	IDUser    uint      `gorm:"not null" json:"id_user"`
}
