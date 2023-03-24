package controller

import (
	"encoding/json"
	"errors"
	"simple-compiler/backend/modules/entities"

	"github.com/gofiber/fiber/v2"
)

func ServerCheck(r fiber.Router) {
	r.Get("/", _servercheck)
}

func Compile(r fiber.Router) {
	r.Post("/run", _compile)
}

func Execute(r fiber.Router) {
	r.Get("/output", _excecute)
}

func _servercheck(c *fiber.Ctx) error {
	return c.SendString("Server is OK.")
}

type Code struct {
	filepath string
	lang     string
}

var code Code

func _compile(c *fiber.Ctx) error {
	clientCode := entities.ClientCode{}
	if err := c.BodyParser(&clientCode); err != nil {
		return err
	}
	code.filepath = GenerateFile(clientCode.Lang, clientCode.Code)
	code.lang = clientCode.Lang

	return nil
}

func _excecute(c *fiber.Ctx) error {
	if code.filepath == "" {
		return errors.New("Error, filepath is empty.")
	}
	if code.lang == "cpp" {
		output := entities.ClientOutput{
			Output: ExecuteCpp(code.filepath),
		}
		u, err := json.Marshal(output)

		if err != nil {
			return nil
		}

		return c.SendString(string(u))

	} else if code.lang == "py" {
		output := entities.ClientOutput{
			Output: ExecutePy(code.filepath),
		}
		u, err := json.Marshal(output)

		if err != nil {
			return nil
		}
		return c.SendString(string(u))
	}

	return c.SendStatus(fiber.ErrBadRequest.Code)
}
