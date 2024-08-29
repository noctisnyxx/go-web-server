package main

import (
	"go-web-server/handler"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	app.Get("/home", func(c fiber.Ctx) error {
		return c.SendString("first web server trial")
	})
	app.Get("/myBookList", handler.GetBookList)
	log.Fatal(app.Listen(":8080"))
}
