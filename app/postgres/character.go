package postgres

import (
	"database/sql"
)

// CharacterService represents a PostgreSQL implementation of app.CharacterService.
type CharacterService struct {
	DB *sql.DB
}

// getCharacter retrieves the character with the given id
func (character *Character) getCharacter(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM characters WHERE id=$1",
		character.ID).Scan(&character.ID)
}

func (character *Character) getCharacters(db *sql.DB, start int, count int) ([]Character, error) {
	rows, err := db.Query(
		"SELECT * FROM characters LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	characters := []Character{}

	for rows.Next() {
		var character Character
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}

	return characters, nil
}

// updateCharacter update the character with the given id
func (character *Character) updateCharacter(db *sql.DB) error {
	_, err :=
		db.Exec(`
		UPDATE characters SET
			name=$1,
			height=$2,
			mass=$3,
			hair_color=$4,
			skin_color=$5,
			eye_color=$6,
			birth_year=$7,
			gender=$8,
			homeworld=$9,
			vehicle_urls=$10,
			created=$11,
			edited=$12,
			url=$13,
		WHERE id=$14`,
			character.Name,
			character.Height,
			character.Mass,
			character.HairColor,
			character.SkinColor,
			character.EyeColor,
			character.BirthYear,
			character.Gender,
			character.Homeworld,
			character.VehicleURLs,
			character.Created,
			character.Edited,
			character.URL,
			character.ID,
		)

	return err
}

// deleteCharacter update the character with the given id
func (character *Character) deleteCharacter(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM characters WHERE id=$1", character.ID)

	return err
}

// createCharacter create a character
func (character *Character) createCharacter(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO vehicles(
			name,
			height,
			mass,
			hair_color,
			skin_color,
			eye_color,
			birth_year,
			gender,
			homeworld,
			vehicle_urls,
			created,
			edited,
			url,
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id`,
		character.Name,
		character.Height,
		character.Mass,
		character.HairColor,
		character.SkinColor,
		character.EyeColor,
		character.BirthYear,
		character.Gender,
		character.Homeworld,
		character.VehicleURLs,
		character.Created,
		character.Edited,
		character.URL).Scan(&character.ID)

	if err != nil {
		return err
	}

	return nil
}
