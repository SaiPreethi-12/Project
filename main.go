package main

import (
	"project/config"
	"project/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	config.ConnectRedis()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Setup(app)

	app.Listen(":3000")
}
