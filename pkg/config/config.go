package config

import (
	"fmt"
	"os"

	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database"
)

type config struct {
	appname           string
	db                database.Connection
	repositoryHandler database.RepositoryHandler
	server            api.Server
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
		fmt.Println(err)
		os.Exit(1)
	}
	c.db = db
	c.repositoryHandler = repositoryHandler
	return c
}

func (c *config) Server(instanceId int) *config {
	server, err := api.NewServerInstanceFactory(instanceId, c.repositoryHandler)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.server = server
	return c
}

func (c *config) StartServer() {
	c.server.Listen()
}

func (c *config) ConnectDatabase() {
	if err := c.db.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *config) DisconnectDatabase() {
	c.db.Disconnect()
}
