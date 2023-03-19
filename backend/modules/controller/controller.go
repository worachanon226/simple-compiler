package controller

import "github.com/gofiber/fiber/v2"

func ServerCheck(r fiber.Router) {
	r.Get("/", _servercheck)
}

func Compile(r fiber.Router) {
	r.Get("/output/:lang", _compile)
}

func _servercheck(c *fiber.Ctx) error {
	return c.SendString("Server is OK.")
}

func _compile(c *fiber.Ctx) error {
	lang := c.Params("lang")
	if lang == "cpp" {
		return c.SendString(GenerateFile("cpp", string(c.Body())))
	} else if lang == "py" {
		return c.SendString(GenerateFile("py", string(c.Body())))
	}

	return c.SendStatus(fiber.ErrBadRequest.Code)
}
