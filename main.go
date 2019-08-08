package main

import "github.com/chewbacca/app/http"

func main() {
	app := http.NewServer()
	app.ListenAndServe()
}
