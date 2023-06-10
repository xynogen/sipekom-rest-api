package config

import (
	"os"

	"github.com/joho/godotenv"
)

var APP_PORT = "APP_PORT"
var SECRET = "SECRET"
var DB_POPULATE = "DB_POPULATE"
var DB_USER = "DB_USER"
var DB_PASSWORD = "DB_PASSWORd"
var DB_HOST = "DB_HOST"
var DB_PORT = "DB_PORT"
var DB_DATABASE = "DB_DATABASE"

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("[Error] Cannot Load .env File")
	}
	return os.Getenv(key)
}
