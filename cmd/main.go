package main

import (
	"log"

	"github.com/gameloungee/server/internal/server"
	"github.com/joho/godotenv"
)

// @Title Gamelounge API
// @version 0.0.0α
// @description API for the Gamelounge project

// @host sgamelounge.onrender.com
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env files found")
	}
}

func main() {
	server := new(server.Server)
	addr := server.MakeAddr()
	server.Run(addr)
}
