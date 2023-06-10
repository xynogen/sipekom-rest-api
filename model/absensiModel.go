package model

import (
	"time"

	"gorm.io/gorm"
)

type Absensi struct {
	gorm.Model
	CheckIn  time.Time `gorm:"not null" json:"checkin"`
	CheckOut time.Time `gorm:"not null" json:"checkout"`
	Lokasi   string    `gorm:"not null" json:"lokasi"`
	IDUser   uint      `gorm:"not null" json:"id_user"`
}
