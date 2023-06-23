package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"not null;unique" json:"username"`
	Password    string `gorm:"not null" json:"password"`
	Role        uint8  `gorm:"not null" json:"role"`
	Photo       string `json:"photo"`
	IsActivated uint8  `gorm:"not null" json:"is_activated"`
}
