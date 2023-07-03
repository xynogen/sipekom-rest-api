package router

import (
	"time"

	"sipekom-rest-api/handler"
	"sipekom-rest-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api")
	api.Use(logger.New())
	api.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Second,
	}))

	api.Get("/", handler.Hello)
	api.Post("/login", handler.Login)
	api.Get("/data/:search_query", handler.GetData)

	user := api.Group("/user")
	user.Use(middleware.Protect())
	user.Get("/", handler.GetAllUser)
	user.Get("/get/:id_user", handler.GetUser)
	user.Get("/data/:id_user", handler.GetUserData)
	user.Post("/create", handler.CreateUser)
	user.Delete("/delete/:id_user", handler.DeleteUser)
	user.Put("/update/:id_user", handler.UpdateUser)

	absen := api.Group("/absen")
	absen.Use(middleware.Protect())
	absen.Get("/:id_user", handler.GetAllAbsen)
	absen.Get("/get/:id_absen", handler.GetAbsen)
	absen.Get("/create/:uri_base64", handler.CreateAbsen)
	absen.Put("/update/:id_absen", handler.UpdateAbsen)
	absen.Delete("/delete/:id_absen", handler.DeleteAbsen)

	elogbook := api.Group("/elogbook")
	elogbook.Use(middleware.Protect())
	elogbook.Get("/", handler.GetAllELogBook)
	elogbook.Get("/get/:id", handler.GetELogBook)
	elogbook.Post("/create", handler.CreateELogBook)
	elogbook.Put("/update/:id_elogbook", handler.UpdateElogBook)
	elogbook.Put("/accepted/:id_elogbook", handler.AcceptedElogBook)
	elogbook.Put("/rejected/:id_elogbook", handler.RejectedElogBook)
	elogbook.Delete("/delete/:id", handler.DeleteELogBook)

	konsulen := api.Group("/konsulen")
	konsulen.Use(middleware.Protect())
	konsulen.Get("/", handler.GetAllKonsulen)
	konsulen.Get("/get/:id", handler.GetKonsulen)
	konsulen.Post("/create", handler.CreateKonsulen)
	konsulen.Put("/update/:id", handler.UpdateKonsulen)

	lokasi := api.Group("/lokasi")
	lokasi.Use(middleware.Protect())
	lokasi.Get("/", handler.GetAllLokasi)

	whoami := api.Group("/whoami")
	whoami.Use(middleware.Protect())
	whoami.Get("/", handler.Whoami)

	qr := api.Group("/qr")
	qr.Use(middleware.Protect())
	qr.Get("/get/:id_lokasi", handler.GetQR)

}
