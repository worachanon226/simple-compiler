package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ServerCheck(r fiber.Router) {
	r.Get("/", _servercheck)
}

func Compile(r fiber.Router) {
	r.Get("/run/:lang", _compile)
}

func Execute(r fiber.Router) {
	r.Get("/output/:lang", _excecute)
}

func _servercheck(c *fiber.Ctx) error {
	return c.SendString("Server is OK.")
}

var filepath = ""

func _compile(c *fiber.Ctx) error {
	lang := c.Params("lang")

	if lang == "cpp" {
		status, file := GenerateFile("cpp", string(c.Body()))
		filepath = file
		return c.SendString(status)
	} else if lang == "py" {
		status, file := GenerateFile("py", string(c.Body()))
		filepath = file
		return c.SendString(status)
	}

	return c.SendStatus(fiber.ErrBadRequest.Code)
}

func _excecute(c *fiber.Ctx) error {
	lang := c.Params("lang")
	if filepath == "" {
		return errors.New("Error, filepath is empty.")
	}
	if lang == "cpp" {
		return c.SendString(ExecuteCpp(filepath))
	} else if lang == "py" {
		return c.SendString(ExecutePy(filepath))
	}

	return c.SendStatus(fiber.ErrBadRequest.Code)
}
