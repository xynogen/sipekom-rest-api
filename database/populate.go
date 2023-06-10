package database

import (
	"fmt"
	"sipekom-rest-api/model"
	"sipekom-rest-api/utils"
)

func PopulateTable() {
	password, _ := utils.HashPassword("password")

	users := []model.User{
		{
			Username: "xyp9x",
			Password: password,
			Level:    model.Level1,
		},
		{
			Username: "masdisini",
			Password: password,
			Level:    model.Level1,
		},
		{
			Username: "FrHaN",
			Password: password,
			Level:    model.Level1,
		},
	}

	prodi := []model.Prodi{
		{
			Prodi: "THTKL",
			Uri:   "absen/thtkl",
		},
		{
			Prodi: "Kesehatan Mata",
			Uri:   "absen/kesehatan_mata",
		},
		{
			Prodi: "Kesehatan Anak",
			Uri:   "absen/kesehatan_anak",
		},
		{
			Prodi: "Dermatologi, Venerologi & Estetika (DVE)",
			Uri:   "absen/dve",
		},
		{
			Prodi: "Ilmu Penyakit Dalam",
			Uri:   "absen/ilmu_penyakit_dalam",
		},
		{
			Prodi: "Obgyn",
			Uri:   "absen/obgyn",
		},
		{
			Prodi: "Neurologi",
			Uri:   "absen/neurologi",
		},
		{
			Prodi: "Patologi Anatomi",
			Uri:   "absen/patologi_anatomi",
		},
		{
			Prodi: "Ilmu Bedah",
			Uri:   "absen/ilmu_bedah",
		},
		{
			Prodi: "Ilmu Penyakit Dalam II",
			Uri:   "absen/ilmu_penyakit_dalam_II",
		},
		{
			Prodi: "Anastesiologi",
			Uri:   "absen/anastesiologi",
		},
	}

	ruangan := []model.Ruangan{
		{
			Ruangan: "IGD",
			Uri:     "absen/igd",
		},
		{
			Ruangan: "ICU",
			Uri:     "absen/icu",
		},
		{
			Ruangan: "POLI",
			Uri:     "absen/poli",
		},
	}

	if userQuery := DB.Create(&users); userQuery.Error != nil {
		fmt.Println("[Info] Cannot Populate User or Data Already Exist")
	}

	if userQuery := DB.Create(&prodi); userQuery.Error != nil {
		fmt.Println("[Info] Cannot Populate Prodi or Data Already Exist")
	}

	if userQuery := DB.Create(&ruangan); userQuery.Error != nil {
		fmt.Println("[Info] Cannot Populate Ruangan or Data Already Exist")
	}

	fmt.Println("[Info] Succes Populate Table")

}
