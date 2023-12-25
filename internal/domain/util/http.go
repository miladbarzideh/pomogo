package util

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetNumericIdentifier(ctx *fiber.Ctx, key string) (int, error) {
	id, err := strconv.Atoi(ctx.Params(key))
	if err != nil {
		return 0, NewBadRequest(err)
	}
	return id, nil
}
