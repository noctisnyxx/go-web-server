package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("first web server trial")
	})
	app.Get("/myBookList", handler.getBookList)
	log.Fatal(app.Listen(":8080"))
}
