package main

import (
	fiber "github.com/gofiber/fiber/v2"

	"obfuscation-challenge-server/routes"
)

func main() {

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "Content-Type,User-Agent")
		return c.Next()
	})

	app.Static("/static", "../client/static")
	app.Static("/", "../client/index.html")

	routes.InitRoutes(app)

	app.Listen(":1337")
}
