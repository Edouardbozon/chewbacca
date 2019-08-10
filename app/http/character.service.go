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

// CharacterHandler represents an HTTP API handler for Character.
type CharacterHandler struct {
	Router           *mux.Router
	CharacterService app.CharacterService
	Logger           *log.Logger
}

// NewCharacterHandler returns a new instance of CharacterHandler.
func NewCharacterHandler(r *mux.Router) *CharacterHandler {
	h := &CharacterHandler{
		Router: r,
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}

	h.Router.HandleFunc("/api/characters", h.getCharacters).Methods("GET")
	h.Router.HandleFunc("/api/character", h.createCharacter).Methods("POST")
	h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.getCharacter).Methods("GET")
	h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.updateCharacter).Methods("PUT")
	h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.deleteCharacter).Methods("DELETE")

	return h
}

func (h *CharacterHandler) getCharacters(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	if offset > 10 || offset < 1 {
		offset = 10
	}
	if limit < 0 {
		limit = 0
	}

	c, err := h.CharacterService.GetCharacters(limit, offset)
	if err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	encodeJSON(w, c, h.Logger)
}

func (h *CharacterHandler) createCharacter(w http.ResponseWriter, r *http.Request) {
	var c *app.Character
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}
	defer r.Body.Close()

	if err := h.CharacterService.CreateCharacter(c); err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	encodeJSON(w, c, h.Logger)
}

func (h *CharacterHandler) getCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		Error(w, err, http.StatusNotFound, h.Logger)
		return
	}

	c, err := h.CharacterService.GetCharacter(id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			Error(w, err, http.StatusNotFound, h.Logger)
		default:
			Error(w, err, http.StatusInternalServerError, h.Logger)
		}
		return
	}

	encodeJSON(w, c, h.Logger)
}

func (h *CharacterHandler) updateCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	var c *app.Character
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		Error(w, err, http.StatusNotFound, h.Logger)
		return
	}
	defer r.Body.Close()
	c.ID = id

	if err := h.CharacterService.UpdateCharacter(c); err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	encodeJSON(w, c, h.Logger)
}

func (h *CharacterHandler) deleteCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		Error(w, err, http.StatusBadRequest, h.Logger)
		return
	}

	if err := h.CharacterService.DeleteCharacter(id); err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	var c *app.Character
	c.ID = id

	encodeJSON(w, c, h.Logger)
}
