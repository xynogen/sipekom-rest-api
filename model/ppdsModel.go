package model

import (
	"time"

	"gorm.io/gorm"
)

type PPDS struct {
	gorm.Model
	Name        string    `json:"name"`
	BirthDate   time.Time `json:"birthdate"`
	BirthPlace  string    `json:"birthplace"`
	NIK         string    `json:"nik"`
	NIM         string    `json:"nim"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phonenumber"`
	Angkatan    int       `json:"2019"`
	Prodi       string    `json:"prodi"`
	Photo       string    `json:"photo"`
	Str         string    `json:"str"`
	Sip         string    `json:"sip"`
	IDUser      uint      `json:"iduser"`
}
