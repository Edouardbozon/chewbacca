package http

import (
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

	h.Router.HandleFunc("/api/characters", h.handleGetCharacters).Methods("GET")
	// h.Router.HandleFunc("/api/character", h.handleGetCharacters).Methods("POST")
	// h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.handleGetCharacters).Methods("GET")
	// h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.handleGetCharacters).Methods("PUT")
	// h.Router.HandleFunc("/api/character/{id:[0-9]+}", h.handleGetCharacters).Methods("DELETE")

	return h
}

func (h *CharacterHandler) handleGetCharacters(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	characters, err := h.CharacterService.GetCharacters(start, count)
	if err != nil {
		Error(w, err, http.StatusInternalServerError, h.Logger)
		return
	}

	encodeJSON(w, characters, h.Logger)
}

// func createCharacter(w http.ResponseWriter, r *http.Request) {
// 	var c Character
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&c); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := postgres.CharacterService.createCharacter(c); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, c)
// }

// func getCharacter(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid c ID")
// 		return
// 	}

// 	c := Character{ID: id}
// 	if err := c.getCharacter(a.DB); err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			respondWithError(w, http.StatusNotFound, "Character not found")
// 		default:
// 			respondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, c)
// }

// func updateCharacter(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid c ID")
// 		return
// 	}

// 	var c Character
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&c); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
// 		return
// 	}
// 	defer r.Body.Close()
// 	c.ID = id

// 	if err := c.updateCharacter(a.DB); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, c)
// }

// func deleteCharacter(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid Character ID")
// 		return
// 	}

// 	c := Character{ID: id}
// 	if err := c.deleteCharacter(a.DB); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
// }
