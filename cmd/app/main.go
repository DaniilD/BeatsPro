package main

import (
	"BeatsPro/internal"
	"BeatsPro/internal/configs"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	server := internal.NewServer()
	application := internal.NewApplication(configs.GetConfigLocator(), server)
	application.Configure()
	application.InitSentry()
	application.Router = internal.NewRouterBuilder().Build()
	application.InitRouts()
	application.Run()
}
