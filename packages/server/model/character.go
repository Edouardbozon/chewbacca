package model

import "database/sql"

// A Character is an individual person or character within the Star Wars universe.
type Character struct {
	Name        string       `json:"name"`
	Height      string       `json:"height"`
	Mass        string       `json:"mass"`
	HairColor   string       `json:"hair_color"`
	SkinColor   string       `json:"skin_color"`
	EyeColor    string       `json:"eye_color"`
	BirthYear   string       `json:"birth_year"`
	Gender      string       `json:"gender"`
	Homeworld   string       `json:"homeworld"`
	VehicleURLs []vehicleURL `json:"vehicles"`
	Created     string       `json:"created"`
	Edited      string       `json:"edited"`
	URL         string       `json:"url"`
}

type characterURL string

// getCharacter retrieves the character with the given id
func (p *Character) getCharacter(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM characters WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
}

// updateCharacter update the character with the given id
func (p *Character) updateCharacter(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE characters SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

// deleteCharacter update the character with the given id
func (p *Character) deleteCharacter(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM characters WHERE id=$1", p.ID)

	return err
}

// createCharacter create a character
func (p *Character) createCharacter(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO characters(name, price) VALUES($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}
