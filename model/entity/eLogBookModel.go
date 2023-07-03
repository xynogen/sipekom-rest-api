package entity

import (
	"time"

	"gorm.io/gorm"
)

type ELogBook struct {
	gorm.Model
	Title         string    `gorm:"not null" json:"title"`
	Jumlah        uint      `gorm:"not null" json:"jumlah"`
	StartTime     time.Time `gorm:"not null" json:"start_time"`
	EndTime       time.Time `gorm:"not null" json:"end_time"`
	Deskripsi     string    `json:"deskripsi"`
	MedicalRecord string    `gorm:"not null" json:"medical_record"`
	IsAccepted    uint8     `gorm:"not null" json:"is_accepted"`
	IDUser        uint      `gorm:"not null" json:"id_user"`
	IDKonsulen    uint      `gorm:"not null" json:"id_konsulen"`
}
