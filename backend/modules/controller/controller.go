package controller

import "github.com/gofiber/fiber/v2"

func ServerCheck(r fiber.Router) {
	r.Get("/", _servercheck)
}

func _servercheck(c *fiber.Ctx) error {
	return c.SendString("Server is OK.")
}
