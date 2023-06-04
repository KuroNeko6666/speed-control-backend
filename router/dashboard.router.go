package router

import (
	"net/http"

	"github.com/KuroNeko6666/speed-control-backend.git/config"
	"github.com/KuroNeko6666/speed-control-backend.git/handler"
	"github.com/KuroNeko6666/speed-control-backend.git/model/response"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v4"
)

func Dashboard(app *fiber.App) {
	group := app.Group("/dashboard", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.SecretKeyApp)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(response.Base{
				Status: http.StatusUnauthorized,
				Data:   err.Error(),
			})
		},
	}))
	group.Get("/user", handler.UserDashboard)
	group.Get("/device", handler.DeviceDashboard)
	group.Get("/data", handler.DeviceDataDashboard)
	group.Get("/count/user", handler.UserCount)
	group.Get("/count/device", handler.DeviceCount)
	group.Get("/count/data", handler.DataCount)
	group.Get("/count/data-user", handler.DataCountUserDevice)
}
