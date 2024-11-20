package main

import (
	"log"

	"github.com/rayfanaqbil/locationgis/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/rayfanaqbil/locationgis/url"
	_ "github.com/rayfanaqbil/locationgis/docs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
