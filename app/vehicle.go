package app

import "time"

// A Vehicle is a single transport craft that does not have hyperdrive capability.
type Vehicle struct {
	ID                   int            `json:"id"`
	Crew                 int            `json:"crew"`
	Passengers           int            `json:"passengers"`
	Length               int            `json:"length"`
	Name                 string         `json:"name"`
	Model                string         `json:"model"`
	Manufacturer         string         `json:"manufacturer"`
	CostInCredits        string         `json:"cost_in_credits"`
	MaxAtmospheringSpeed string         `json:"max_atmosphering_speed"`
	CargoCapacity        string         `json:"cargo_capacity"`
	Consumables          string         `json:"consumables"`
	VehicleClass         string         `json:"vehicle_class"`
	PilotURLs            []characterURL `json:"pilots"`
	Created              time.Time      `json:"created"`
	Edited               time.Time      `json:"edited"`
	URL                  string         `json:"url"`
}

type vehicleURL string

// A VehicleService is a domain service that manage the Vehicle entity.
type VehicleService interface {
	GetVehicle(id int) (*Vehicle, error)
	GetVehicles(limit int, offset int) ([]*Vehicle, error)
	UpdateVehicle(v *Vehicle) error
	CreateVehicle(v *Vehicle) error
	DeleteVehicle(id int) error
}
