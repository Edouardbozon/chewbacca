package postgres

import (
	"database/sql"

	"github.com/edouardbozon/chewbacca/app"
)

// Ensure CharacterService is implementing the app.CharacterService
var _ app.CharacterService = &CharacterService{}

// CharacterService represents the PostgreSQL implementation of app.CharacterService.
type CharacterService struct {
	DB *sql.DB
}

// GetCharacter get a character from DB according to the given id
func (s *CharacterService) GetCharacter(id int) (*app.Character, error) {
	var c app.Character
	row := s.DB.QueryRow("SELECT * FROM characters WHERE id=$1", id)
	if err := row.Scan(&c.ID); err != nil {
		return nil, err
	}
	return &c, nil
}

// GetCharacters get a list of characters from DB according to the given limit and offset
func (s *CharacterService) GetCharacters(limit int, offset int) ([]*app.Character, error) {
	rows, err := s.DB.Query(
		"SELECT * FROM characters LIMIT $1 OFFSET $2",
		limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	characters := []*app.Character{}

	for rows.Next() {
		var c *app.Character
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		characters = append(characters, c)
	}

	return characters, nil
}

// UpdateCharacter update the character in DB with the given id
func (s *CharacterService) UpdateCharacter(c *app.Character) error {
	_, err :=
		s.DB.Exec(`
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
			c.Name,
			c.Height,
			c.Mass,
			c.HairColor,
			c.SkinColor,
			c.EyeColor,
			c.BirthYear,
			c.Gender,
			c.Homeworld,
			c.VehicleURLs,
			c.Created,
			c.Edited,
			c.URL,
			c.ID,
		)

	return err
}

// DeleteCharacter delete the character in DB with the given id
func (s *CharacterService) DeleteCharacter(id int) error {
	_, err := s.DB.Exec("DELETE FROM characters WHERE id=$1", id)

	return err
}

// CreateCharacter create a character in DB
func (s *CharacterService) CreateCharacter(c *app.Character) error {
	err := s.DB.QueryRow(
		`INSERT INTO vehicles (
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
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id`,
		c.Name,
		c.Height,
		c.Mass,
		c.HairColor,
		c.SkinColor,
		c.EyeColor,
		c.BirthYear,
		c.Gender,
		c.Homeworld,
		c.VehicleURLs,
		c.Created,
		c.Edited,
		c.URL).Scan(&c.ID)

	if err != nil {
		return err
	}

	return nil
}
