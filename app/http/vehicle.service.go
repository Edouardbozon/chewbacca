package http

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/chewbacca/app"
	"github.com/gorilla/mux"
)

// VehicleHandler represents an HTTP API handler for Vehicle.
type VehicleHandler struct {
	Router         *mux.Router
	VehicleService app.VehicleService
	Logger         *log.Logger
}

// NewVehicleHandler returns a new instance of VehicleHandler.
func NewVehicleHandler(r *mux.Router) *VehicleHandler {
	h := &VehicleHandler{
		Router: r,
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}

	h.Router.HandleFunc("/api/characters", h.handleGetCharacters).Methods("GET")
	// h.Router.HandleFunc("/api/character", h.handleGetCharacters).Methods("POST")
	// h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.handleGetCharacters).Methods("GET")
	// h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.handleGetCharacters).Methods("PUT")
	// h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.handleGetCharacters).Methods("DELETE")

	return h
}

func (h *VehicleHandler) handleGetCharacters(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))

	if limit > 10 || limit < 1 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	vehicles, err := h.VehicleService.GetVehicles(limit, offset)
	if err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	encodeJSON(w, vehicles, h.Logger)
}

// func createVehicle(w http.ResponseWriter, r *http.Request) {
// 	var v Vehicle
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&v); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := postgres.VehicleService.createVehicle(v); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, v)
// }

// func getVehicle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid v ID")
// 		return
// 	}

// 	v := Vehicle{ID: id}
// 	if err := v.getVehicle(a.DB); err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			respondWithError(w, http.StatusNotFound, "Vehicle not found")
// 		default:
// 			respondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, v)
// }

// func updateVehicle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid v ID")
// 		return
// 	}

// 	var v Vehicle
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&v); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
// 		return
// 	}
// 	defer r.Body.Close()
// 	v.ID = id

// 	if err := v.updateVehicle(a.DB); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, v)
// }

// func deleteVehicle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid Vehicle ID")
// 		return
// 	}

// 	v := Vehicle{ID: id}
// 	if err := v.deleteVehicle(a.DB); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
// }
