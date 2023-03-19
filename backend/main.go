package main

import (
	"os"
	"simple-compiler/backend/config"
	"simple-compiler/backend/modules/servers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
	cfg := new(config.Config)
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	s := servers.NewServer(cfg)
	s.Start()
}
