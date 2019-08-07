package app

// A Character is an individual person or character within the Star Wars universe.
type Character struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Height      int          `json:"height"`
	Mass        int          `json:"mass"`
	HairColor   string       `json:"hair_color"`
	SkinColor   string       `json:"skin_color"`
	EyeColor    string       `json:"eye_color"`
	BirthYear   string       `json:"birth_year"`
	Gender      string       `json:"gender"`
	Homeworld   string       `json:"homeworld"`
	VehicleURLs []vehicleURL `json:"vehicles"`
	Created     bool         `json:"created"`
	Edited      bool         `json:"edited"`
	URL         string       `json:"url"`
}

type characterURL string

// A CharacterService is a domain service that manage the Character entity.
type CharacterService interface {
	getCharacter(id int) (*Character, error)
	getCharacters() ([]*Character, error)
	createCharacter(u *Character) error
	deleteCharacter(id int) error
}
