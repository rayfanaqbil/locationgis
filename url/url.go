package url

import (
	"github.com/rayfanaqbil/locationgis/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)
	page.Group("/api")
	page.Post("/findnearestroad", controller.FindNearestRoad)

	// Menambahkan Swagger UI menggunakan manual URL dan konfigurasi
	page.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json", // Ganti dengan path yang sesuai dengan hasil 'swag init'
	}))
}
