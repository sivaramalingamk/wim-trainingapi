package domain
type VehicleData struct {
	ID        string `json:"id"`
	weight int `json:"weight"`
	vehicle_speed int `json:"vehicle_speed"`
	acceleration int `json:"acceleration"`
	heading_direction int `json:"heading_direction"`
	coolent_temp int `json:"coolent_temp"`
	oil_pressure int `json:"oil_pressure"`
	intake_air_temp int `json:"intake_air_temp"`
	rpm int `json:"rpm"`
	engine_load int `json:"engine_load"`
	elevation_angle int `json:"elevation_angle"`
	o2 int `json:"o_2"`
	fuel_flow int `json:"fuel_flow"`
}
