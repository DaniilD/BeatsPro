package internal

import (
	"BeatsPro/internal/configs"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Application struct {
	configLocator *configs.ConfigLocator
	server        *Server
	Router        *Router
}

func NewApplication(configLocator *configs.ConfigLocator, server *Server) *Application {
	return &Application{
		configLocator: configLocator,
		server:        server,
	}
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
	driver := os.Getenv("DB_DRIVER")
	sslMode := os.Getenv("DB_SSLMODE")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	application.configLocator.DBConfigInstance().Host = host
	application.configLocator.DBConfigInstance().Name = name
	application.configLocator.DBConfigInstance().User = user
	application.configLocator.DBConfigInstance().Password = password
	application.configLocator.DBConfigInstance().Port = port
	application.configLocator.DBConfigInstance().Driver = driver
	application.configLocator.DBConfigInstance().SslMode = sslMode
}

func (application *Application) Run() {
	port := application.configLocator.ServerConfigInstance().Port
	host := application.configLocator.ServerConfigInstance().Host

	if err := application.server.Run(host, fmt.Sprint(port), application.Router); err != nil {
		log.Fatal(err)
	}
}

func (application *Application) InitRouts() {
	application.Router.InitRouts()
}
