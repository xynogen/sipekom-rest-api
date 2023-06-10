package model

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
	Angkatan    int    `json:"angkatan"`
	Prodi       string `json:"prodi"`
	Photo       string `json:"photo"`
	Str         string `json:"str"`
	Sip         string `json:"sip"`
	Kompetensi  string `json:"kompetensi"`
	IDUser      uint   `json:"id_user"`
}
