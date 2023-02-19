package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Home(ctx *fiber.Ctx) error {

	return ctx.Render("index", nil)
}
