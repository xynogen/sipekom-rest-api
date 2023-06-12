package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"not null;unique" json:"username"`
	Password    string `gorm:"not null" json:"password"`
	Level       uint8  `gorm:"not null" json:"level"`
	IsActivated uint8  `gorm:"not null" json:"is_activated"`
}
