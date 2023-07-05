package database

import (
	"fmt"
	"sipekom-rest-api/config"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/static"
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
			Role:        static.RoleAdmin,
			Photo:       "photo/xyp9x.jpg",
			IsActivated: static.Activated,
		},
		{
			Username:    "masdisini",
			Password:    password,
			Role:        static.RoleAdmin,
			Photo:       "photo/masdisini.jpg",
			IsActivated: static.Activated,
		},
		{
			Username:    "FrHaN",
			Password:    password,
			Role:        static.RoleAdmin,
			Photo:       "photo/frhan.jpg",
			IsActivated: static.Activated,
		},
		{
			Username:    "ujang",
			Password:    password,
			Role:        static.RoleKonsulen,
			Photo:       "photo/ujang.jpg",
			IsActivated: static.Activated,
		},
		{
			Username:    "tarung",
			Password:    password,
			Role:        static.RoleKonsulen,
			Photo:       "photo/tarung.jpg",
			IsActivated: static.Activated,
		},
		{
			Username:    "s1mple",
			Password:    password,
			Role:        static.RoleMahasiswa,
			Photo:       "photo/s1mple.jpg",
			IsActivated: static.Activated,
		},
		{
			Username:    "zywoo",
			Password:    password,
			Role:        static.RoleMahasiswa,
			Photo:       "photo/zywoo.jpg",
			IsActivated: static.Activated,
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
			Str:         "str/s1mple.pdf",
			Sip:         "Sip/s1mple.pdf",
			Kompetensi:  static.KompetensiSenior,
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
			Str:         "str/zywoo.pdf",
			Sip:         "Sip/zywoo.pdf",
			Kompetensi:  static.KompetensiJunior,
			IDUser:      7,
		},
	}

	lokasi := []entity.Lokasi{
		{
			Lokasi: "THTKL",
			Uri:    "thtkl",
		},
		{
			Lokasi: "Kesehatan Mata",
			Uri:    "kesehatan_mata",
		},
		{
			Lokasi: "Kesehatan Anak",
			Uri:    "kesehatan_anak",
		},
		{
			Lokasi: "Dermatologi, Venerologi & Estetika (DVE)",
			Uri:    "dve",
		},
		{
			Lokasi: "Ilmu Penyakit Dalam",
			Uri:    "ilmu_penyakit_dalam",
		},
		{
			Lokasi: "Obgyn",
			Uri:    "obgyn",
		},
		{
			Lokasi: "Neurologi",
			Uri:    "neurologi",
		},
		{
			Lokasi: "Patologi Anatomi",
			Uri:    "patologi_anatomi",
		},
		{
			Lokasi: "Ilmu Bedah",
			Uri:    "ilmu_bedah",
		},
		{
			Lokasi: "Ilmu Penyakit Dalam II",
			Uri:    "ilmu_penyakit_dalam_II",
		},
		{
			Lokasi: "Anastesiologi",
			Uri:    "anastesiologi",
		},
		{
			Lokasi: "IGD",
			Uri:    "igd",
		},
		{
			Lokasi: "ICU",
			Uri:    "icu",
		},
		{
			Lokasi: "POLI",
			Uri:    "poli",
		},
	}

	eLogBook := []entity.ELogBook{
		{
			Title:         "Mengukur Suhu Matahari",
			Jumlah:        1,
			StartTime:     time.Now().Add(time.Hour * -1),
			EndTime:       time.Now().Add(time.Hour),
			Deskripsi:     "Menggunakan Cahaya Matahari",
			MedicalRecord: "Aman",
			IsAccepted:    static.AccOnReview,
			IDUser:        6,
			IDKonsulen:    4,
		},
		{
			Title:         "Mengukur Suhu Matahari",
			Jumlah:        1,
			StartTime:     time.Now().Add(time.Hour * -1),
			EndTime:       time.Now().Add(time.Hour),
			Deskripsi:     "Menggunakan Cahaya Matahari",
			MedicalRecord: "Aman",
			IsAccepted:    static.AccOnReview,
			IDUser:        6,
			IDKonsulen:    4,
		},
		{
			Title:         "Mengukur Suhu Matahari",
			Jumlah:        1,
			StartTime:     time.Now().Add(time.Hour * -1),
			EndTime:       time.Now().Add(time.Hour),
			Deskripsi:     "Menggunakan Cahaya Matahari",
			MedicalRecord: "Aman",
			IsAccepted:    static.AccOnReview,
			IDUser:        7,
			IDKonsulen:    5,
		},
		{
			Title:         "Mengukur Suhu Matahari",
			Jumlah:        1,
			StartTime:     time.Now().Add(time.Hour * -1),
			EndTime:       time.Now().Add(time.Hour),
			Deskripsi:     "Menggunakan Cahaya Matahari",
			MedicalRecord: "Aman",
			IsAccepted:    static.AccOnReview,
			IDUser:        7,
			IDKonsulen:    5,
		},
	}

	absensi := []entity.Absensi{
		{
			Absen:     time.Now().Add(time.Hour * 1),
			AbsenFlag: static.AbsenCheckIn,
			Lokasi:    "Kesehatan Mata",
			IDUser:    6,
		},
		{
			Absen:     time.Now().Add(time.Hour * 1),
			AbsenFlag: static.AbsenCheckOut,
			Lokasi:    "Kesehatan Mata",
			IDUser:    6,
		},
		{
			Absen:     time.Now().Add(time.Hour * 1),
			AbsenFlag: static.AbsenCheckIn,
			Lokasi:    "Kesehatan Anak",
			IDUser:    7,
		},
		{
			Absen:     time.Now().Add(time.Hour * 1),
			AbsenFlag: static.AbsenCheckOut,
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

	if err := DB.Create(&absensi).Error; err != nil {
		fmt.Println("[Info] Cannot Populate Absen or Data Already Exist")
	}

	fmt.Println("[Info] Succes Populate Table")

}
