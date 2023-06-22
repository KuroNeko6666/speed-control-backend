package app

import (
	"log"

	"github.com/KuroNeko6666/speed-control-backend.git/config"
	"github.com/KuroNeko6666/speed-control-backend.git/database"
	"github.com/KuroNeko6666/speed-control-backend.git/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RunApp() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowMethods:     "*",
		AllowCredentials: true,
	}))
	database.ConnectDB()

	router.Auth(app)
	router.Device(app)
	router.DeviceData(app)
	router.User(app)
	router.Dashboard(app)

	log.Fatal(app.Listen(config.ServerAddress()))
}
