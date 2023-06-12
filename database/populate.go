package database

import (
	"fmt"
	"sipekom-rest-api/config"
	"sipekom-rest-api/model"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/utils"
	"strconv"
	"time"
)

func PopulateTable() {
	populate, err := strconv.ParseBool(config.Env("DB_POPULATE"))
	if err != nil {
		fmt.Println("[Error] Failed to Parse DB_POPULATE on .env")
		return
	}

	if !populate {
		fmt.Println("[Info] Table Will Not Populate")
		return
	}

	password, _ := utils.HashPassword("password")

	users := []entity.User{
		{
			Username:    "xyp9x",
			Password:    password,
			Level:       model.LevelAdmin,
			IsActivated: model.Activated,
		},
		{
			Username:    "masdisini",
			Password:    password,
			Level:       model.LevelAdmin,
			IsActivated: model.Activated,
		},
		{
			Username:    "FrHaN",
			Password:    password,
			Level:       model.LevelAdmin,
			IsActivated: model.Activated,
		},
		{
			Username:    "ujang",
			Password:    password,
			Level:       model.LevelKonsulen,
			IsActivated: model.Activated,
		},
		{
			Username:    "tarung",
			Password:    password,
			Level:       model.LevelKonsulen,
			IsActivated: model.Activated,
		},
		{
			Username:    "s1mple",
			Password:    password,
			Level:       model.LevelMahasiswa,
			IsActivated: model.Activated,
		},
		{
			Username:    "zywoo",
			Password:    password,
			Level:       model.LevelMahasiswa,
			IsActivated: model.Activated,
		},
	}

	konsulen := []entity.Konsulen{
		{
			Name:      "Ujang Sudrajat",
			Spesialis: "Kriptografi",
			IDUser:    4,
		},
		{
			Name:      "Tarung Flanker",
			Spesialis: "Networking",
			IDUser:    5,
		},
	}

	mahasiswa := []entity.PPDS{
		{
			Name:        "Alexander",
			BirthDate:   "1 Dec 2001",
			BirthPlace:  "Kyiv",
			NIK:         "31000000",
			NIM:         "2100000",
			Address:     "Kyiv Sekitar Bakhmut",
			PhoneNumber: "08318010",
			Angkatan:    2016,
			Prodi:       "Teknik Mesin",
			Photo:       "photo/s1mple.jpg",
			Str:         "str/s1mple.pdf",
			Sip:         "Sip/s1mple.pdf",
			Kompetensi:  model.KompetensiSenior,
			IDUser:      6,
		},
		{
			Name:        "Mathieu",
			BirthDate:   "1 Dec 2001",
			BirthPlace:  "Paris",
			NIK:         "31000000",
			NIM:         "2100000",
			Address:     "Paris Sekitar Sedan",
			PhoneNumber: "08318010",
			Angkatan:    2018,
			Prodi:       "Teknik Sipil",
			Photo:       "photo/zywoo.jpg",
			Str:         "str/zywoo.pdf",
			Sip:         "Sip/zywoo.pdf",
			Kompetensi:  model.KompetensiJunior,
			IDUser:      7,
		},
	}

	lokasi := []entity.Lokasi{
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

	eLogBook := []entity.ELogBook{
		{
			Name:           "Mengukur Suhu Matahari",
			Jumlah:         1,
			StartTime:      time.Now().Add(time.Hour * -1),
			EndTime:        time.Now().Add(time.Hour),
			Deskripsi:      "Menggunakan Cahaya Matahari",
			Medical_Record: "Aman",
			IDKonsulen:     1,
			IDUser:         6,
		},
		{
			Name:           "Mengukur Suhu Matahari",
			Jumlah:         1,
			StartTime:      time.Now().Add(time.Hour * -1),
			EndTime:        time.Now().Add(time.Hour),
			Deskripsi:      "Menggunakan Cahaya Matahari",
			Medical_Record: "Aman",
			IDKonsulen:     1,
			IDUser:         6,
		},
		{
			Name:           "Mengukur Suhu Matahari",
			Jumlah:         1,
			StartTime:      time.Now().Add(time.Hour * -1),
			EndTime:        time.Now().Add(time.Hour),
			Deskripsi:      "Menggunakan Cahaya Matahari",
			Medical_Record: "Aman",
			IDKonsulen:     1,
			IDUser:         7,
		},
		{
			Name:           "Mengukur Suhu Matahari",
			Jumlah:         1,
			StartTime:      time.Now().Add(time.Hour * -1),
			EndTime:        time.Now().Add(time.Hour),
			Deskripsi:      "Menggunakan Cahaya Matahari",
			Medical_Record: "Aman",
			IDKonsulen:     1,
			IDUser:         7,
		},
	}

	konsul := []entity.Konsul{
		{
			IDMahasiswa: 6,
			IDKonsulen:  4,
		},
		{
			IDMahasiswa: 7,
			IDKonsulen:  5,
		},
	}

	absensi := []entity.Absensi{
		{
			Absen:     time.Now().Add(time.Hour * 1),
			AbsenFlag: model.FlagAbsenCheckIn,
			Lokasi:    "Kesehatan Mata",
			IDUser:    6,
		},
		{
			Absen:     time.Now().Add(time.Hour * 1),
			AbsenFlag: model.FlagAbsenCheckOut,
			Lokasi:    "Kesehatan Mata",
			IDUser:    6,
		},
		{
			Absen:     time.Now().Add(time.Hour * 1),
			AbsenFlag: model.FlagAbsenCheckIn,
			Lokasi:    "Kesehatan Anak",
			IDUser:    7,
		},
		{
			Absen:     time.Now().Add(time.Hour * 1),
			AbsenFlag: model.FlagAbsenCheckOut,
			Lokasi:    "Kesehatan Anak",
			IDUser:    7,
		},
	}

	if err := DB.Create(&users).Error; err != nil {
		fmt.Println("[Info] Cannot Populate User or Data Already Exist")
	}

	if err := DB.Create(&lokasi).Error; err != nil {
		fmt.Println("[Info] Cannot Populate Lokasi or Data Already Exist")
	}

	if err := DB.Create(&konsulen).Error; err != nil {
		fmt.Println("[Info] Cannot Populate Konsulen or Data Already Exist")
	}

	if err := DB.Create(&mahasiswa).Error; err != nil {
		fmt.Println("[Info] Cannot Populate Mahasiwa or Data Already Exist")
	}

	if err := DB.Create(&eLogBook).Error; err != nil {
		fmt.Println("[Info] Cannot Populate E-Log Book or Data Already Exist")
	}

	if err := DB.Create(&konsul).Error; err != nil {
		fmt.Println("[Info] Cannot Populate Konsul or Data Already Exist")
	}

	if err := DB.Create(&absensi).Error; err != nil {
		fmt.Println("[Info] Cannot Populate Absen or Data Already Exist")
	}

	fmt.Println("[Info] Succes Populate Table")

}
