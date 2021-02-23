package main

import (
	"BeatsProApi/internal"
	"BeatsProApi/internal/configs"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}


func main() {
	application := internal.NewApplication(configs.GetConfigLocator())
	application.Configure()

}
