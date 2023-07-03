package database

import (
	"fmt"

	"sipekom-rest-api/config"
	"sipekom-rest-api/model/entity"

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
		&entity.User{},
		&entity.PPDS{},
		&entity.Konsulen{},
		&entity.Lokasi{},
		&entity.ELogBook{},
		&entity.Absensi{},
	)

}
