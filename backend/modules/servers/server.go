package servers

import (
	"log"
	"simple-compiler/backend/config"
	"simple-compiler/backend/pkg/utils"

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
	fiberConn, err := utils.ConnectionBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port

	log.Printf("server has been started on %s:%s", host, port)

	if err := s.App.Listen(fiberConn); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
