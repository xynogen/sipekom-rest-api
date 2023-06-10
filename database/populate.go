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

	lokasi := []model.Lokasi{
		{
			Lokasi: "THTKL",
			Uri:    "absen/thtkl",
		},
		{
			Lokasi: "Kesehatan Mata",
			Uri:    "absen/kesehatan_mata",
		},
		{
			Lokasi: "Kesehatan Anak",
			Uri:    "absen/kesehatan_anak",
		},
		{
			Lokasi: "Dermatologi, Venerologi & Estetika (DVE)",
			Uri:    "absen/dve",
		},
		{
			Lokasi: "Ilmu Penyakit Dalam",
			Uri:    "absen/ilmu_penyakit_dalam",
		},
		{
			Lokasi: "Obgyn",
			Uri:    "absen/obgyn",
		},
		{
			Lokasi: "Neurologi",
			Uri:    "absen/neurologi",
		},
		{
			Lokasi: "Patologi Anatomi",
			Uri:    "absen/patologi_anatomi",
		},
		{
			Lokasi: "Ilmu Bedah",
			Uri:    "absen/ilmu_bedah",
		},
		{
			Lokasi: "Ilmu Penyakit Dalam II",
			Uri:    "absen/ilmu_penyakit_dalam_II",
		},
		{
			Lokasi: "Anastesiologi",
			Uri:    "absen/anastesiologi",
		},
		{
			Lokasi: "IGD",
			Uri:    "absen/igd",
		},
		{
			Lokasi: "ICU",
			Uri:    "absen/icu",
		},
		{
			Lokasi: "POLI",
			Uri:    "absen/poli",
		},
	}

	if userQuery := DB.Create(&users); userQuery.Error != nil {
		fmt.Println("[Info] Cannot Populate User or Data Already Exist")
	}

	if userQuery := DB.Create(&lokasi); userQuery.Error != nil {
		fmt.Println("[Info] Cannot Populate Lokasi or Data Already Exist")
	}

	fmt.Println("[Info] Succes Populate Table")

}
