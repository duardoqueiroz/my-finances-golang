package main

import (
	"fmt"
	"net/http"

	"github.com/duardoqueiroz/my-finances-golang/pkg/config"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database/repositories/postgres"
)

func main() {
	config.ReadConfig()

	db := postgres.SetupCredentials(config.C.Database.User, config.C.Database.Password, config.C.Database.Port, config.C.Database.Name, config.C.Database.Host)
	db.Connect()
	defer db.Disconnect()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(fmt.Sprintf(":%s", config.C.Server.Port), nil)
}
