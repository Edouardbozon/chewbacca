package main

import "github.com/edouardbozon/chewbacca/app/http"

func main() {
	app := http.NewServer()
	app.ListenAndServe()
}
