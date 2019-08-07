package main

import (
	"database/sql"
	"encoding/json"
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

// Run app
func (a *App) Run(addr string) {
	log.Println(fmt.Sprintf("Server started on: http://localhost%s", addr))
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	// Vehicle
	a.Router.HandleFunc("/vehicles", a.getVehicles).Methods("GET")
	a.Router.HandleFunc("/vehicle", a.createVehicle).Methods("POST")
	a.Router.HandleFunc("/vehicle/{id:[0-9]+}", a.getVehicle).Methods("GET")
	a.Router.HandleFunc("/vehicle/{id:[0-9]+}", a.updateVehicle).Methods("PUT")
	a.Router.HandleFunc("/vehicle/{id:[0-9]+}", a.deleteVehicle).Methods("DELETE")

	// Character
	a.Router.HandleFunc("/characters", a.getCharacters).Methods("GET")
	a.Router.HandleFunc("/character", a.createCharacter).Methods("POST")
	a.Router.HandleFunc("/character/{id:[0-9]+}", a.getCharacter).Methods("GET")
	a.Router.HandleFunc("/character/{id:[0-9]+}", a.updateCharacter).Methods("PUT")
	a.Router.HandleFunc("/character/{id:[0-9]+}", a.deleteCharacter).Methods("DELETE")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
