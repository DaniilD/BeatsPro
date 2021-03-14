package main

import (
	"BeatsPro/internal"
	"BeatsPro/internal/configs"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"log"
	"time"
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

	application.Router = internal.NewRouterBuilder().Build()

	//init sentry
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: "https://57fadf5d1c524ba5ba175f60daa60812@o551473.ingest.sentry.io/5674965",
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Environment: "",
		Release:     "my-project-name@1.0.0",
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage(" привет я даниил")

	application.InitRouts()
	application.Run()
}
