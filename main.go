package main

import (
	"sipekom-rest-api/config"
	"sipekom-rest-api/database"
	_ "sipekom-rest-api/docs"
	"sipekom-rest-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func main() {
	APP_PORT := config.Env(config.APP_PORT)

	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	app.Use(recover.New())

	database.ConnectDB()

	app.Get("/docs/*", swagger.HandlerDefault)
	router.SetupRouter(app)

	app.Listen(":" + APP_PORT)

}
