package main

import (

	"github.com/gofiber/fiber"
)
	
func main() {
	app := fiber.New()

	app.Static("/", "./site", fiber.Static{
		Compress:   true,
		ByteRange:  true,
		Browse:     true,
		Index:      "index.html"
	})

	app.Listen(":5080")

}
