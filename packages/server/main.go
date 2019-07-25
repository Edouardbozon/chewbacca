package main

import (
	"os"

	"./app"
)

func main() {
	a := app.App{}
	a.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	a.Run(":8080")
}
