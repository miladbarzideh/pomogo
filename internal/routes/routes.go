package routes

import "github.com/gofiber/fiber/v2"

func Initialize(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Response!")
	})
}
