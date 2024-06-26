package main

import (
	"fmt"
	"net/http"

	"github.com/duardoqueiroz/my-finances-golang/pkg/config"
)

func main() {
	config.ReadConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(fmt.Sprintf(":%s", config.C.Server.Port), nil)
}
