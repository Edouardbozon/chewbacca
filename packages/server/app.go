package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	a.Router.HandleFunc("/vehicles", a.getVehicles).Methods("GET")
	a.Router.HandleFunc("/vehicle", a.createVehicle).Methods("POST")
	a.Router.HandleFunc("/vehicle/{id:[0-9]+}", a.getVehicle).Methods("GET")
	a.Router.HandleFunc("/vehicle/{id:[0-9]+}", a.updateVehicle).Methods("PUT")
	a.Router.HandleFunc("/vehicle/{id:[0-9]+}", a.deleteVehicle).Methods("DELETE")
}

func (a *App) getVehicles(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	vehicles, err := getVehicles(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, vehicles)
}

func (a *App) createVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle Vehicle
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&vehicle); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := vehicle.createVehicle(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, vehicle)
}

func (a *App) getVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid vehicle ID")
		return
	}

	vehicle := Vehicle{ID: id}
	if err := vehicle.getVehicle(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Vehicle not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, vehicle)
}

func (a *App) updateVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid vehicle ID")
		return
	}

	var vehicle Vehicle
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&vehicle); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	vehicle.ID = id

	if err := vehicle.updateVehicle(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, vehicle)
}

func (a *App) deleteVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Vehicle ID")
		return
	}

	vehicle := Vehicle{ID: id}
	if err := vehicle.deleteVehicle(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
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
