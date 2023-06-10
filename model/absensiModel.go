package model

import (
	"time"

	"gorm.io/gorm"
)

type Absensi struct {
	gorm.Model
	CheckIn  time.Time `gorm:"not null" json:"checkin"`
	CheckOut time.Time `gorm:"not null" json:"checkout"`
	NIM      string    `gorm:"not null" json:"nim"`
	Name     string    `gorm:"not null" json:"name"`
	Lokasi   string    `gorm:"not null" json:"lokasi"`
}
