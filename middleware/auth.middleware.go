package middleware

import (
	"fmt"

	"github.com/KuroNeko6666/speed-control-backend.git/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Auth(ctx *fiber.Ctx) error {
	store := session.New()
	res, err := store.Get(ctx)

	if err != nil {
		return handler.UnAuthorized(ctx)
	}
	defer res.Save()

	fmt.Print(res.Keys())
	fmt.Print(res.Get("id"))
	fmt.Print(res.ID())

	return handler.Success(ctx, res.Keys())

	return ctx.Next()
}
