package main

import "database/sql"

// A Vehicle is a single transport craft that does not have hyperdrive capability.
type Vehicle struct {
	ID                   int            `json:"id"`
	Name                 string         `json:"name"`
	Model                string         `json:"model"`
	Manufacturer         string         `json:"manufacturer"`
	CostInCredits        string         `json:"cost_in_credits"`
	Length               int            `json:"length"`
	MaxAtmospheringSpeed string         `json:"max_atmosphering_speed"`
	Crew                 string         `json:"crew"`
	Passengers           string         `json:"passengers"`
	CargoCapacity        string         `json:"cargo_capacity"`
	Consumables          string         `json:"consumables"`
	VehicleClass         string         `json:"vehicle_class"`
	PilotURLs            []characterURL `json:"pilots"`
	Created              bool           `json:"created"`
	Edited               bool           `json:"edited"`
	URL                  string         `json:"url"`
}

type vehicleURL string

// getVehicle retrieves the vehicle with the given id
func (vehicle *Vehicle) getVehicle(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM vehicles WHERE id=$1",
		vehicle.ID).Scan(&vehicle.ID)
}

func getVehicles(db *sql.DB, start int, count int) ([]Vehicle, error) {
	rows, err := db.Query(
		"SELECT * FROM vehicles LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	vehicles := []Vehicle{}

	for rows.Next() {
		var vehicle Vehicle
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
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
