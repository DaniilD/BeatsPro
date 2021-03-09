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
	router := internal.NewRouterBuilder().Build()
	application := internal.NewApplication(configs.GetConfigLocator(), server, router)
	application.Configure()
	application.InitRouts()
	application.Run()
}
