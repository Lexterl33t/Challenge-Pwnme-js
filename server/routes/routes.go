package routes

import (
	"github.com/gofiber/fiber/v2"

	"obfuscation-challenge-server/controller"
)

func InitRoutes(app *fiber.App) {
	app.Post("/welcome", controller.Welcome)
}
