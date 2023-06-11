package router

import (
	"sipekom-rest-api/handler"
	"sipekom-rest-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)
	api.Post("/login", handler.Login)

	user := api.Group("/user")
	user.Use(middleware.Protect())
	user.Get("/", handler.GetAllUser)
	user.Get("/create", handler.CreateUser)
	user.Get("/get/:id", handler.GetUser)
	user.Delete("/delete/:id", handler.DeleteUser)
	user.Put("/update/:id", handler.UpdateUser)
}
