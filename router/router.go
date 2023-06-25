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
	absen.Get("/create/:uri_base64", handler.CreateAbsen)
	absen.Put("/update/:id", handler.UpdateAbsen)
	absen.Delete("/delete/:id", handler.DeleteAbsen)

	elogbook := api.Group("/elogbook")
	elogbook.Use(middleware.Protect())
	elogbook.Get("/", handler.GetAllELogBook)
	elogbook.Get("/get/:id", handler.GetELogBook)
	elogbook.Post("/create", handler.CreateELogBook)
	elogbook.Put("/update/:id", handler.UpdateElogBook)
	elogbook.Delete("/delete/:id", handler.DeleteELogBook)

	konsulen := api.Group("/konsulen")
	konsulen.Use(middleware.Protect())
	konsulen.Get("/", handler.GetAllKonsulen)
	konsulen.Get("/get/:id", handler.GetKonsulen)
	konsulen.Post("/create", handler.CreateKonsulen)
	konsulen.Put("/update/:id", handler.UpdateKonsulen)

	check := api.Group("/check")
	check.Use(middleware.Protect())
	check.Get("/", handler.Check)

	qr := api.Group("/qr")
	qr.Get("/get/:id_lokasi", handler.GetQR)

}
