package postgres

import (
	"database/sql"
	"fmt"
	"log"

	// Driver used to connect PostgreSQL database
	_ "github.com/lib/pq"
)

// Client represents a client to the underlying PostgreSQL database.
type Client struct {
	DB *sql.DB
}

// Open opens and initializes the PostgreSQL database.
func (a *Client) Open(user string, password string, dbname string, host string) {
	connection :=
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, dbname, host)

	var err error
	a.DB, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
}
