package main

import (
	"fmt"
	"os"

	"github.com/duardoqueiroz/my-finances-golang/pkg/config"
	"github.com/duardoqueiroz/my-finances-golang/pkg/factories"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api/routes"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database/repositories/postgres"
)

func main() {
	conf, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db := postgres.SetupCredentials(conf.Database.User, conf.Database.Password, conf.Database.Port, conf.Database.Name, conf.Database.Host)
	if err := db.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Disconnect()

	conn := db.Connection()

	repoFactory := factories.NewRepositoryFactory(postgres.NewUserRepository(conn))

	server := routes.LoadRoutes(repoFactory)
	server.Debug = true
	address := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)
	server.Logger.Info(server.Start(address))
}
