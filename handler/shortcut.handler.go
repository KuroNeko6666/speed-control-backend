package handler

import (
	"net/http"

	"github.com/KuroNeko6666/speed-control-backend.git/model/response"
	"github.com/gofiber/fiber/v2"
)

func BadRequest(ctx *fiber.Ctx, err string) error {
	status := http.StatusBadRequest
	return ctx.Status(status).JSON(response.Base{
		Status: status,
		Data:   err,
	})
}

func InternalServer(ctx *fiber.Ctx, err string) error {
	status := http.StatusInternalServerError
	return ctx.Status(status).JSON(response.Base{
		Status: status,
		Data:   err,
	})
}

func NotFound(ctx *fiber.Ctx) error {
	status := http.StatusNotFound
	return ctx.Status(status).JSON(response.Base{
		Status: status,
		Data:   http.StatusText(status),
	})
}

func UnAuthorized(ctx *fiber.Ctx) error {
	status := http.StatusUnauthorized
	return ctx.Status(status).JSON(response.Base{
		Status: status,
		Data:   http.StatusText(status),
	})
}

func SuccessString(ctx *fiber.Ctx) error {
	status := http.StatusOK
	return ctx.Status(status).JSON(response.Base{
		Status: status,
		Data:   http.StatusText(status),
	})
}

func Success(ctx *fiber.Ctx, data interface{}) error {
	status := http.StatusOK
	return ctx.Status(status).JSON(response.Base{
		Status: status,
		Data:   data,
	})
}
