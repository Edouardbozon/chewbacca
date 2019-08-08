package postgres

import (
	"database/sql"

	"github.com/chewbacca/app"
	_ "github.com/lib/pq"
)

// Ensure VehicleService is implementing the app.VehicleService
var _ app.VehicleService = &VehicleService{}

// VehicleService represents the PostgreSQL implementation of app.VehicleService.
type VehicleService struct {
	db *sql.DB
}

func (s *VehicleService) GetVehicle(id int) (*app.Vehicle, error) {
	var v app.Vehicle
	row := s.db.QueryRow("SELECT * FROM vehicles WHERE id=$1", id)
	if err := row.Scan(&v.ID); err != nil {
		return nil, err
	}
	return &v, nil
}

func (s *VehicleService) GetVehicles(limit int, offset int) ([]*app.Vehicle, error) {
	rows, err := s.db.Query(
		"SELECT * FROM vehicles LIMIT $1 OFFSET $2",
		limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	vehicles := []*app.Vehicle{}

	for rows.Next() {
		var v *app.Vehicle
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, v)
	}

	return vehicles, nil
}

// updateVehicle update the v with the given id
// func (s *VehicleService) updateVehicle(id int) error {
// 	_, err :=
// 		s.db.Exec(`
// 		UPDATE vehicles SET
// 			name=$1,
// 			model=$2,
// 			manufacturer=$3,
// 			cost_in_credits=$4,
// 			length=$5,
// 			max_atmosphering_speed=$6,
// 			crew=$7,
// 			passengers=$8,
// 			cargo_capacity=$9,
// 			consumables=$10,
// 			vehicle_class=$11,
// 			pilot_urls=$12,
// 			created=$13,
// 			edited=$14,
// 			url=$15
// 		WHERE id=$16`,
// 			v.Name,
// 			v.Model,
// 			v.Manufacturer,
// 			v.CostInCredits,
// 			v.Length,
// 			v.MaxAtmospheringSpeed,
// 			v.Crew,
// 			v.Passengers,
// 			v.CargoCapacity,
// 			v.Consumables,
// 			v.VehicleClass,
// 			v.PilotURLs,
// 			v.Created,
// 			v.Edited,
// 			v.URL,
// 			v.ID,
// 		)

// 	return err
// }

// deleteVehicle update the v with the given id
func (s *VehicleService) DeleteVehicle(id int) error {
	_, err := s.db.Exec("DELETE FROM vehicles WHERE id=$1", id)

	return err
}

// CreateVehicle create a v
func (s *VehicleService) CreateVehicle(v *app.Vehicle) error {
	err := s.db.QueryRow(
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
		v.Name,
		v.Model,
		v.Manufacturer,
		v.CostInCredits,
		v.Length,
		v.MaxAtmospheringSpeed,
		v.Crew,
		v.Passengers,
		v.CargoCapacity,
		v.Consumables,
		v.VehicleClass,
		v.PilotURLs,
		v.Created,
		v.Edited,
		v.URL).Scan(&v.ID)

	if err != nil {
		return err
	}

	return nil
}