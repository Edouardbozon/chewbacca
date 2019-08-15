package main

import (
	"log"
	"os"

	"github.com/edouardbozon/chewbacca/app/http"
	"github.com/edouardbozon/chewbacca/app/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("APP_DB_USERNAME")
	pwd := os.Getenv("APP_DB_PASSWORD")
	db := os.Getenv("APP_DB_NAME")
	host := os.Getenv("APP_DB_HOST")

	client := &postgres.Client{}
	client.Open(username, pwd, db, host)

	defer client.DB.Close()

	app := http.NewServer()

	app.Handler.CharacterHandler.CharacterService = &postgres.CharacterService{DB: client.DB}
	app.Handler.VehicleHandler.VehicleService = &postgres.VehicleService{DB: client.DB}

	app.ListenAndServe()
}
