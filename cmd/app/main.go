package main

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/config"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := config.New().AppName("My finances").Logger(logger.ZapInstance)
	app.Database(database.PostgresInstance).ConnectDatabase()
	defer app.DisconnectDatabase()

	app.Server(api.EchoInstance).StartServer()
}
