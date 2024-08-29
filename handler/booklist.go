package handler

import (
	"github.com/gofiber/fiber/v3"
)

func GetBookList(c fiber.Ctx) error {
	return c.SendString("getbooklist endpoint")
}
