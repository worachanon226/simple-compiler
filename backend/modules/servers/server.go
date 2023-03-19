package servers

import (
	"simple-compiler/backend/config"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
	Cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		App: fiber.New(),
		Cfg: cfg,
	}
}

func (s *Server) Start() {

}
