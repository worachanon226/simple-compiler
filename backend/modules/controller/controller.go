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

func _servercheck(c *fiber.Ctx) error {
	return c.SendString("Server is OK.")
}

func _compile(c *fiber.Ctx) error {
	clientCode := entities.ClientCode{}

	if err := c.BodyParser(&clientCode); err != nil {
		return err
	}

	filepath := GenerateFile(clientCode.Lang, clientCode.Code)
	if clientCode.Lang == "cpp" {
		output := entities.ClientOutput{
			Output: ExecuteCpp(filepath),
		}
		u, err := json.Marshal(output)

		if err != nil {
			return err
		}

		return c.SendString(string(u))

	} else if clientCode.Lang == "py" {
		output := entities.ClientOutput{
			Output: ExecutePy(filepath),
		}
		u, err := json.Marshal(output)

		if err != nil {
			return err
		}
		return c.SendString(string(u))
	}
	return errors.New("Language is wrong.")
}
