package main

import (
	"os"

	"github.com/edouardbozon/chewbacca/app/http"
	"github.com/edouardbozon/chewbacca/app/postgres"
)

func main() {
	client := &postgres.Client{}
	client.Open(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	app := http.NewServer()
	app.ListenAndServe()
}
