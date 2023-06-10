package database

import (
	"fmt"
	"strconv"

	"sipekom-rest-api/config"
	"sipekom-rest-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Env(config.DB_USER),
		config.Env(config.DB_PASSWORD),
		config.Env(config.DB_HOST),
		config.Env(config.DB_PORT),
		config.Env(config.DB_DATABASE),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("[Error] Failed to Connect to Database")
	}

	fmt.Println("[Info] Connected to Database")

	DB.AutoMigrate(
		&model.User{},
		&model.PPDS{},
		&model.Konsul{},
		&model.Konsulen{},
		&model.Lokasi{},
		&model.ELogBook{},
		&model.Absensi{},
	)

	populate, err := strconv.ParseBool(config.Env("DB_POPULATE"))
	if err != nil {
		fmt.Println("[Error] Failed to Parse DB_POPULATE on .env")
	}

	if populate {
		PopulateTable()
	} else {
		fmt.Println("[Info] Table Will Not Populate")
	}

}
