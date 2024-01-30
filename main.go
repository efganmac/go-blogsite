package main

import (
	"go-blogsite/config"
	"go-blogsite/server"
)

func main() {

	cfg, _ := config.SetupEnv()

	server.StartServer(cfg)
}
