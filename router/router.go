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
	user.Get("/get/:id", handler.GetUser)
	user.Post("/create", handler.CreateUser)
	user.Delete("/delete/:id", handler.DeleteUser)
	user.Put("/update/:id", handler.UpdateUser)

	absen := api.Group("/absen")
	absen.Use(middleware.Protect())
	absen.Get("/", handler.GetAllAbsen)
	absen.Get("/get/:id", handler.GetAbsen)
	absen.Get("/create/:lokasi", handler.CreateAbsen)
	absen.Put("/update/:id", handler.UpdateAbsen)
	absen.Delete("/delete/:id", handler.DeleteAbsen)

	elogbook := api.Group("/elogbook")
	elogbook.Use(middleware.Protect())
	elogbook.Get("/", handler.GetAllELogBook)
	elogbook.Get("/get/:id", handler.GetELogBook)

	konsulen := api.Group("/konsulen")
	konsulen.Use(middleware.Protect())
	konsulen.Get("/", handler.GetAllKonsulen)
	konsulen.Get("/get/:id", handler.GetKonsulen)

}
