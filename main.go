package main

import (
	"os"

	"github.com/edouardbozon/chewbacca/app/http"
	"github.com/edouardbozon/chewbacca/app/postgres"
)

func main() {
	username := os.Getenv("APP_DB_USERNAME")
	pwd := os.Getenv("APP_DB_PASSWORD")
	db := os.Getenv("APP_DB_NAME")

	client := &postgres.Client{}
	client.Open(username, pwd, db)

	defer client.DB.Close()

	app := http.NewServer()

	app.Handler.CharacterHandler.CharacterService = &postgres.CharacterService{DB: client.DB}
	app.Handler.VehicleHandler.VehicleService = &postgres.VehicleService{DB: client.DB}

	app.ListenAndServe()
}
