package config

import (
	"log"

	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/logger"
)

type config struct {
	appname           string
	db                database.Connection
	repositoryHandler database.RepositoryHandler
	server            api.Server
	logger logger.Logger
}

func New() *config {
	return &config{}
}

func (c *config) AppName(appname string) *config {
	c.appname = appname
	return c
}

func (c *config) Database(instanceId int) *config {
	db, repositoryHandler, err := database.NewSQLDatabaseFactory(instanceId)
	if err != nil {
		log.Fatalln("error configuring database: %w", err)
	}

	c.logger.InfoF("Database configured successfully!")

	c.db = db
	c.repositoryHandler = repositoryHandler
	return c
}

func (c *config) Server(instanceId int) *config {
	server, err := api.NewServerInstanceFactory(instanceId, c.repositoryHandler, c.logger)
	if err != nil {
		log.Fatalln("error initializing webserver: %w", err)
	}

	c.logger.InfoF("WebServer configured successfully!")

	c.server = server
	return c
}

func (c *config) Logger(instanceId int) *config {
	logger, err := logger.NewLoggerFactory(instanceId)
	if err != nil {
		log.Fatalln("error initializing logger: %w", err)
	}
	
	c.logger = logger
	
	c.logger.InfoF("Log configured successfully!")
	
	return c
}

func (c *config) StartServer() {
	c.server.Listen()
}

func (c *config) ConnectDatabase() {
	if err := c.db.Connect(); err != nil {
		log.Fatalln("error connecting to database: %w", err)
	}
}

func (c *config) DisconnectDatabase() {
	if err := c.db.Disconnect(); err != nil {
		log.Fatalln("error disconnecting to database: %w", err)
	}
}
