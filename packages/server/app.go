package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App representation
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize app
func (a *App) Initialize(user string, password string, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	log.Print(connectionString)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	defer a.DB.Close()

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	// a.Router.HandleFunc("/characters", a.getProducts).Methods("GET")
	// a.Router.HandleFunc("/character", a.createProduct).Methods("POST")
	// a.Router.HandleFunc("/character/{id:[0-9]+}", a.getProduct).Methods("GET")
	// a.Router.HandleFunc("/character/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	// a.Router.HandleFunc("/character/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

// Run app
func (a *App) Run(addr string) {
	log.Println(fmt.Sprintf("Server started on: http://localhost%s", addr))
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
