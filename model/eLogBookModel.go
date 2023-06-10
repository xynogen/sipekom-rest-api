package model

import (
	"time"

	"gorm.io/gorm"
)

type ELogBook struct {
	gorm.Model
	Name           string    `gorm:"not null" json:"name"`
	Jumlah         uint      `gorm:"not null" son:"jumlah"`
	Waktu          time.Time `gorm:"not null" json:"waktu"`
	Deskripsi      string    `json:"deskripsi"`
	Supervisor     string    `gorm:"not null" json:"suoervisor"`
	Medical_Record string    `gorm:"not null" json:"medical_record"`
}
