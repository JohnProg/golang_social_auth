package main

import (
	"github.com/golang_social_auth/settings"
	"log"

	"net/http"
	"github.com/golang_social_auth/server"
)

const (
	configPath = "config.json"
)

func main() {
	app := server.Server{}
	if err := settings.Read(configPath); err != nil {
		log.Fatal("Could not read config file at " + configPath + " " + err.Error())
	}

	app.Initialize()
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
