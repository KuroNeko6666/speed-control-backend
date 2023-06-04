package router

import (
	"net/http"

	"github.com/KuroNeko6666/speed-control-backend.git/config"
	"github.com/KuroNeko6666/speed-control-backend.git/handler"
	"github.com/KuroNeko6666/speed-control-backend.git/model/response"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v4"
)

func Auth(app *fiber.App) {
	app.Get("/auth", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.SecretKeyApp)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(response.Base{
				Status: http.StatusUnauthorized,
				Data:   err.Error(),
			})
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			return handler.SuccessString(ctx)
		},
	}))
	app.Post("/auth", handler.Login)
	app.Delete("/auth", handler.Logout)
	app.Post("/register", handler.Register)
}
