package main

type Planet struct {
	Name     string `json:"name"`
	Distance int    `json:"distance"`
}

type Planets []Planet

type Vehicle struct {
	Name        string `json:"name"`
	TotalNo     int    `json:"total_no"`
	MaxDistance int    `json:"max_distance"`
	Speed       int    `json:"speed"`
}

type Vehicles []Vehicle

type FindFalconeReq struct {
	Token        string   `json:"token"`
	PlanetNames  []string `json:"planet_names"`
	VehicleNames []string `json:"vehicle_names"`
}
