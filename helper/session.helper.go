package helper

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func GetStore(ctx *fiber.Ctx) (*session.Session, error) {
	store := session.New()
	res, err := store.Get(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SetSession(ctx *fiber.Ctx, key string) error {
	sess, err := GetStore(ctx)
	if err != nil {
		return err
	}
	sess.Set("user_id", key)
	sess.SetExpiry(time.Hour * 24)
	if err := sess.Save(); err != nil {
		return err
	}
	return nil
}

func GetSession(ctx *fiber.Ctx) (string, error) {
	sess, err := GetStore(ctx)
	if err != nil {
		return "", err
	}

	userID := sess.Get("user_id")

	if userID == nil {
		return "", errors.New(http.StatusText(http.StatusUnauthorized))
	}

	return fmt.Sprintf("%v", userID), nil

}

func DeleteSession(ctx *fiber.Ctx) error {
	sess, err := GetStore(ctx)
	if err != nil {
		return err
	}

	err = sess.Destroy()

	if err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return nil

}
