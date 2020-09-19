package main

import (

	"github.com/gofiber/fiber"
)
	
func main() {
	app := fiber.New()

	app.Static("/", "./site")

	app.Static("/prefix", "./site")

	app.Static("*", "./site/index.html")

	app.Listen(":5080")

}