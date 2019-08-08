package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler is a collection of all the service handlers.
type Handler struct {
	CharacterHandler *CharacterHandler
	VehicleHandler   *VehicleHandler
	Router           *mux.Router
}

// NewHandler return a Handler instance
func NewHandler() *Handler {
	r := mux.NewRouter()

	return &Handler{
		Router:           r,
		CharacterHandler: NewCharacterHandler(r),
		VehicleHandler:   NewVehicleHandler(r),
	}
}

// Error writes an API error message to the response and logger.
func Error(w http.ResponseWriter, err error, code int, logger *log.Logger) {
	// Log error.
	logger.Printf("http error: %s (code=%d)", err, code)

	// Hide internal errors to client
	if code == http.StatusInternalServerError {
		err = &internalErr{msg: "Internal Server Error"}
	}

	// Write generic error response.
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err.Error())
}

// encodeJSON encodes v to w in JSON format. Error() is called if encoding fails.
func encodeJSON(w http.ResponseWriter, v interface{}, logger *log.Logger) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		Error(w, err, http.StatusInternalServerError, logger)
	}
}
