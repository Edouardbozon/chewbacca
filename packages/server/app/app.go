package app

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// App representation
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize app
func (a *App) Initialize(user, password, dbname string) {}

// Run app
func (a *App) Run(addr string) {}
