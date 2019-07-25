package model

import "database/sql"

// A Vehicle is a single transport craft that does not have hyperdrive capability.
type Vehicle struct {
	ID                   string       `json:id`
	Name                 string       `json:"name"`
	Model                string       `json:"model"`
	Manufacturer         string       `json:"manufacturer"`
	CostInCredits        string       `json:"cost_in_credits"`
	Length               int          `json:"length"`
	MaxAtmospheringSpeed string       `json:"max_atmosphering_speed"`
	Crew                 string       `json:"crew"`
	Passengers           string       `json:"passengers"`
	CargoCapacity        string       `json:"cargo_capacity"`
	Consumables          string       `json:"consumables"`
	VehicleClass         string       `json:"vehicle_class"`
	PilotURLs            []vehicleURL `json:"pilots"`
	Created              bool         `json:"created"`
	Edited               bool         `json:"edited"`
	URL                  string       `json:"url"`
}

type vehicleURL string

// VehiclesCreateQuery create vehicles table
const VehiclesCreateQuery = `
	DROP TABLE IF EXISTS vehicles;
	CREATE TABLE vehicles (
		id SERIAL,

		name TEXT UNIQUE NOT NULL,
		model TEXT NOT NULL,
		manufacturer TEXT NOT NULL,
		cost_in_credits TEXT NOT NULL,
		length INT NOT NULL,
		max_atmosphering_speed TEXT NOT NULL,
		crew TEXT NOT NULL,
		passengers TEXT NOT NULL,
		cargo_capacity TEXT NOT NULL,
		consumables TEXT NOT NULL,
		vehicle_class TEXT NOT NULL,
		pilot_urls 
		pilots TEXT NOT NULL,
		created boolean DEFAULT FALSE
		edited boolean DEFAULT FALSE
		url TEXT NOT NULL,

		PRIMARY KEY (id),
		FOREIGN KEY (pilot_urls) REFERENCES books(id) ON DELETE CASCADE,
		UNIQUE (url)
	);
`

// getVehicle retrieves the vehicle with the given id
func (vehicle *Vehicle) getVehicle(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM vehicles WHERE id=$1",
		vehicle.ID).Scan(&vehicle.ID)
}

// updateVehicle update the vehicle with the given id
func (vehicle *Vehicle) updateVehicle(db *sql.DB) error {
	_, err :=
		db.Exec(`
		UPDATE vehicles SET
			name=$1,
			model=$2,
			manufacturer=$3,
			cost_in_credits=$4,
			length=$5,
			max_atmosphering_speed=$6,
			crew=$7,
			passengers=$8,
			cargo_capacity=$9,
			consumables=$10,
			vehicle_class=$11,
			pilot_urls=$12,
			created=$13,
			edited=$14,
			url=$15
		WHERE id=$16`,
			vehicle.Name,
			vehicle.Model,
			vehicle.Manufacturer,
			vehicle.CostInCredits,
			vehicle.Length,
			vehicle.MaxAtmospheringSpeed,
			vehicle.Crew,
			vehicle.Passengers,
			vehicle.CargoCapacity,
			vehicle.Consumables,
			vehicle.VehicleClass,
			vehicle.PilotURLs,
			vehicle.Created,
			vehicle.Edited,
			vehicle.URL,
			vehicle.ID,
		)

	return err
}

// deleteVehicle update the vehicle with the given id
func (vehicle *Vehicle) deleteVehicle(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM vehicles WHERE id=$1", vehicle.ID)

	return err
}

// createVehicle create a vehicle
func (vehicle *Vehicle) createVehicle(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO vehicles(
			name,
			model,
			manufacturer,
			cost_in_credits,
			length,
			max_atmosphering_speed,
			crew,
			passengers,
			cargo_capacity,
			consumables,
			vehicle_class,
			pilot_urls,
			created,
			edited,
			url
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		RETURNING id`,
		vehicle.Name,
		vehicle.Model,
		vehicle.Manufacturer,
		vehicle.CostInCredits,
		vehicle.Length,
		vehicle.MaxAtmospheringSpeed,
		vehicle.Crew,
		vehicle.Passengers,
		vehicle.CargoCapacity,
		vehicle.Consumables,
		vehicle.VehicleClass,
		vehicle.PilotURLs,
		vehicle.Created,
		vehicle.Edited,
		vehicle.URL).Scan(&vehicle.ID)

	if err != nil {
		return err
	}

	return nil
}
