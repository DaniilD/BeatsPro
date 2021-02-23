package internal

import (
	"BeatsPro/internal/configs"
	"log"
	"os"
	"strconv"
)

type Application struct {
	configLocator *configs.ConfigLocator
}

func NewApplication(configLocator *configs.ConfigLocator) *Application {
	return &Application{configLocator: configLocator}
}

func (application *Application) Configure() {
	application.configureServer()
	application.configureDB()
}

func (application *Application) configureServer() {
	host := os.Getenv("APPLICATION_HOST")
	port, err := strconv.Atoi(os.Getenv("APPLICATION_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	application.configLocator.ServerConfigInstance().Host = host
	application.configLocator.ServerConfigInstance().Port = port

}

func (application *Application) configureDB() {
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	application.configLocator.DBConfigInstance().Host = host
	application.configLocator.DBConfigInstance().Name = name
	application.configLocator.DBConfigInstance().User = user
	application.configLocator.DBConfigInstance().Password = password
	application.configLocator.DBConfigInstance().Port = port
}
