package entity

import (
	"gorm.io/gorm"
)

type PPDS struct {
	gorm.Model
	Name        string `json:"name"`
	BirthDate   string `json:"birthdate"`
	BirthPlace  string `json:"birthplace"`
	NIK         string `json:"nik"`
	NIM         string `json:"nim"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Angkatan    uint   `json:"angkatan"`
	Prodi       string `json:"prodi"`
	Str         string `json:"str"`
	Sip         string `json:"sip"`
	Kompetensi  uint8  `json:"kompetensi"`
	IDUser      uint   `gorm:"not null;unique" json:"id_user"`
}
