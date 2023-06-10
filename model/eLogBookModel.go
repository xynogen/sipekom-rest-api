package model

import (
	"time"

	"gorm.io/gorm"
)

type ELogBook struct {
	gorm.Model
	Name           string    `gorm:"not null" json:"name"`
	Jumlah         uint      `gorm:"not null" json:"jumlah"`
	StartTime      time.Time `gorm:"not null" json:"start_time"`
	EndTime        time.Time `gorm:"not null" json:"end_time"`
	Deskripsi      string    `json:"deskripsi"`
	Medical_Record string    `gorm:"not null" json:"medical_record"`
	KonsulenID     uint      `gorm:"not null" json:"id_konsulen"`
}
