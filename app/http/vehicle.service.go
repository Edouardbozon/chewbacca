package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/edouardbozon/chewbacca/app"
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

	h.Router.HandleFunc("/api/vehicles", h.getCharacters).Methods("GET")
	h.Router.HandleFunc("/api/character", h.createVehicle).Methods("POST")
	h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.getVehicle).Methods("GET")
	h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.updateVehicle).Methods("PUT")
	h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.deleteVehicle).Methods("DELETE")

	return h
}

func (h *VehicleHandler) getCharacters(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))

	if limit > 10 || limit < 1 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	v, err := h.VehicleService.GetVehicles(limit, offset)
	if err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	encodeJSON(w, v, h.Logger)
}

func (h *VehicleHandler) createVehicle(w http.ResponseWriter, r *http.Request) {
	var v *app.Vehicle
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&v); err != nil {
		Error(w, err, http.StatusBadRequest, h.Logger)
		return
	}
	defer r.Body.Close()

	if err := h.VehicleService.CreateVehicle(v); err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	encodeJSON(w, v, h.Logger)
}

func (h *VehicleHandler) getVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		Error(w, err, http.StatusBadRequest, h.Logger)
		return
	}

	v, err := h.VehicleService.GetVehicle(id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			Error(w, err, http.StatusNotFound, h.Logger)
		default:
			Error(w, err, http.StatusInternalServerError, h.Logger)
		}
		return
	}

	encodeJSON(w, v, h.Logger)
}

func (h *VehicleHandler) updateVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		Error(w, err, http.StatusBadRequest, h.Logger)
		return
	}

	var v *app.Vehicle
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&v); err != nil {
		Error(w, err, http.StatusBadRequest, h.Logger)
		return
	}
	defer r.Body.Close()
	v.ID = id

	if err := h.VehicleService.UpdateVehicle(v); err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	encodeJSON(w, v, h.Logger)
}

func (h *VehicleHandler) deleteVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		Error(w, err, http.StatusBadRequest, h.Logger)
		return
	}

	if err := h.VehicleService.DeleteVehicle(id); err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	var v *app.Vehicle
	v.ID = id

	encodeJSON(w, v, h.Logger)
}
